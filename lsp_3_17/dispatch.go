package lsp

import "github.com/two-hundred/ls-builder/common"

// Dispatcher provides a convenient way to dispatch
// requests and notification to the client with types
// for known requests and notifications that a server can
// send to a client in LSP 3.17.0.
type Dispatcher struct {
	ctx *common.LSPContext
}

// NewDispatcher creates a new instance of a dispatcher.
// A dispatcher is wrapped around a current LSP context
// and should live as long as the underlying LSP context.
func NewDispatcher(ctx *common.LSPContext) *Dispatcher {
	return &Dispatcher{ctx: ctx}
}

// Context returns the underlying LSP context.
func (d *Dispatcher) Context() *common.LSPContext {
	return d.ctx
}

// Progress notifies the client of progress for a specific task.
func (d *Dispatcher) Progress(params ProgressParams) error {
	return d.ctx.Notify(MethodProgress, params)
}

// CancelRequest cancels a request with the client.
func (d *Dispatcher) CancelRequest(params CancelParams) error {
	return d.ctx.Notify(MethodCancelRequest, params)
}

// RegisterCapability registers a new capability with the client.
func (d *Dispatcher) RegisterCapability(params RegistrationParams) error {
	return d.ctx.Call(ClientRegisterCapability, params, nil)
}

// UnregisterCapability de-registers a capability with the client.
func (d *Dispatcher) UnregisterCapability(params UnregistrationParams) error {
	return d.ctx.Call(ClientUnregisterCapability, params, nil)
}

// LogTrace sends a notification to log the trace of the server’s execution.
// The amount of content of these notifications depends on the current
// `trace` configuration. If `trace`, the server should not send any `logTrace`
// notification. If `trace` is `messages`, the server should not add the `verbose`
// field in the `LogTraceParams`.
func (d *Dispatcher) LogTrace(params LogTraceParams) error {
	return d.ctx.Notify(MethodLogTrace, params)
}

// CodeLensRefresh requests the client to refresh all code lenses.
func (d *Dispatcher) CodeLensRefresh() error {
	return d.ctx.Call(MethodCodeLensRefresh, nil, nil)
}

// SemanticTokensRefresh requests the client to refresh the editors for which
// this server provides semantic tokens.
// As a result the client should ask the server to recompute the semantic tokens for these editors.
// This is useful if a server detects a project wide configuration change which requires a
// re-calculation of all semantic tokens. Note that the client still has the freedom
// to delay the re-calculation of the semantic tokens if for example an editor is currently not visible.
func (d *Dispatcher) SemanticTokensRefresh() error {
	return d.ctx.Call(MethodSemanticTokensRefresh, nil, nil)
}

// InlayHintRefresh requests the client to refresh inlay hints
// currently shown in editors.
func (d *Dispatcher) InlayHintRefresh() error {
	return d.ctx.Call(MethodInlayHintRefresh, nil, nil)
}

// InlineValueRefresh requests the client to refresh inline values
// currently shown in editors.
func (d *Dispatcher) InlineValueRefresh() error {
	return d.ctx.Call(MethodInlineValueRefresh, nil, nil)
}

// PublishDiagnostics sends diagnostics from the server to the client to signal
// results of validation runs.
func (d *Dispatcher) PublishDiagnostics(params PublishDiagnosticsParams) error {
	return d.ctx.Notify(MethodPublishDiagnostics, params)
}

// DiagnosticsRefresh requests that the client refreshes all needed document
// and workspace diagnostics. This is useful if the server detects a project wide
// configuration change which requires a re-calculation of all diagnostics.
func (d *Dispatcher) DiagnosticsRefresh() error {
	return d.ctx.Call(MethodDiagnosticsRefresh, nil, nil)
}

// WorkspaceConfiguration requests the client to fetch the configuration settings
// for the given scopes and configuration sections within a workspace.
func (d *Dispatcher) WorkspaceConfiguration(params ConfigurationParams) ([]LSPAny, error) {
	var result []LSPAny
	err := d.ctx.Call(MethodWorkspaceConfiguration, params, &result)
	return result, err
}

// WorkspaceFolders requests that the client fetches the workspace folders
// that are currently open.
func (d *Dispatcher) WorkspaceFolders() ([]WorkspaceFolder, error) {
	var result []WorkspaceFolder
	err := d.ctx.Call(MethodWorkspaceFolders, nil, &result)
	return result, err
}

// ApplyWorkspaceEdit requests that the client applies a workspace edit.
func (d *Dispatcher) ApplyWorkspaceEdit(params ApplyWorkspaceEditParams) (*ApplyWorkspaceEditResult, error) {
	var result ApplyWorkspaceEditResult
	err := d.ctx.Call(MethodWorkspaceApplyEdit, params, &result)
	return &result, err
}

// ShowMessageNotification sends a notification to the client to show a message
// without waiting for a response.
func (d *Dispatcher) ShowMessageNotification(params ShowMessageParams) error {
	return d.ctx.Notify(MethodShowMessageNotification, params)
}

// ShowMessageRequest sends a request to the client to show a message where the request
// can pass actions and wait for an answer from the client.
func (d *Dispatcher) ShowMessageRequest(params ShowMessageRequestParams) (*MessageActionItem, error) {
	var result MessageActionItem
	err := d.ctx.Call(MethodShowMessageRequest, params, &result)
	return &result, err
}

// LogMessage sends a notification to the client to log a message.
func (d *Dispatcher) LogMessage(params LogMessageParams) error {
	return d.ctx.Notify(MethodLogMessage, params)
}
