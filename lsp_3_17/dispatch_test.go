package lsp

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/two-hundred/ls-builder/server"
)

type DispatchTestSuite struct {
	suite.Suite
}

func (s *DispatchTestSuite) Test_server_sends_initialized_notification() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := newTestLSPContext(ctx, container.serverConn)
	dispatcher := NewDispatcher(lspCtx)
	err := dispatcher.Initialized()
	s.Require().NoError(err)

	// Wait some time for the client to receive the message as JSON RPC
	// notifications do not wait for a response.
	time.Sleep(50 * time.Millisecond)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()
	// Verify that the client received the initialized notification.
	s.Require().Len(container.clientReceivedMessages, 1)
	var message InitializedParams
	err = json.Unmarshal(*container.clientReceivedMessages[0], &message)
	s.Require().NoError(err)
}

func (s *DispatchTestSuite) Test_server_sends_register_capability_message() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := newTestLSPContext(ctx, container.serverConn)
	dispatcher := NewDispatcher(lspCtx)

	registerCapabilityParams := RegistrationParams{
		Registrations: []Registration{
			{
				ID:     "1",
				Method: "textDocument/didOpen",
				RegisterOptions: map[string]interface{}{
					"documentSelector": []interface{}{
						map[string]interface{}{
							"language": "go",
							"scheme":   "file",
						},
					},
				},
			},
		},
	}
	err := dispatcher.RegisterCapability(registerCapabilityParams)
	s.Require().NoError(err)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()
	// Verify that the client received the register capability message.
	s.Require().Len(container.clientReceivedMessages, 1)
	var message RegistrationParams
	err = json.Unmarshal(*container.clientReceivedMessages[0], &message)
	s.Require().NoError(err)
	s.Require().Equal(registerCapabilityParams, message)
}

func TestDispatchTestSuite(t *testing.T) {
	suite.Run(t, new(DispatchTestSuite))
}
