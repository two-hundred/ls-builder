package server

import (
	"context"

	"github.com/sourcegraph/jsonrpc2"
	"github.com/two-hundred/ls-builder/common"
)

// NewLSPContext creates a new LSP context from the given connection and request.
func NewLSPContext(ctx context.Context, conn *jsonrpc2.Conn, request *jsonrpc2.Request) *common.LSPContext {
	lspContext := &common.LSPContext{
		Notify: func(method string, params any) error {
			return conn.Notify(ctx, method, params)
		},
		Call: func(method string, params any, result any) error {
			return conn.Call(ctx, method, params, result)
		},
	}

	if request == nil {
		return lspContext
	}

	lspContext.Method = request.Method
	if request.Params != nil {
		lspContext.Params = *request.Params
	}

	return lspContext
}
