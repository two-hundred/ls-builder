package common

import (
	"context"
	"encoding/json"
)

// NotifyFunc is a signature for a function that sends JSON-RPC notifications.
// This should only be used directly in the case where lsp.Dispatcher doesn't implement
// the method you need to call.
type NotifyFunc func(method string, params any) error

// CallFunc is a signature for a function that sends JSON-RPC requests.
// The result parameter should be a pointer to the value that the result should be unmarshaled into.
// This should only be used directly in the case where lsp.Dispatcher doesn't implement
// the method you need to call.
type CallFunc func(method string, params any, result any) error

// LSPContext contains the context for an LSP request from a client.
type LSPContext struct {
	Method  string
	Params  json.RawMessage
	Notify  NotifyFunc
	Call    CallFunc
	Context context.Context
}

// Handler provides an interface for handling LSP requests.
type Handler interface {
	Handle(ctx *LSPContext) (r any, validMethod bool, validParams bool, err error)
}

// HandlerFunc provides a convenient way to define a handler with a plain function.
type HandlerFunc func(ctx *LSPContext) (r any, validMethod bool, validParams bool, err error)

func (f HandlerFunc) Handle(ctx *LSPContext) (r any, validMethod bool, validParams bool, err error) {
	return f(ctx)
}
