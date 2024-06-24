package lsp

import (
	"encoding/json"

	"github.com/two-hundred/ls-builder/common"
)

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

// WithCodeLensResolveHandler sets the handler for the `codeLens/resolve` request.
func WithCodeLensResolveHandler(handler CodeLensResolveHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetCodeLensResolveHandler(handler)
	}
}

// WithFoldingRangeHandler sets the handler for the `textDocument/foldingRange` request.
func WithFoldingRangeHandler(handler FoldingRangeHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetFoldingRangeHandler(handler)
	}
}

// WithSelectionRangeHandler sets the handler for the `textDocument/selectionRange` request.
func WithSelectionRangeHandler(handler SelectionRangeHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetSelectionRangeHandler(handler)
	}
}

// WithDocumentSymbolHandler sets the handler for the `textDocument/documentSymbol` request.
func WithDocumentSymbolHandler(handler DocumentSymbolHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetDocumentSymbolHandler(handler)
	}
}

// WithSemanticTokensFullHandler sets the handler for the `textDocument/semanticTokens/full` request.
func WithSemanticTokensFullHandler(handler SemanticTokensFullHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetSemanticTokensFullHandler(handler)
	}
}

// WithSemanticTokensFullDeltaHandler sets the handler for the `textDocument/semanticTokens/full/delta` request.
func WithSemanticTokensFullDeltaHandler(handler SemanticTokensFullDeltaHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetSemanticTokensFullDeltaHandler(handler)
	}
}

// WithSemanticTokensRangeHandler sets the handler for the `textDocument/semanticTokens/range` request.
func WithSemanticTokensRangeHandler(handler SemanticTokensRangeHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetSemanticTokensRangeHandler(handler)
	}
}

// WithInlayHintHandler sets the handler for the `textDocument/inlayHint` request.
func WithInlayHintHandler(handler InlayHintHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetInlayHintHandler(handler)
	}
}

// WithInlayHintResolveHandler sets the handler for the `inlayHint/resolve` request.
func WithInlayHintResolveHandler(handler InlayHintResolveHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetInlayHintResolveHandler(handler)
	}
}

// WithInlineValueHandler sets the handler for the `inlineValue` request.
func WithInlineValueHandler(handler InlineValueHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetInlineValueHandler(handler)
	}
}

// WithMonikerHandler sets the handler for the `textDocument/moniker` request.
func WithMonikerHandler(handler MonikerHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetMonikerHandler(handler)
	}
}

// WithCompletionHandler sets teh handler for the `textDocument/completion` request.
func WithCompletionHandler(handler CompletionHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetCompletionHandler(handler)
	}
}

// WithCompletionItemResolveHandler sets the handler for the `completionItem/resolve` request.
func WithCompletionItemResolveHandler(handler CompletionItemResolveHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetCompletionItemResolveHandler(handler)
	}
}

// WithDocumentDiagnosticsHandler sets the handler for the `textDocument/diagnostics` request.
func WithDocumentDiagnosticsHandler(handler DocumentDiagnosticHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetDocumentDiagnosticsHandler(handler)
	}
}

// WithWorkspaceDiagnosticHandler sets the handler for the `workspace/diagnostics` request.
func WithWorkspaceDiagnosticHandler(handler WorkspaceDiagnosticHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetWorkspaceDiagnosticHandler(handler)
	}
}

// WithSignatureHelpHandler sets the handler for the `textDocument/signatureHelp` request.
func WithSignatureHelpHandler(handler SignatureHelpHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetSignatureHelpHandler(handler)
	}
}

// WithCodeActionHandler sets the handler for the `textDocument/codeAction` request.
func WithCodeActionHandler(handler CodeActionHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetCodeActionHandler(handler)
	}
}

// WithCodeActionResolveHandler sets the handler for the `codeAction/resolve` request.
func WithCodeActionResolveHandler(handler CodeActionResolveHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetCodeActionResolveHandler(handler)
	}
}

// WithDocumentColorHandler sets the handler for the `textDocument/documentColor` request.
func WithDocumentColorHandler(handler DocumentColorHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetDocumentColorHandler(handler)
	}
}

// WithDocumentColorPresentationHandler sets the handler for the
// `textDocument/colorPresentation` request.
func WithDocumentColorPresentationHandler(handler DocumentColorPresentationHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetDocumentColorPresentationHandler(handler)
	}
}

// WithDocumentFormattingHandler sets the handler for the
// `textDocument/formatting` request.
func WithDocumentFormattingHandler(handler DocumentFormattingHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetDocumentFormattingHandler(handler)
	}
}

// WithDocumentRangeFormattingHandler sets the handler for the
// `textDocument/rangeFormatting` request.
func WithDocumentRangeFormattingHandler(handler DocumentRangeFormattingHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetDocumentRangeFormattingHandler(handler)
	}
}

// WithDocumentOnTypeFormattingHandler sets the handler for the
// `textDocument/onTypeFormatting` request.
func WithDocumentOnTypeFormattingHandler(handler DocumentOnTypeFormattingHandlerFunc) HandlerOption {
	return func(root *Handler) {
		root.SetDocumentOnTypeFormattingHandler(handler)
	}
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

// SetCodeLensResolveHandler sets the handler for the `codeLens/resolve` request.
func (h *Handler) SetCodeLensResolveHandler(handler CodeLensResolveHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.codelensResolve = handler
	h.messageHandlers[MethodCodeLensResolve] = createCodeLensResolveHandler(h)
}

// SetFoldingRangeHandler sets the handler for the `textDocument/foldingRange` request.
func (h *Handler) SetFoldingRangeHandler(handler FoldingRangeHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.foldingRange = handler
	h.messageHandlers[MethodFoldingRange] = createFoldingRangeHandler(h)
}

// SetSelectionRangeHandler sets the handler for the `textDocument/selectionRange` request.
func (h *Handler) SetSelectionRangeHandler(handler SelectionRangeHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.selectionRange = handler
	h.messageHandlers[MethodSelectionRange] = createSelectionRangeHandler(h)
}

// SetDocumentSymbolHandler sets the handler for the `textDocument/documentSymbol` request.
func (h *Handler) SetDocumentSymbolHandler(handler DocumentSymbolHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.documentSymbol = handler
	h.messageHandlers[MethodDocumentSymbol] = createDocumentSymbolHandler(h)
}

// SetSemanticTokensFullHandler sets the handler for the `textDocument/semanticTokens/full` request.
func (h *Handler) SetSemanticTokensFullHandler(handler SemanticTokensFullHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.semanticTokensFull = handler
	h.messageHandlers[MethodSemanticTokensFull] = createSemanticTokensFullHandler(h)
}

// SetSemanticTokensFullDeltaHandler sets the handler for the `textDocument/semanticTokens/full/delta` request.
func (h *Handler) SetSemanticTokensFullDeltaHandler(handler SemanticTokensFullDeltaHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.semanticTokensFullDelta = handler
	h.messageHandlers[MethodSemanticTokensFullDelta] = createSemanticTokensFullDeltaHandler(h)
}

// SetSemanticTokensRangeHandler sets the handler for the `textDocument/semanticTokens/range` request.
func (h *Handler) SetSemanticTokensRangeHandler(handler SemanticTokensRangeHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.semanticTokensRange = handler
	h.messageHandlers[MethodSemanticTokensRange] = createSemanticTokensRangeHandler(h)
}

// SetInlayHintHandler sets the handler for the `textDocument/inlayHint` request.
func (h *Handler) SetInlayHintHandler(handler InlayHintHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.inlayHint = handler
	h.messageHandlers[MethodInlayHint] = createInlayHintHandler(h)
}

// SetInlayHintResolveHandler sets the handler for the `inlayHint/resolve` request.
func (h *Handler) SetInlayHintResolveHandler(handler InlayHintResolveHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.inlayHintResolve = handler
	h.messageHandlers[MethodInlayHintResolve] = createInlayHintResolveHandler(h)
}

// SetInlineValueHandler sets the handler for the `inlineValue` request.
func (h *Handler) SetInlineValueHandler(handler InlineValueHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.inlineValue = handler
	h.messageHandlers[MethodInlineValue] = createInlineValueHandler(h)
}

// SetMonikerHandler sets the handler for the `textDocument/moniker` request.
func (h *Handler) SetMonikerHandler(handler MonikerHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.moniker = handler
	h.messageHandlers[MethodMoniker] = createMonikerHandler(h)
}

// SetCompletionHandler sets the handler for the `textDocument/completion` request.
func (h *Handler) SetCompletionHandler(handler CompletionHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.completion = handler
	h.messageHandlers[MethodCompletion] = createCompletionHandler(h)
}

// SetCompletionItemResolveHandler sets the handler for the `completionItem/resolve` request.
func (h *Handler) SetCompletionItemResolveHandler(handler CompletionItemResolveHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.completionItemResolve = handler
	h.messageHandlers[MethodCompletionItemResolve] = createCompletionItemResolveHandler(h)
}

// SetDocumentDiagnosticsHandler sets the handler for the `textDocument/diagnostics` request.
func (h *Handler) SetDocumentDiagnosticsHandler(handler DocumentDiagnosticHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.documentDiagnostics = handler
	h.messageHandlers[MethodDocumentDiagnostic] = createDocumentDiagnosticsHandler(h)
}

// SetWorkspaceDiagnosticHandler sets the handler for the `workspace/diagnostics` request.
func (h *Handler) SetWorkspaceDiagnosticHandler(handler WorkspaceDiagnosticHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.workspaceDiagnostics = handler
	h.messageHandlers[MethodWorkspaceDiagnostic] = createWorkspaceDiagnosticHandler(h)
}

// SetSignatureHelpHandler sets the handler for the `textDocument/signatureHelp` request.
func (h *Handler) SetSignatureHelpHandler(handler SignatureHelpHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.signatureHelp = handler
	h.messageHandlers[MethodSignatureHelp] = createSignatureHelpHandler(h)
}

// SetCodeActionHandler sets the handler for the `textDocument/codeAction` request.
func (h *Handler) SetCodeActionHandler(handler CodeActionHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.codeAction = handler
	h.messageHandlers[MethodCodeAction] = createCodeActionHandler(h)
}

// SetCodeActionResolveHandler sets the handler for the `codeAction/resolve` request.
func (h *Handler) SetCodeActionResolveHandler(handler CodeActionResolveHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.codeActionResolve = handler
	h.messageHandlers[MethodCodeActionResolve] = createCodeActionResolveHandler(h)
}

// SetDocumentColorHandler sets the handler for the `textDocument/documentColor` request.
func (h *Handler) SetDocumentColorHandler(handler DocumentColorHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.documentColor = handler
	h.messageHandlers[MethodDocumentColor] = createDocumentColorHandler(h)
}

// SetDocumentColorPresentationHandler sets the handler for the `textDocument/colorPresentation` request.
func (h *Handler) SetDocumentColorPresentationHandler(handler DocumentColorPresentationHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.documentColorPresentation = handler
	h.messageHandlers[MethodDocumentColorPresentation] = createDocumentColorPresentationHandler(h)
}

// SetDocumentFormattingHandler sets the handler for the `textDocument/formatting` request.
func (h *Handler) SetDocumentFormattingHandler(handler DocumentFormattingHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.documentFormatting = handler
	h.messageHandlers[MethodDocumentFormatting] = createDocumentFormattingHandler(h)
}

// SetDocumentRangeFormattingHandler sets the handler for the `textDocument/rangeFormatting` request.
func (h *Handler) SetDocumentRangeFormattingHandler(handler DocumentRangeFormattingHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.documentRangeFormatting = handler
	h.messageHandlers[MethodDocumentRangeFormatting] = createDocumentRangeFormattingHandler(h)
}

// SetDocumentOnTypeFormattingHandler sets the handler for the `textDocument/onTypeFormatting` request.
func (h *Handler) SetDocumentOnTypeFormattingHandler(handler DocumentOnTypeFormattingHandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.documentOnTypeFormatting = handler
	h.messageHandlers[MethodDocumentOnTypeFormatting] = createDocumentOnTypeFormattingHandler(h)
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

func createCodeLensResolveHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.codelensResolve != nil {
				var params CodeLens
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.codelensResolve(ctx, &params)
				}
			}
			return
		},
	)
}

func createFoldingRangeHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.foldingRange != nil {
				var params FoldingRangeParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.foldingRange(ctx, &params)
				}
			}
			return
		},
	)
}

func createSelectionRangeHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.selectionRange != nil {
				var params SelectionRangeParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.selectionRange(ctx, &params)
				}
			}
			return
		},
	)
}

func createDocumentSymbolHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.documentSymbol != nil {
				var params DocumentSymbolParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.documentSymbol(ctx, &params)
				}
			}
			return
		},
	)
}

func createSemanticTokensFullHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.semanticTokensFull != nil {
				var params SemanticTokensParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.semanticTokensFull(ctx, &params)
				}
			}
			return
		},
	)
}

func createSemanticTokensFullDeltaHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.semanticTokensFullDelta != nil {
				var params SemanticTokensDeltaParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.semanticTokensFullDelta(ctx, &params)
				}
			}
			return
		},
	)
}

func createSemanticTokensRangeHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.semanticTokensRange != nil {
				var params SemanticTokensRangeParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.semanticTokensRange(ctx, &params)
				}
			}
			return
		},
	)
}

func createInlayHintHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.inlayHint != nil {
				var params InlayHintParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.inlayHint(ctx, &params)
				}
			}
			return
		},
	)
}

func createInlayHintResolveHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.inlayHintResolve != nil {
				var params InlayHint
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.inlayHintResolve(ctx, &params)
				}
			}
			return
		},
	)
}

func createInlineValueHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.inlineValue != nil {
				var params InlineValueParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.inlineValue(ctx, &params)
				}
			}
			return
		},
	)
}

func createMonikerHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.moniker != nil {
				var params MonikerParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.moniker(ctx, &params)
				}
			}
			return
		},
	)
}

func createCompletionHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.completion != nil {
				var params CompletionParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.completion(ctx, &params)
				}
			}
			return
		},
	)
}

func createCompletionItemResolveHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.completionItemResolve != nil {
				var params CompletionItem
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.completionItemResolve(ctx, &params)
				}
			}
			return
		},
	)
}

func createDocumentDiagnosticsHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.documentDiagnostics != nil {
				var params DocumentDiagnosticParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.documentDiagnostics(ctx, &params)
				}
			}
			return
		},
	)
}

func createWorkspaceDiagnosticHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.workspaceDiagnostics != nil {
				var params WorkspaceDiagnosticParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.workspaceDiagnostics(ctx, &params)
				}
			}
			return
		},
	)
}

func createSignatureHelpHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.signatureHelp != nil {
				var params SignatureHelpParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.signatureHelp(ctx, &params)
				}
			}
			return
		},
	)
}

func createCodeActionHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.codeAction != nil {
				var params CodeActionParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.codeAction(ctx, &params)
				}
			}
			return
		},
	)
}

func createCodeActionResolveHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.codeActionResolve != nil {
				var params CodeAction
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.codeActionResolve(ctx, &params)
				}
			}
			return
		},
	)
}

func createDocumentColorHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.documentColor != nil {
				var params DocumentColorParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.documentColor(ctx, &params)
				}
			}
			return
		},
	)
}

func createDocumentColorPresentationHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.documentColorPresentation != nil {
				var params ColorPresentationParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.documentColorPresentation(ctx, &params)
				}
			}
			return
		},
	)
}

func createDocumentFormattingHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.documentFormatting != nil {
				var params DocumentFormattingParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.documentFormatting(ctx, &params)
				}
			}
			return
		},
	)
}

func createDocumentRangeFormattingHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.documentRangeFormatting != nil {
				var params DocumentRangeFormattingParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.documentRangeFormatting(ctx, &params)
				}
			}
			return
		},
	)
}

func createDocumentOnTypeFormattingHandler(root *Handler) common.Handler {
	return common.HandlerFunc(
		func(
			ctx *common.LSPContext,
		) (r any, validMethod bool, validParams bool, err error) {
			validMethod = true
			if root.documentOnTypeFormatting != nil {
				var params DocumentOnTypeFormattingParams
				if err = json.Unmarshal(ctx.Params, &params); err == nil {
					validParams = true
					r, err = root.documentOnTypeFormatting(ctx, &params)
				}
			}
			return
		},
	)
}
