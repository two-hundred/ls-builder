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

func (s *DispatchTestSuite) Test_server_sends_push_diagnostics_notification() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	diagnosticSeverityWarning := DiagnosticSeverityWarning
	diagnosticCode := Integer(1405)
	version := Integer(8)
	source := "typescript"
	params := PublishDiagnosticsParams{
		URI:     "file:///path/to/file.go",
		Version: &version,
		Diagnostics: []Diagnostic{
			{
				Range: Range{
					Start: Position{
						Line:      10,
						Character: 5,
					},
					End: Position{
						Line:      50,
						Character: 20,
					},
				},
				Severity: &diagnosticSeverityWarning,
				Code: &IntOrString{
					IntVal: &diagnosticCode,
				},
				Source: &source,
				Tags:   []DiagnosticTag{DiagnosticTagUnnecessary, DiagnosticTagDeprecated},
			},
		},
	}
	err := dispatcher.PublishDiagnostics(params)
	s.Require().NoError(err)

	// Let some time pass as this is a notification,
	// which by definition in LSP and JSON-RPC is fire-and-forget
	// so we can't know at this point if the client has received the message.
	time.Sleep(10 * time.Millisecond)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()
	// Verify that the client received the refresh message.
	s.Require().Len(container.clientReceivedMessages, 1)

	var message PublishDiagnosticsParams
	err = json.Unmarshal(*container.clientReceivedMessages[0], &message)
	s.Require().NoError(err)
	s.Require().Equal(params, message)

	// Verify the method name.
	s.Require().Len(container.clientReceivedMethods, 1)
	s.Require().Equal(MethodPublishDiagnostics, container.clientReceivedMethods[0])
}

func (s *DispatchTestSuite) Test_server_sends_diagnostics_refresh_request() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	err := dispatcher.DiagnosticsRefresh()
	s.Require().NoError(err)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()

	// Verify the method name.
	s.Require().Len(container.clientReceivedMethods, 1)
	s.Require().Equal(MethodDiagnosticsRefresh, container.clientReceivedMethods[0])
}

func (s *DispatchTestSuite) Test_server_sends_workspace_configuration_request() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	scopeURI := "file:///path/to/file.go"
	section := "editor"
	params := ConfigurationParams{
		Items: []ConfigurationItem{
			{
				ScopeURI: &scopeURI,
				Section:  &section,
			},
		},
	}
	target := map[string]interface{}{}
	err := dispatcher.WorkspaceConfiguration(params, &target)
	s.Require().NoError(err)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()

	// Verify that the client received the configuration message.
	s.Require().Len(container.clientReceivedMessages, 1)

	var message ConfigurationParams
	err = json.Unmarshal(*container.clientReceivedMessages[0], &message)
	s.Require().NoError(err)
	s.Require().Equal(params, message)

	// Verify the method name.
	s.Require().Len(container.clientReceivedMethods, 1)
	s.Require().Equal(MethodWorkspaceConfiguration, container.clientReceivedMethods[0])
}

func (s *DispatchTestSuite) Test_server_sends_workspace_folders_request() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	_, err := dispatcher.WorkspaceFolders()
	s.Require().NoError(err)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()

	// Verify that the client received the configuration message.
	s.Require().Len(container.clientReceivedMessages, 1)

	// Verify the method name.
	s.Require().Len(container.clientReceivedMethods, 1)
	s.Require().Equal(MethodWorkspaceFolders, container.clientReceivedMethods[0])
}

func (s *DispatchTestSuite) Test_server_sends_apply_workspace_edit_request() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	editLabel := "Apply workspace edit"
	params := ApplyWorkspaceEditParams{
		Label: &editLabel,
		Edit: WorkspaceEdit{
			Changes: map[DocumentURI][]TextEdit{
				"file:///path/to/file.go": {
					{
						Range: &Range{
							Start: Position{
								Line:      10,
								Character: 5,
							},
							End: Position{
								Line:      50,
								Character: 20,
							},
						},
						NewText: "new text",
					},
				},
			},
		},
	}
	_, err := dispatcher.ApplyWorkspaceEdit(params)
	s.Require().NoError(err)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()

	// Verify that the client received the apply workspace edit message.
	s.Require().Len(container.clientReceivedMessages, 1)
	var message ApplyWorkspaceEditParams
	err = json.Unmarshal(*container.clientReceivedMessages[0], &message)
	s.Require().NoError(err)
	s.Require().Equal(params, message)

	// Verify the method name.
	s.Require().Len(container.clientReceivedMethods, 1)
	s.Require().Equal(MethodWorkspaceApplyEdit, container.clientReceivedMethods[0])
}

func (s *DispatchTestSuite) Test_server_sends_window_show_message_notification() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	showMessageParams := ShowMessageParams{
		Type:    MessageTypeInfo,
		Message: "Something interesting happened",
	}

	err := dispatcher.ShowMessageNotification(showMessageParams)
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
	var message ShowMessageParams
	err = json.Unmarshal(*container.clientReceivedMessages[0], &message)
	s.Require().NoError(err)
	s.Require().Equal(showMessageParams, message)

	// Verify the method name.
	s.Require().Len(container.clientReceivedMethods, 1)
	s.Require().Equal(MethodShowMessageNotification, container.clientReceivedMethods[0])
}

func (s *DispatchTestSuite) Test_server_sends_show_message_request() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	params := ShowMessageRequestParams{
		Type:    MessageTypeWarning,
		Message: "Something unusual happened",
		Actions: []MessageActionItem{
			{
				Title: "Retry",
			},
		},
	}
	_, err := dispatcher.ShowMessageRequest(params)
	s.Require().NoError(err)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()

	// Verify that the client received the apply workspace edit message.
	s.Require().Len(container.clientReceivedMessages, 1)
	var message ShowMessageRequestParams
	err = json.Unmarshal(*container.clientReceivedMessages[0], &message)
	s.Require().NoError(err)
	s.Require().Equal(params, message)

	// Verify the method name.
	s.Require().Len(container.clientReceivedMethods, 1)
	s.Require().Equal(MethodShowMessageRequest, container.clientReceivedMethods[0])
}

func (s *DispatchTestSuite) Test_server_sends_window_log_message_notification() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	logMessageParams := LogMessageParams{
		Type:    MessageTypeError,
		Message: "Something bad happened",
	}

	err := dispatcher.LogMessage(logMessageParams)
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
	var message LogMessageParams
	err = json.Unmarshal(*container.clientReceivedMessages[0], &message)
	s.Require().NoError(err)
	s.Require().Equal(logMessageParams, message)

	// Verify the method name.
	s.Require().Len(container.clientReceivedMethods, 1)
	s.Require().Equal(MethodLogMessage, container.clientReceivedMethods[0])
}

func (s *DispatchTestSuite) Test_server_sends_create_work_done_progress_request() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	token := "progress-test-token"
	params := WorkDoneProgressCreateParams{
		Token: &ProgressToken{
			StrVal: &token,
		},
	}
	err := dispatcher.CreateWorkDoneProgress(params)
	s.Require().NoError(err)

	// Acquire a lock on the received message list shared between goroutines.
	container.mu.Lock()
	defer container.mu.Unlock()

	// Verify that the client received the apply workspace edit message.
	s.Require().Len(container.clientReceivedMessages, 1)
	var message WorkDoneProgressCreateParams
	err = json.Unmarshal(*container.clientReceivedMessages[0], &message)
	s.Require().NoError(err)
	s.Require().Equal(params, message)

	// Verify the method name.
	s.Require().Len(container.clientReceivedMethods, 1)
	s.Require().Equal(MethodWorkDoneProgressCreate, container.clientReceivedMethods[0])
}

func (s *DispatchTestSuite) Test_server_sends_window_telemetry_event_notification() {
	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	serverHandler := newTestServerHandler()
	container := createTestConnectionsContainer(serverHandler)

	lspCtx := server.NewLSPContext(ctx, container.serverConn, nil)
	dispatcher := NewDispatcher(lspCtx)

	telemetryEventParams := map[string]interface{}{
		"traceId": "test-trace-id",
		"message": "Something interesting happened",
	}

	err := dispatcher.Telemetry(telemetryEventParams)
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
	var message map[string]interface{}
	err = json.Unmarshal(*container.clientReceivedMessages[0], &message)
	s.Require().NoError(err)
	s.Require().Equal(telemetryEventParams, message)

	// Verify the method name.
	s.Require().Len(container.clientReceivedMethods, 1)
	s.Require().Equal(MethodTelemetryEvent, container.clientReceivedMethods[0])
}

func TestDispatchTestSuite(t *testing.T) {
	suite.Run(t, new(DispatchTestSuite))
}
