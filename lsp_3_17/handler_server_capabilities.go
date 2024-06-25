package lsp

// CreateServerCapabilities creates a server capabilities object
// to be sent to the client during initialization.
// This derives a base set of capabilities from the configured handlers
// that can be modified before being sent to the client.
// All handlers that are not dynamically registered must be set
// before calling this method.
//
// For notebook synchronisation events, the server capabilities
// need to be set with notebook selectors to indicate which
// notebooks should be supported.
// See: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#notebookDocument_synchronization
// (Go to the "Server Capability" section)
func (h *Handler) CreateServerCapabilities() ServerCapabilities {
	var capabilities ServerCapabilities

	h.applyTextDocumentSyncCapabilities(&capabilities)
	h.applyNotebookDocumentSyncCapabilities(&capabilities)
	h.applyLanguageFeaturesSet1Capabilities(&capabilities)
	h.applyLanguageFeaturesSet2Capabilities(&capabilities)
	h.applyWorkspaceFeaturesCapabilities(&capabilities)

	return capabilities
}

var (
	True = true
)

func (h *Handler) applyTextDocumentSyncCapabilities(capabilities *ServerCapabilities) {
	if (h.textDocumentDidOpen != nil) && (h.textDocumentDidClose != nil) {
		prepareEmptyTextDocumentSyncOptions(capabilities)
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).OpenClose = &True
	}

	if h.textDocumentDidChange != nil {
		prepareEmptyTextDocumentSyncOptions(capabilities)
		incremental := TextDocumentSyncKindIncremental
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).Change = &incremental
	}

	if h.textDocumentWillSave != nil {
		prepareEmptyTextDocumentSyncOptions(capabilities)
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).WillSave = &True
	}

	if h.textDocumentWillSaveWaitUntil != nil {
		prepareEmptyTextDocumentSyncOptions(capabilities)
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).WillSaveWaitUntil = &True
	}

	if h.textDocumentDidSave != nil {
		prepareEmptyTextDocumentSyncOptions(capabilities)
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).Save = &True
	}
}

func prepareEmptyTextDocumentSyncOptions(capabilities *ServerCapabilities) {
	if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
		capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
	}
}

func (h *Handler) applyNotebookDocumentSyncCapabilities(capabilities *ServerCapabilities) {
	if h.notebookDocumentDidSave != nil {
		prepareEmptyNotebookDocumentSyncOptions(capabilities)
		capabilities.NotebookDocumentSync.(*NotebookDocumentSyncOptions).Save = &True
	}
}

func prepareEmptyNotebookDocumentSyncOptions(capabilities *ServerCapabilities) {
	if _, ok := capabilities.NotebookDocumentSync.(*NotebookDocumentSyncOptions); !ok {
		capabilities.NotebookDocumentSync = &NotebookDocumentSyncOptions{}
	}
}

func (h *Handler) applyLanguageFeaturesSet1Capabilities(capabilities *ServerCapabilities) {
	if h.completion != nil {
		capabilities.CompletionProvider = &CompletionOptions{}
	}

	if h.hover != nil {
		capabilities.HoverProvider = true
	}

	if h.signatureHelp != nil {
		capabilities.SignatureHelpProvider = &SignatureHelpOptions{}
	}

	if h.gotoDeclaration != nil {
		capabilities.DeclarationProvider = true
	}

	if h.gotoDefinition != nil {
		capabilities.DefinitionProvider = true
	}

	if h.gotoTypeDefinition != nil {
		capabilities.TypeDefinitionProvider = true
	}

	if h.gotoImplementation != nil {
		capabilities.ImplementationProvider = true
	}

	if h.findReferences != nil {
		capabilities.ReferencesProvider = true
	}

	if h.prepareCallHierarchy != nil {
		capabilities.CallHierarchyProvider = true
	}

	if h.prepareTypeHierarchy != nil {
		capabilities.TypeHierarchyProvider = true
	}

	if h.documentHighlight != nil {
		capabilities.DocumentHighlightProvider = true
	}

	if h.documentLink != nil {
		capabilities.DocumentLinkProvider = &DocumentLinkOptions{}
	}

	if h.codeLens != nil {
		capabilities.CodeLensProvider = &CodeLensOptions{}
	}

	if h.foldingRange != nil {
		capabilities.FoldingRangeProvider = true
	}

	if h.selectionRange != nil {
		capabilities.SelectionRangeProvider = true
	}
}

func (h *Handler) applyLanguageFeaturesSet2Capabilities(capabilities *ServerCapabilities) {
	if h.documentSymbol != nil {
		capabilities.DocumentSymbolProvider = true
	}

	if h.semanticTokensFull != nil {
		prepareEmptySemanticTokensProvider(capabilities)
		capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Full = true
	}

	if h.semanticTokensFullDelta != nil {
		prepareEmptySemanticTokensProvider(capabilities)
		capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Full = SemanticDelta{
			Delta: &True,
		}
	}

	if h.semanticTokensRange != nil {
		prepareEmptySemanticTokensProvider(capabilities)
		capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Range = true
	}

	if h.inlineValue != nil {
		capabilities.InlineValueProvider = true
	}

	if h.inlayHint != nil {
		capabilities.InlayHintProvider = true
	}

	if h.moniker != nil {
		capabilities.MonikerProvider = true
	}

	if h.documentDiagnostics != nil {
		capabilities.DiagnosticProvider = true
	}

	if h.codeAction != nil {
		capabilities.CodeActionProvider = true
	}

	if h.documentColor != nil {
		capabilities.ColorProvider = true
	}

	if h.documentFormatting != nil {
		capabilities.DocumentFormattingProvider = true
	}

	if h.documentRangeFormatting != nil {
		capabilities.DocumentRangeFormattingProvider = true
	}

	if h.documentOnTypeFormatting != nil {
		capabilities.DocumentOnTypeFormattingProvider = &DocumentOnTypeFormattingOptions{}
	}

	if h.documentRename != nil {
		capabilities.RenameProvider = true
	}

	if h.documentLinkedEditingRange != nil {
		capabilities.LinkedEditingRangeProvider = true
	}
}

func prepareEmptySemanticTokensProvider(capabilities *ServerCapabilities) {
	if _, ok := capabilities.SemanticTokensProvider.(*SemanticTokensOptions); !ok {
		capabilities.SemanticTokensProvider = &SemanticTokensOptions{}
	}
}

func (h *Handler) applyWorkspaceFeaturesCapabilities(capabilities *ServerCapabilities) {
	if h.workspaceSymbol != nil {
		capabilities.WorkspaceSymbolProvider = true
	}

	if h.workspaceDidCreateFiles != nil {
		prepareEmptyWorkspaceCapabilities(capabilities)
		capabilities.Workspace.FileOperations.DidCreate = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if h.workspaceWillCreateFiles != nil {
		prepareEmptyWorkspaceCapabilities(capabilities)
		capabilities.Workspace.FileOperations.WillCreate = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if h.workspaceDidRenameFiles != nil {
		prepareEmptyWorkspaceCapabilities(capabilities)
		capabilities.Workspace.FileOperations.DidRename = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if h.workspaceWillRenameFiles != nil {
		prepareEmptyWorkspaceCapabilities(capabilities)
		capabilities.Workspace.FileOperations.WillRename = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if h.workspaceDidDeleteFiles != nil {
		prepareEmptyWorkspaceCapabilities(capabilities)
		capabilities.Workspace.FileOperations.DidDelete = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if h.workspaceWillDeleteFiles != nil {
		prepareEmptyWorkspaceCapabilities(capabilities)
		capabilities.Workspace.FileOperations.WillDelete = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if h.workspaceExecuteCommand != nil {
		capabilities.ExecuteCommandProvider = &ExecuteCommandOptions{}
	}
}

func prepareEmptyWorkspaceCapabilities(capabilities *ServerCapabilities) {
	if capabilities.Workspace == nil {
		capabilities.Workspace = &ServerWorkspaceCapabilities{
			FileOperations: &WorkspaceFileOperationServerCapabilities{},
		}
	}
}
