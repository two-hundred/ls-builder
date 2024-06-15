package lsp

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/two-hundred/ls-builder/common"
	"github.com/two-hundred/ls-builder/server"
	"go.uber.org/zap"
)

type HandlerTestSuite struct {
	suite.Suite
}

func (s *HandlerTestSuite) Test_calls_cancel_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *CancelParams, 1)
	serverHandler := NewHandler(
		WithCancelRequestHandler(
			func(ctx *common.LSPContext, params *CancelParams) error {
				callChan <- params
				return nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)
	clientDispatcher := NewDispatcher(clientLSPContext)

	cancelID := "test-request-id"
	cancelParams := CancelParams{
		ID: &IntOrString{StrVal: &cancelID},
	}
	err = clientDispatcher.CancelRequest(cancelParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(cancelParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_progress_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *ProgressParams, 1)
	serverHandler := NewHandler(
		WithProgressHandler(
			func(ctx *common.LSPContext, params *ProgressParams) error {
				callChan <- params
				return nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)
	clientDispatcher := NewDispatcher(clientLSPContext)

	progressToken := "test-progress-token"
	progressParams := ProgressParams{
		Token: &ProgressToken{StrVal: &progressToken},
	}
	err = clientDispatcher.Progress(progressParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(progressParams, *receivedParams)
	}
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
