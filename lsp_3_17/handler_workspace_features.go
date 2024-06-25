package lsp

import (
	"encoding/json"

	"github.com/two-hundred/ls-builder/common"
)

// WithWorkspaceSymbolHandler sets the handler for the `workspace/symbol` request.
func WithWorkspaceSymbolHandler(handler WorkspaceSymbolHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetWorkspaceSymbolHandler(handler)
	}
}

// SetWorkspaceSymbolHandler sets the handler for the `workspace/symbol` request.
func (h *Handler) SetWorkspaceSymbolHandler(handler WorkspaceSymbolHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.workspaceSymbol = handler
	h.messageHandlers[MethodWorkspaceSymbol] = createWorkspaceSymbolHandler(h)
}

func createWorkspaceSymbolHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.workspaceSymbol != nil {
				var params WorkspaceSymbolParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.workspaceSymbol(ctx, &params)
				}
			}
			return
		},
	)
}
