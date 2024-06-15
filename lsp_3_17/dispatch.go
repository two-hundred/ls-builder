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

// Initialized notifies the client that the server has been initialized.
func (d *Dispatcher) Initialized() error {
	return d.ctx.Notify(MethodInitialized, InitializedParams{})
}

// RegisterCapability registers a new capability with the client.
func (d *Dispatcher) RegisterCapability(params RegistrationParams) error {
	return d.ctx.Call(ClientRegisterCapability, params, nil)
}
