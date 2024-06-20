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
// message handlers provided as options or set after creation using `Set*Handler`
// methods.
type Handler struct {
	// Base Protocol
	cancelRequest CancelRequestHandlerFunc
	progress      ProgressHandlerFunc

	// Lifecycle Messages
	initialize  InitializeHandlerFunc
	initialized InitializedHandlerFunc
	shutdown    ShutdownHandlerFunc
	exit        ExitHandlerFunc

	// Trace Messages
	setTrace SetTraceHandlerFunc

	// Text Document Synchronization
	textDocumentDidOpen           TextDocumentDidOpenHandlerFunc
	textDocumentDidChange         TextDocumentDidChangeHandlerFunc
	textDocumentWillSave          TextDocumentWillSaveHandlerFunc
	textDocumentWillSaveWaitUntil TextDocumentWillSaveWaitUntilHandlerFunc
	textDocumentDidSave           TextDocumentDidSaveHandlerFunc
	textDocumentDidClose          TextDocumentDidCloseHandlerFunc

	// Notebook Document Synchronisation
	notebookDocumentDidOpen   NotebookDocumentDidOpenHandlerFunc
	notebookDocumentDidChange NotebookDocumentDidChangeHandlerFunc
	notebookDocumentDidSave   NotebookDocumentDidSaveHandlerFunc
	notebookDocumentDidClose  NotebookDocumentDidCloseHandlerFunc

	// Language Features
	gotoDeclaration            GotoDeclarationHandlerFunc
	gotoDefinition             GotoDefinitionHandlerFunc
	gotoTypeDefinition         GotoTypeDefinitionHandlerFunc
	gotoImplementation         GotoImplementationHandlerFunc
	findReferences             FindReferencesHandlerFunc
	prepareCallHierarchy       PrepareCallHierarchyHandlerFunc
	callHierarchyIncomingCalls CallHierarchyIncomingCallsHandlerFunc
	callHierarchyOutgoingCalls CallHierarchyOutgoingCallsHandlerFunc
	prepareTypeHierarchy       PrepareTypeHierarchyHandlerFunc
	typeHierarchySupertypes    TypeHierarchySupertypesHandlerFunc
	typeHierarchySubtypes      TypeHierarchySubtypesHandlerFunc
	documentHighlight          DocumentHighlightHandlerFunc
	documentLink               DocumentLinkHandlerFunc
	documentLinkResolve        DocumentLinkResolveHandlerFunc
	hover                      HoverHandlerFunc
	codeLens                   CodeLensHandlerFunc

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

// WithInitializedHandler sets the handler for the `initialized` notification.
func WithInitializedHandler(handler InitializedHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetInitializedHandler(handler)
	}
}

// WithSetTraceHandler sets the handler for the `$/setTrace` notification.
func WithSetTraceHandler(handler SetTraceHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetTraceHandler(handler)
	}
}

// WithShutdownHandler sets the handler for the `shutdown` request.
func WithShutdownHandler(handler ShutdownHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetShutdownHandler(handler)
	}
}

// WithExitHandler sets the handler for the `exit` notification.
func WithExitHandler(handler ExitHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetExitHandler(handler)
	}
}

// WithTextDocumentDidOpenHandler sets the handler
// for the `textDocument/didOpen` notification.
func WithTextDocumentDidOpenHandler(handler TextDocumentDidOpenHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetTextDocumentDidOpenHandler(handler)
	}
}

// WithTextDocumentDidChangeHandler sets the handler
// for the `textDocument/didChange` notification.
func WithTextDocumentDidChangeHandler(handler TextDocumentDidChangeHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetTextDocumentDidChangeHandler(handler)
	}
}

// WithTextDocumentWillSaveHandler sets the handler
// for the `textDocument/willSave` notification.
func WithTextDocumentWillSaveHandler(handler TextDocumentWillSaveHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetTextDocumentWillSaveHandler(handler)
	}
}

// WithTextDocumentWillSaveWaitUntilHandler sets the handler
// for the `textDocument/willSaveWaitUntil` request.
func WithTextDocumentWillSaveWaitUntilHandler(handler TextDocumentWillSaveWaitUntilHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetTextDocumentWillSaveWaitUntilHandler(handler)
	}
}

// WithTextDocumentDidSaveHandler sets the handler
// for the `textDocument/didSave` notification.
func WithTextDocumentDidSaveHandler(handler TextDocumentDidSaveHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetTextDocumentDidSaveHandler(handler)
	}
}

// WithTextDocumentDidCloseHandler sets the handler
// for the `textDocument/didClose` notification.
func WithTextDocumentDidCloseHandler(handler TextDocumentDidCloseHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetTextDocumentDidCloseHandler(handler)
	}
}

// WithNotebookDocumentDidOpenHandler sets the handler
// for the `notebookDocument/didOpen` notification.
func WithNotebookDocumentDidOpenHandler(handler NotebookDocumentDidOpenHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetNotebookDocumentDidOpenHandler(handler)
	}
}

// WithNotebookDocumentDidChangeHandler sets the handler
// for the `notebookDocument/didChange` notification.
func WithNotebookDocumentDidChangeHandler(handler NotebookDocumentDidChangeHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetNotebookDocumentDidChangeHandler(handler)
	}
}

// WithNotebookDocumentDidSaveHandler sets the handler
// for the `notebookDocument/didSave` notification.
func WithNotebookDocumentDidSaveHandler(handler NotebookDocumentDidSaveHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetNotebookDocumentDidSaveHandler(handler)
	}
}

// WithNotebookDocumentDidCloseHandler sets the handler
// for the `notebookDocument/didClose` notification.
func WithNotebookDocumentDidCloseHandler(handler NotebookDocumentDidCloseHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetNotebookDocumentDidCloseHandler(handler)
	}
}

// WithGotoDeclarationHandler sets the handler for the `textDocument/declaration` request.
func WithGotoDeclarationHandler(handler GotoDeclarationHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetGotoDeclarationHandler(handler)
	}
}

// WithGotoDefinitionHandler sets the handler for the `textDocument/definition` request.
func WithGotoDefinitionHandler(handler GotoDefinitionHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetGotoDefinitionHandler(handler)
	}
}

// WithGotoTypeDefinitionHandler sets the handler for the `textDocument/typeDefinition` request.
func WithGotoTypeDefinitionHandler(handler GotoTypeDefinitionHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetGotoTypeDefinitionHandler(handler)
	}
}

// WithGotoImplementationHandler sets the handler for the `textDocument/implementation` request.
func WithGotoImplementationHandler(handler GotoImplementationHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetGotoImplementationHandler(handler)
	}
}

// WithFindReferencesHandler sets the handler for the `textDocument/references` request.
func WithFindReferencesHandler(handler FindReferencesHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetFindReferencesHandler(handler)
	}
}

// WithPrepareCallHierarchyHandler sets the handler for the `textDocument/prepareCallHierarchy` request.
func WithPrepareCallHierarchyHandler(handler PrepareCallHierarchyHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetPrepareCallHierarchyHandler(handler)
	}
}

// WithCallHierarchyIncomingCallsHandler sets the handler for the `textDocument/incomingCalls` request.
func WithCallHierarchyIncomingCallsHandler(handler CallHierarchyIncomingCallsHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetCallHierarchyIncomingCallsHandler(handler)
	}
}

// WithCallHierarchyOutgoingCallsHandler sets the handler for the `textDocument/outgoingCalls` request.
func WithCallHierarchyOutgoingCallsHandler(handler CallHierarchyOutgoingCallsHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetCallHierarchyOutgoingCallsHandler(handler)
	}
}

// WithPrepareTypeHierarchyHandler sets the handler for the `textDocument/prepareTypeHierarchy` request.
func WithPrepareTypeHierarchyHandler(handler PrepareTypeHierarchyHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetPrepareTypeHierarchyHandler(handler)
	}
}

// WithTypeHierarchySupertypesHandler sets the handler for the `typeHierarchy/supertypes` request.
func WithTypeHierarchySupertypesHandler(handler TypeHierarchySupertypesHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetTypeHierarchySupertypesHandler(handler)
	}
}

// WithTypeHierarchySubtypesHandler sets the handler for the `typeHierarchy/subtypes` request.
func WithTypeHierarchySubtypesHandler(handler TypeHierarchySubtypesHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetTypeHierarchySubtypesHandler(handler)
	}
}

// WithDocumentHighlightHandler sets the handler for the `textDocument/documentHighlight` request.
func WithDocumentHighlightHandler(handler DocumentHighlightHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetDocumentHighlightHandler(handler)
	}
}

// WithDocumentLinkHandler sets the handler for the `textDocument/documentLink` request.
func WithDocumentLinkHandler(handler DocumentLinkHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetDocumentLinkHandler(handler)
	}
}

// WithDocumentLinkResolveHandler sets the handler for the `documentLink/resolve` request.
func WithDocumentLinkResolveHandler(handler DocumentLinkResolveHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetDocumentLinkResolveHandler(handler)
	}
}

// WithHoverHandler sets the handler for the `textDocument/hover` request.
func WithHoverHandler(handler HoverHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetHoverHandler(handler)
	}
}

// WithCodeLensHandler sets the handler for the `textDocument/codeLens` request.
func WithCodeLensHandler(handler CodeLensHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetCodeLensHandler(handler)
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

// SetInitializedHandler sets the handler for the `initialized` notification.
func (h *Handler) SetInitializedHandler(handler InitializedHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.initialized = handler
	h.messageHandlers[MethodInitialized] = createInitializedHandler(h)
}

// SetTraceHandler sets the handler for the `$/setTrace` notification.
func (h *Handler) SetTraceHandler(handler SetTraceHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.setTrace = handler
	h.messageHandlers[MethodSetTrace] = createSetTraceHandler(h)
}

// SetShutdownHandler sets the handler for the `shutdown` request.
func (h *Handler) SetShutdownHandler(handler ShutdownHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.shutdown = handler
	h.messageHandlers[MethodShutdown] = createShutdownHandler(h)
}

// SetExitHandler sets the handler for the `exit` notification.
func (h *Handler) SetExitHandler(handler ExitHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.exit = handler
	h.messageHandlers[MethodExit] = createExitHandler(h)
}

// SetTextDocumentDidOpenHandler sets the handler
// for the `textDocument/didOpen` notification.
func (h *Handler) SetTextDocumentDidOpenHandler(handler TextDocumentDidOpenHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.textDocumentDidOpen = handler
	h.messageHandlers[MethodTextDocumentDidOpen] = createSetTextDocumentDidOpenHandler(h)
}

// SetTextDocumentDidChangeHandler sets the handler
// for the `textDocument/didChange` notification.
func (h *Handler) SetTextDocumentDidChangeHandler(handler TextDocumentDidChangeHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.textDocumentDidChange = handler
	h.messageHandlers[MethodTextDocumentDidChange] = createTextDocumentDidChangeHandler(h)
}

// SetTextDocumentWillSaveHandler sets the handler
// for the `textDocument/willSave` notification.
func (h *Handler) SetTextDocumentWillSaveHandler(handler TextDocumentWillSaveHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.textDocumentWillSave = handler
	h.messageHandlers[MethodTextDocumentWillSave] = createTextDocumentWillSaveHandler(h)
}

// SetTextDocumentWillSaveWaitUntilHandler sets the handler
// for the `textDocument/willSaveWaitUntil` request.
func (h *Handler) SetTextDocumentWillSaveWaitUntilHandler(handler TextDocumentWillSaveWaitUntilHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.textDocumentWillSaveWaitUntil = handler
	h.messageHandlers[MethodTextDocumentWillSaveWaitUntil] = createTextDocumentWillSaveWaitUntilHandler(h)
}

// SetTextDocumentDidSaveHandler sets the handler
// for the `textDocument/didSave` notification.
func (h *Handler) SetTextDocumentDidSaveHandler(handler TextDocumentDidSaveHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.textDocumentDidSave = handler
	h.messageHandlers[MethodTextDocumentDidSave] = createTextDocumentDidSaveHandler(h)
}

// SetTextDocumentDidCloseHandler sets the handler
// for the `textDocument/didClose` notification.
func (h *Handler) SetTextDocumentDidCloseHandler(handler TextDocumentDidCloseHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.textDocumentDidClose = handler
	h.messageHandlers[MethodTextDocumentDidClose] = createTextDocumentDidCloseHandler(h)
}

// SetNotebookDocumentDidOpenHandler sets the handler
// for the `notebookDocument/didOpen` notification.
func (h *Handler) SetNotebookDocumentDidOpenHandler(handler NotebookDocumentDidOpenHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.notebookDocumentDidOpen = handler
	h.messageHandlers[MethodNotebookDocumentDidOpen] = createNotebookDocumentDidOpenHandler(h)
}

// SetNotebookDocumentDidChangeHandler sets the handler
// for the `notebookDocument/didChange` notification.
func (h *Handler) SetNotebookDocumentDidChangeHandler(handler NotebookDocumentDidChangeHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.notebookDocumentDidChange = handler
	h.messageHandlers[MethodNotebookDocumentDidChange] = createNotebookDocumentDidChangeHandler(h)
}

// SetNotebookDocumentDidSaveHandler sets the handler
// for the `notebookDocument/didSave` notification.
func (h *Handler) SetNotebookDocumentDidSaveHandler(handler NotebookDocumentDidSaveHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.notebookDocumentDidSave = handler
	h.messageHandlers[MethodNotebookDocumentDidSave] = createNotebookDocumentDidSaveHandler(h)
}

// SetNotebookDocumentDidCloseHandler sets the handler
// for the `notebookDocument/didClose` notification.
func (h *Handler) SetNotebookDocumentDidCloseHandler(handler NotebookDocumentDidCloseHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.notebookDocumentDidClose = handler
	h.messageHandlers[MethodNotebookDocumentDidClose] = createNotebookDocumentDidCloseHandler(h)
}

// SetGotoDeclarationHandler sets the handler for the `textDocument/declaration` request.
func (h *Handler) SetGotoDeclarationHandler(handler GotoDeclarationHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.gotoDeclaration = handler
	h.messageHandlers[MethodGotoDeclaration] = createGotoDeclarationHandler(h)
}

// SetGotoDefinitionHandler sets the handler for the `textDocument/definition` request.
func (h *Handler) SetGotoDefinitionHandler(handler GotoDefinitionHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.gotoDefinition = handler
	h.messageHandlers[MethodGotoDefinition] = createGotoDefinitionHandler(h)
}

// SetGotoTypeDefinitionHandler sets the handler for the `textDocument/typeDefinition` request.
func (h *Handler) SetGotoTypeDefinitionHandler(handler GotoTypeDefinitionHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.gotoTypeDefinition = handler
	h.messageHandlers[MethodGotoTypeDefinition] = createGotoTypeDefinitionHandler(h)
}

// SetGotoImplementationHandler sets the handler for the `textDocument/implementation` request.
func (h *Handler) SetGotoImplementationHandler(handler GotoImplementationHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.gotoImplementation = handler
	h.messageHandlers[MethodGotoImplementation] = createGotoImplementationHandler(h)
}

// SetFindReferencesHandler sets the handler for the `textDocument/references` request.
func (h *Handler) SetFindReferencesHandler(handler FindReferencesHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.findReferences = handler
	h.messageHandlers[MethodFindReferences] = createFindReferencesHandler(h)
}

// SetPrepareCallHierarchyHandler sets the handler for the `textDocument/prepareCallHierarchy` request.
func (h *Handler) SetPrepareCallHierarchyHandler(handler PrepareCallHierarchyHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.prepareCallHierarchy = handler
	h.messageHandlers[MethodPrepareCallHierarchy] = createPrepareCallHierarchyHandler(h)
}

// SetCallHierarchyIncomingCallsHandler sets the handler for the `textDocument/incomingCalls` request.
func (h *Handler) SetCallHierarchyIncomingCallsHandler(handler CallHierarchyIncomingCallsHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.callHierarchyIncomingCalls = handler
	h.messageHandlers[MethodCallHierarchyIncomingCalls] = createCallHierarchyIncomingCallsHandler(h)
}

// SetCallHierarchyOutgoingCallsHandler sets the handler for the `textDocument/outgoingCalls` request.
func (h *Handler) SetCallHierarchyOutgoingCallsHandler(handler CallHierarchyOutgoingCallsHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.callHierarchyOutgoingCalls = handler
	h.messageHandlers[MethodCallHierarchyOutgoingCalls] = createCallHierarchyOutgoingCallsHandler(h)
}

// SetPrepareTypeHierarchyHandler sets the handler for the `textDocument/prepareTypeHierarchy` request.
func (h *Handler) SetPrepareTypeHierarchyHandler(handler PrepareTypeHierarchyHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.prepareTypeHierarchy = handler
	h.messageHandlers[MethodPrepareTypeHierarchy] = createPrepareTypeHierarchyHandler(h)
}

// SetTypeHierarchySupertypesHandler sets the handler for the `typeHierarchy/supertypes` request.
func (h *Handler) SetTypeHierarchySupertypesHandler(handler TypeHierarchySupertypesHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.typeHierarchySupertypes = handler
	h.messageHandlers[MethodTypeHierarchySupertypes] = createTypeHierarchySupertypesHandler(h)
}

// SetTypeHierarchySubtypesHandler sets the handler for the `typeHierarchy/subtypes` request.
func (h *Handler) SetTypeHierarchySubtypesHandler(handler TypeHierarchySubtypesHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.typeHierarchySubtypes = handler
	h.messageHandlers[MethodTypeHierarchySubtypes] = createTypeHierarchySubtypesHandler(h)
}

// SetDocumentHighlightHandler sets the handler for the `textDocument/documentHighlight` request.
func (h *Handler) SetDocumentHighlightHandler(handler DocumentHighlightHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.documentHighlight = handler
	h.messageHandlers[MethodDocumentHighlight] = createDocumentHighlightHandler(h)
}

// SetDocumentLinkHandler sets the handler for the `textDocument/documentLink` request.
func (h *Handler) SetDocumentLinkHandler(handler DocumentLinkHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.documentLink = handler
	h.messageHandlers[MethodDocumentLink] = createDocumentLinkHandler(h)
}

// SetDocumentLinkResolveHandler sets the handler for the `documentLink/resolve` request.
func (h *Handler) SetDocumentLinkResolveHandler(handler DocumentLinkResolveHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.documentLinkResolve = handler
	h.messageHandlers[MethodDocumentLinkResolve] = createDocumentLinkResolveHandler(h)
}

// SetHoverHandler sets the handler for the `textDocument/hover` request.
func (h *Handler) SetHoverHandler(handler HoverHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.hover = handler
	h.messageHandlers[MethodHover] = createHoverHandler(h)
}

// SetCodeLensHandler sets the handler for the `textDocument/codeLens` request.
func (h *Handler) SetCodeLensHandler(handler CodeLensHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.codeLens = handler
	h.messageHandlers[MethodCodeLens] = createCodeLensHandler(h)
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

func createInitializedHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.initialized != nil {
				validMethod = true
				var params InitializedParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.initialized(ctx, &params)
				}
			}
			return
		},
	)
}

func createSetTraceHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.setTrace != nil {
				validMethod = true
				var params SetTraceParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.setTrace(ctx, &params)
				}
			}
			return
		},
	)
}

func createShutdownHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			root.SetInitialized(false)
			if root.shutdown != nil {
				validMethod = true
				validParams = true
				err = root.shutdown(ctx)
			}
			return
		},
	)
}

func createExitHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			// Note that the server will close the
			// connection after we handle it here.
			if root.exit != nil {
				validMethod = true
				validParams = true
				err = root.exit(ctx)
			}
			return
		},
	)
}

func createSetTextDocumentDidOpenHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.textDocumentDidOpen != nil {
				validMethod = true
				var params DidOpenTextDocumentParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.textDocumentDidOpen(ctx, &params)
				}
			}
			return
		},
	)
}

func createTextDocumentDidChangeHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.textDocumentDidChange != nil {
				validMethod = true
				var params DidChangeTextDocumentParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.textDocumentDidChange(ctx, &params)
				}
			}
			return
		},
	)
}

func createTextDocumentWillSaveHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.textDocumentWillSave != nil {
				validMethod = true
				var params WillSaveTextDocumentParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.textDocumentWillSave(ctx, &params)
				}
			}
			return
		},
	)
}

func createTextDocumentWillSaveWaitUntilHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.textDocumentWillSaveWaitUntil != nil {
				validMethod = true
				var params WillSaveTextDocumentParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.textDocumentWillSaveWaitUntil(ctx, &params)
				}
			}
			return
		},
	)
}

func createTextDocumentDidSaveHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.textDocumentDidSave != nil {
				validMethod = true
				var params DidSaveTextDocumentParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.textDocumentDidSave(ctx, &params)
				}
			}
			return
		},
	)
}

func createTextDocumentDidCloseHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.textDocumentDidClose != nil {
				validMethod = true
				var params DidCloseTextDocumentParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.textDocumentDidClose(ctx, &params)
				}
			}
			return
		},
	)
}

func createNotebookDocumentDidOpenHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.notebookDocumentDidOpen != nil {
				validMethod = true
				var params DidOpenNotebookDocumentParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.notebookDocumentDidOpen(ctx, &params)
				}
			}
			return
		},
	)
}

func createNotebookDocumentDidChangeHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.notebookDocumentDidChange != nil {
				validMethod = true
				var params DidChangeNotebookDocumentParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.notebookDocumentDidChange(ctx, &params)
				}
			}
			return
		},
	)
}

func createNotebookDocumentDidSaveHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.notebookDocumentDidSave != nil {
				validMethod = true
				var params DidSaveNotebookDocumentParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.notebookDocumentDidSave(ctx, &params)
				}
			}
			return
		},
	)
}

func createNotebookDocumentDidCloseHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.notebookDocumentDidClose != nil {
				validMethod = true
				var params DidCloseNotebookDocumentParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					err = root.notebookDocumentDidClose(ctx, &params)
				}
			}
			return
		},
	)
}

func createGotoDeclarationHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.gotoDeclaration != nil {
				validMethod = true
				var params DeclarationParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.gotoDeclaration(ctx, &params)
				}
			}
			return
		},
	)
}

func createGotoDefinitionHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.gotoDefinition != nil {
				validMethod = true
				var params DefinitionParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.gotoDefinition(ctx, &params)
				}
			}
			return
		},
	)
}

func createGotoTypeDefinitionHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.gotoTypeDefinition != nil {
				validMethod = true
				var params TypeDefinitionParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.gotoTypeDefinition(ctx, &params)
				}
			}
			return
		},
	)
}

func createGotoImplementationHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.gotoImplementation != nil {
				validMethod = true
				var params ImplementationParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.gotoImplementation(ctx, &params)
				}
			}
			return
		},
	)
}

func createFindReferencesHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.findReferences != nil {
				validMethod = true
				var params ReferencesParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.findReferences(ctx, &params)
				}
			}
			return
		},
	)
}

func createPrepareCallHierarchyHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.prepareCallHierarchy != nil {
				validMethod = true
				var params CallHierarchyPrepareParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.prepareCallHierarchy(ctx, &params)
				}
			}
			return
		},
	)
}

func createCallHierarchyIncomingCallsHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.callHierarchyIncomingCalls != nil {
				validMethod = true
				var params CallHierarchyIncomingCallsParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.callHierarchyIncomingCalls(ctx, &params)
				}
			}
			return
		},
	)
}

func createCallHierarchyOutgoingCallsHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.callHierarchyOutgoingCalls != nil {
				validMethod = true
				var params CallHierarchyOutgoingCallsParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.callHierarchyOutgoingCalls(ctx, &params)
				}
			}
			return
		},
	)
}

func createPrepareTypeHierarchyHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.prepareTypeHierarchy != nil {
				validMethod = true
				var params TypeHierarchyPrepareParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.prepareTypeHierarchy(ctx, &params)
				}
			}
			return
		},
	)
}

func createTypeHierarchySupertypesHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.typeHierarchySupertypes != nil {
				validMethod = true
				var params TypeHierarchySupertypesParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.typeHierarchySupertypes(ctx, &params)
				}
			}
			return
		},
	)
}

func createTypeHierarchySubtypesHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.typeHierarchySubtypes != nil {
				validMethod = true
				var params TypeHierarchySubtypesParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.typeHierarchySubtypes(ctx, &params)
				}
			}
			return
		},
	)
}

func createDocumentHighlightHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.documentHighlight != nil {
				validMethod = true
				var params DocumentHighlightParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.documentHighlight(ctx, &params)
				}
			}
			return
		},
	)
}

func createDocumentLinkHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.documentLink != nil {
				validMethod = true
				var params DocumentLinkParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.documentLink(ctx, &params)
				}
			}
			return
		},
	)
}

func createDocumentLinkResolveHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			if root.documentLinkResolve != nil {
				validMethod = true
				var params DocumentLink
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.documentLinkResolve(ctx, &params)
				}
			}
			return
		},
	)
}

func createHoverHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.hover != nil {
				var params HoverParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.hover(ctx, &params)
				}
			}
			return
		},
	)
}

func createCodeLensHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.codeLens != nil {
				var params CodeLensParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.codeLens(ctx, &params)
				}
			}
			return
		},
	)
}
