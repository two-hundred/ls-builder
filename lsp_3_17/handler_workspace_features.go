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

// WithWorkspaceSymbolResolveHandler sets the handler for the `workspaceSymbol/resolve` request.
func WithWorkspaceSymbolResolveHandler(handler WorkspaceSymbolResolveHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetWorkspaceSymbolResolveHandler(handler)
	}
}

// WithWorkspaceDidChangeConfigurationHandler sets the handler for the `workspace/didChangeConfiguration` notification.
func WithWorkspaceDidChangeConfigurationHandler(handler WorkspaceDidChangeConfigurationHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetWorkspaceDidChangeConfigurationHandler(handler)
	}
}

// WithWorkspaceDidChangeFoldersHandler sets the handler for the `workspace/didChangeWorkspaceFolders` notification.
func WithWorkspaceDidChangeFoldersHandler(handler WorkspaceDidChangeFoldersHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetWorkspaceDidChangeFoldersHandler(handler)
	}
}

// WithWorkspaceWillCreateFilesHandler sets the handler for the `workspace/willCreateFiles` request.
func WithWorkspaceWillCreateFilesHandler(handler WorkspaceWillCreateFilesHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetWorkspaceWillCreateFilesHandler(handler)
	}
}

// WithWorkspaceDidCreateFilesHandler sets the handler for the `workspace/didCreateFiles` notification.
func WithWorkspaceDidCreateFilesHandler(handler WorkspaceDidCreateFilesHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetWorkspaceDidCreateFilesHandler(handler)
	}
}

// SetWorkspaceSymbolHandler sets the handler for the `workspace/symbol` request.
func (h *Handler) SetWorkspaceSymbolHandler(handler WorkspaceSymbolHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.workspaceSymbol = handler
	h.messageHandlers[MethodWorkspaceSymbol] = createWorkspaceSymbolHandler(h)
}

// SetWorkspaceSymbolResolveHandler sets the handler for the `workspaceSymbol/resolve` request.
func (h *Handler) SetWorkspaceSymbolResolveHandler(handler WorkspaceSymbolResolveHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.workspaceSymbolResolve = handler
	h.messageHandlers[MethodWorkspaceSymbolResolve] = createWorkspaceSymbolResolveHandler(h)
}

// SetWorkspaceDidChangeConfigurationHandler sets the handler for the `workspace/didChangeConfiguration` notification.
func (h *Handler) SetWorkspaceDidChangeConfigurationHandler(handler WorkspaceDidChangeConfigurationHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.workspaceDidChangeConfiguration = handler
	h.messageHandlers[MethodWorkspaceDidChangeConfiguration] = createWorkspaceDidChangeConfigurationHandler(h)
}

// SetWorkspaceDidChangeFoldersHandler sets the handler for the `workspace/didChangeWorkspaceFolders` notification.
func (h *Handler) SetWorkspaceDidChangeFoldersHandler(handler WorkspaceDidChangeFoldersHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.workspaceDidChangeFolders = handler
	h.messageHandlers[MethodWorkspaceDidChangeFolders] = createWorkspaceDidChangeFoldersHandler(h)
}

// SetWorkspaceWillCreateFilesHandler sets the handler for the `workspace/willCreateFiles` request.
func (h *Handler) SetWorkspaceWillCreateFilesHandler(handler WorkspaceWillCreateFilesHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.workspaceWillCreateFiles = handler
	h.messageHandlers[MethodWorkspaceWillCreateFiles] = createWorkspaceWillCreateFilesHandler(h)
}

// SetWorkspaceDidCreateFilesHandler sets the handler for the `workspace/didCreateFiles` notification.
func (h *Handler) SetWorkspaceDidCreateFilesHandler(handler WorkspaceDidCreateFilesHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.workspaceDidCreateFiles = handler
	h.messageHandlers[MethodWorkspaceDidCreateFiles] = createWorkspaceDidCreateFilesHandler(h)
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

func createWorkspaceSymbolResolveHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.workspaceSymbolResolve != nil {
				var params WorkspaceSymbol
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.workspaceSymbolResolve(ctx, &params)
				}
			}
			return
		},
	)
}

func createWorkspaceDidChangeConfigurationHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.workspaceDidChangeConfiguration != nil {
				var params DidChangeConfigurationParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.workspaceDidChangeConfiguration(ctx, &params)
				}
			}
			return
		},
	)
}

func createWorkspaceDidChangeFoldersHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.workspaceDidChangeFolders != nil {
				var params DidChangeWorkspaceFoldersParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.workspaceDidChangeFolders(ctx, &params)
				}
			}
			return
		},
	)
}

func createWorkspaceWillCreateFilesHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.workspaceWillCreateFiles != nil {
				var params CreateFilesParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.workspaceWillCreateFiles(ctx, &params)
				}
			}
			return
		},
	)
}

func createWorkspaceDidCreateFilesHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.workspaceDidCreateFiles != nil {
				var params CreateFilesParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.workspaceDidCreateFiles(ctx, &params)
				}
			}
			return
		},
	)
}
