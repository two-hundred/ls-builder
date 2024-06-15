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

func (s *HandlerTestSuite) Test_calls_initialize_request_handler_and_sets_initialized_state() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	version := "0.0.1"
	openClose := true
	initializeResult := InitializeResult{
		Capabilities: ServerCapabilities{
			PositionEncoding: PositionEncodingKindUTF8,
			TextDocumentSync: TextDocumentSyncOptions{
				OpenClose: &openClose,
				Change:    &TextDocumentSyncKindIncremental,
			},
		},
		ServerInfo: &InitializeResultServerInfo{
			Name:    "test-language-server",
			Version: &version,
		},
	}

	serverHandler := NewHandler(
		WithInitializeHandler(
			func(ctx *common.LSPContext, params *InitializeParams) (any, error) {
				return initializeResult, nil
			},
		),
	)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	// Assert before the LSP initialisation process.
	s.Require().False(serverHandler.IsInitialized())

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	willSave := true
	didSave := true
	initializeParams := InitializeParams{
		ClientInfo: &InitializeClientInfo{
			Name:    "test-client",
			Version: "0.1.0",
		},
		Capabilities: ClientCapabilities{
			TextDocument: &TextDocumentClientCapabilities{
				Synchronization: &TextDocumentSyncClientCapabilities{
					WillSave: &willSave,
					DidSave:  &didSave,
				},
			},
			General: &GeneralClientCapabilities{
				PositionEncodings: []PositionEncodingKind{PositionEncodingKindUTF8},
			},
		},
	}

	returnedResult := InitializeResult{}
	err = clientLSPContext.Call(MethodInitialize, initializeParams, &returnedResult)
	s.Require().NoError(err)
	s.Require().Equal(initializeResult, returnedResult)
	s.Require().True(serverHandler.IsInitialized())
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
