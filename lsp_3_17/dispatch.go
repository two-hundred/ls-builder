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
