package lsp

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/two-hundred/ls-builder/common"
)

// Handler defines a set of message handlers that allows the server
// to respond to client notifications and requests.
// Server capabilities can be derived from the handlers defined here.
// Instances of handlers should be created with the `NewHandler` function with
// message handlers provided as options or set after creation using `SetHandler`.
type Handler struct {
	// Base Protocol
	cancelRequest CancelRequestHandlerFunc
	progress      ProgressHandlerFunc

	// Lifecycle Messages
	initialize InitializeHandlerFunc
	// initialized InitializedHandlerFunc
	// shutdown    ShutdownHandlerFunc
	// exit        ExitHandlerFunc

	// Trace Messages
	// setTrace SetTraceHandlerFunc

	// Text Document Synchronization
	// textDocumentDidOpen TextDocumentDidOpenHandlerFunc

	isInitialized bool
	// Provides a mapping of method names to the respective handlers
	// that are wrappers around the user-provided handler functions that will unmarshal params
	// and optionally set some state before calling the user-provided handler.
	messageHandlers map[string]common.Handler
	mu              sync.Mutex
}

// HandlerOption is a function that can be used to configure a handler
// with options such as message handlers.
type HandlerOption func(*Handler)

// WithCancelRequestHandler sets the handler for the `$/cancelRequest` notification.
func WithCancelRequestHandler(handler CancelRequestHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetCancelRequestHandler(handler)
	}
}

// WithProgressHandler sets the handler for the `$/progress` notification.
func WithProgressHandler(handler ProgressHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetProgressHandler(handler)
	}
}

// WithInitializeHandler sets the handler for the `initialize` request.
func WithInitializeHandler(handler InitializeHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetInitializeHandler(handler)
	}
}

// NewHandler creates a new instance of a handler, optionally,
// with a provided set of method handlers.
func NewHandler(opts ...HandlerOption) *Handler {
	h := &Handler{
		messageHandlers: make(map[string]common.Handler),
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// SetCancelRequestHandler sets the handler for the `$/cancelRequest` notification.
func (h *Handler) SetCancelRequestHandler(handler CancelRequestHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.cancelRequest = handler
	h.messageHandlers[MethodCancelRequest] = createCancelRequestHandler(h)
}

// SetProgressHandler sets the handler for the `$/progress` notification.
func (h *Handler) SetProgressHandler(handler ProgressHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.progress = handler
	h.messageHandlers[MethodProgress] = createProgressRequestHandler(h)
}

// SetInitializeHandler sets the handler for the `initialize` request.
func (h *Handler) SetInitializeHandler(handler InitializeHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.initialize = handler
	h.messageHandlers[MethodInitialize] = createInitializeHandler(h)
}

// Fulfils the common.Handler interface.
func (h *Handler) Handle(ctx *common.LSPContext) (r any, validMethod bool, validParams bool, err error) {
	if !h.IsInitialized() && ctx.Method != MethodInitialize {
		return nil, true, true, fmt.Errorf("server is not initialized")
	}

	messageHandler, hasHandler := h.messageHandlers[ctx.Method]
	if hasHandler {
		return messageHandler.Handle(ctx)
	}

	return
}

// IsInitialized returns whether or not the connection to the client
// has been initialized as per "Lifecycle Messages" of the LSP specification.
func (h *Handler) IsInitialized() bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.isInitialized
}

// SetInitialized sets the initialized state of the connection to the client
// as per "Lifecycle Messages" of the LSP specification.
func (h *Handler) SetInitialized(initialized bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.isInitialized = initialized
}

func createCancelRequestHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.cancelRequest != nil {
				validMethod = true
				var params CancelParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.cancelRequest(ctx, &params)
				}
			}
			return
		},
	)
}

func createProgressRequestHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.progress != nil {
				validMethod = true
				var params ProgressParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.progress(ctx, &params)
				}
			}
			return
		},
	)
}

func createInitializeHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.initialize != nil {
				validMethod = true
				var params InitializeParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					if r, err = root.initialize(ctx, &params); err == nil {
						root.SetInitialized(true)
					}
				}
			}
			return
		},
	)
}
