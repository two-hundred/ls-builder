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

func (s *DispatchTestSuite) Test_server_sends_progress_notification() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	progressToken := "progress-test-token"
	progressParams := ProgressParams{
		Token: &IntOrString{StrVal: &progressToken},
	}

	err := dispatcher.Progress(progressParams)
	s.Require().NoError(err)

	// Let some time pass as this is a notification,
	// which by definition in LSP and JSON-RPC is fire-and-forget
	// so we can't know at this point if the client has received the message.
	time.Sleep(10 * time.Millisecond)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()
	// Verify that the client received the progress message.
	s.Require().Len(container.clientReceivedMessages, 1)
	var message ProgressParams
	err = json.Unmarshal(*container.clientReceivedMessages[0], &message)
	s.Require().NoError(err)
	s.Require().Equal(progressParams, message)
}

func (s *DispatchTestSuite) Test_server_sends_cancel_request_notification() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	requestID := "test-request-id"
	cancelParams := CancelParams{
		ID: &IntOrString{StrVal: &requestID},
	}

	err := dispatcher.CancelRequest(cancelParams)
	s.Require().NoError(err)

	// Let some time pass as this is a notification,
	// which by definition in LSP and JSON-RPC is fire-and-forget
	// so we can't know at this point if the client has received the message.
	time.Sleep(10 * time.Millisecond)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()
	// Verify that the client received the cancel request message.
	s.Require().Len(container.clientReceivedMessages, 1)
	var message CancelParams
	err = json.Unmarshal(*container.clientReceivedMessages[0], &message)
	s.Require().NoError(err)
	s.Require().Equal(cancelParams, message)
}

func (s *DispatchTestSuite) Test_server_sends_log_trace_notification() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	logTraceParams := LogTraceParams{
		Message: "Something interesting happened",
	}

	err := dispatcher.LogTrace(logTraceParams)
	s.Require().NoError(err)

	// Let some time pass as this is a notification,
	// which by definition in LSP and JSON-RPC is fire-and-forget
	// so we can't know at this point if the client has received the message.
	time.Sleep(10 * time.Millisecond)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()
	// Verify that the client received the cancel request message.
	s.Require().Len(container.clientReceivedMessages, 1)
	var message LogTraceParams
	err = json.Unmarshal(*container.clientReceivedMessages[0], &message)
	s.Require().NoError(err)
	s.Require().Equal(logTraceParams, message)
}

func (s *DispatchTestSuite) Test_server_sends_register_capability_message() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
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

func (s *DispatchTestSuite) Test_server_sends_deregister_capability_message() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	deregisterCapabilityParams := UnregistrationParams{
		Unregistrations: []Unregistration{
			{
				ID:     "1",
				Method: "textDocument/didOpen",
			},
		},
	}
	err := dispatcher.UnregisterCapability(deregisterCapabilityParams)
	s.Require().NoError(err)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()
	// Verify that the client received the de-register capability message.
	s.Require().Len(container.clientReceivedMessages, 1)
	var message UnregistrationParams
	err = json.Unmarshal(*container.clientReceivedMessages[0], &message)
	s.Require().NoError(err)
	s.Require().Equal(deregisterCapabilityParams, message)
}

func (s *DispatchTestSuite) Test_server_sends_refresh_code_lens_message() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	err := dispatcher.CodeLensRefresh()
	s.Require().NoError(err)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()
	// Verify that the client received the refresh message.
	s.Require().Len(container.clientReceivedMessages, 1)
	// Verify the method name.
	s.Require().Len(container.clientReceivedMethods, 1)
	s.Require().Equal(MethodCodeLensRefresh, container.clientReceivedMethods[0])
}

func (s *DispatchTestSuite) Test_server_sends_refresh_semantic_tokens_message() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	err := dispatcher.SemanticTokensRefresh()
	s.Require().NoError(err)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()
	// Verify that the client received the refresh message.
	s.Require().Len(container.clientReceivedMessages, 1)
	// Verify the method name.
	s.Require().Len(container.clientReceivedMethods, 1)
	s.Require().Equal(MethodSemanticTokensRefresh, container.clientReceivedMethods[0])
}

func (s *DispatchTestSuite) Test_server_sends_refresh_inlay_hints_message() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	err := dispatcher.InlayHintRefresh()
	s.Require().NoError(err)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()
	// Verify that the client received the refresh message.
	s.Require().Len(container.clientReceivedMessages, 1)
	// Verify the method name.
	s.Require().Len(container.clientReceivedMethods, 1)
	s.Require().Equal(MethodInlayHintRefresh, container.clientReceivedMethods[0])
}

func (s *DispatchTestSuite) Test_server_sends_refresh_inline_value_message() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	err := dispatcher.InlineValueRefresh()
	s.Require().NoError(err)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()
	// Verify that the client received the refresh message.
	s.Require().Len(container.clientReceivedMessages, 1)
	// Verify the method name.
	s.Require().Len(container.clientReceivedMethods, 1)
	s.Require().Equal(MethodInlineValueRefresh, container.clientReceivedMethods[0])
}

func TestDispatchTestSuite(t *testing.T) {
	suite.Run(t, new(DispatchTestSuite))
}
