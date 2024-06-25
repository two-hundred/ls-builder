package lsp

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/two-hundred/ls-builder/common"
)

type HandlerServerCapabilitiesTestSuite struct {
	suite.Suite
}

func (s *HandlerServerCapabilitiesTestSuite) Test_create_server_capabilities() {
	h := NewHandler(
		WithTextDocumentDidOpenHandler(
			func(ctx *common.LSPContext, params *DidOpenTextDocumentParams) error {
				return nil
			},
		),
		WithTextDocumentDidCloseHandler(
			func(ctx *common.LSPContext, params *DidCloseTextDocumentParams) error {
				return nil
			},
		),
		WithTextDocumentDidChangeHandler(
			func(ctx *common.LSPContext, params *DidChangeTextDocumentParams) error {
				return nil
			},
		),
		WithTextDocumentWillSaveHandler(
			func(ctx *common.LSPContext, params *WillSaveTextDocumentParams) error {
				return nil
			},
		),
		WithTextDocumentWillSaveWaitUntilHandler(
			func(ctx *common.LSPContext, params *WillSaveTextDocumentParams) ([]TextEdit, error) {
				return nil, nil
			},
		),
		WithTextDocumentDidSaveHandler(
			func(ctx *common.LSPContext, params *DidSaveTextDocumentParams) error {
				return nil
			},
		),
		WithNotebookDocumentDidSaveHandler(
			func(ctx *common.LSPContext, params *DidSaveNotebookDocumentParams) error {
				return nil
			},
		),
		WithCompletionHandler(
			func(ctx *common.LSPContext, params *CompletionParams) (any, error) {
				return nil, nil
			},
		),
		WithHoverHandler(
			func(ctx *common.LSPContext, params *HoverParams) (*Hover, error) {
				return nil, nil
			},
		),
		WithSignatureHelpHandler(
			func(ctx *common.LSPContext, params *SignatureHelpParams) (*SignatureHelp, error) {
				return nil, nil
			},
		),
		WithGotoDeclarationHandler(
			func(ctx *common.LSPContext, params *DeclarationParams) (any, error) {
				return nil, nil
			},
		),
		WithGotoDefinitionHandler(
			func(ctx *common.LSPContext, params *DefinitionParams) (any, error) {
				return nil, nil
			},
		),
		WithGotoTypeDefinitionHandler(
			func(ctx *common.LSPContext, params *TypeDefinitionParams) (any, error) {
				return nil, nil
			},
		),
		WithGotoImplementationHandler(
			func(ctx *common.LSPContext, params *ImplementationParams) (any, error) {
				return nil, nil
			},
		),
		WithFindReferencesHandler(
			func(ctx *common.LSPContext, params *ReferencesParams) ([]Location, error) {
				return nil, nil
			},
		),
		WithPrepareCallHierarchyHandler(
			func(ctx *common.LSPContext, params *CallHierarchyPrepareParams) ([]CallHierarchyItem, error) {
				return nil, nil
			},
		),
		WithPrepareTypeHierarchyHandler(
			func(ctx *common.LSPContext, params *TypeHierarchyPrepareParams) ([]TypeHierarchyItem, error) {
				return nil, nil
			},
		),
		WithDocumentHighlightHandler(
			func(ctx *common.LSPContext, params *DocumentHighlightParams) ([]DocumentHighlight, error) {
				return nil, nil
			},
		),
		WithDocumentLinkHandler(
			func(ctx *common.LSPContext, params *DocumentLinkParams) ([]DocumentLink, error) {
				return nil, nil
			},
		),
		WithCodeLensHandler(
			func(ctx *common.LSPContext, params *CodeLensParams) ([]CodeLens, error) {
				return nil, nil
			},
		),
		WithFoldingRangeHandler(
			func(ctx *common.LSPContext, params *FoldingRangeParams) ([]FoldingRange, error) {
				return nil, nil
			},
		),
		WithSelectionRangeHandler(
			func(ctx *common.LSPContext, params *SelectionRangeParams) ([]SelectionRange, error) {
				return nil, nil
			},
		),
		WithDocumentSymbolHandler(
			func(ctx *common.LSPContext, params *DocumentSymbolParams) (any, error) {
				return nil, nil
			},
		),
		WithSemanticTokensFullHandler(
			func(ctx *common.LSPContext, params *SemanticTokensParams) (*SemanticTokens, error) {
				return nil, nil
			},
		),
		WithSemanticTokensFullDeltaHandler(
			func(ctx *common.LSPContext, params *SemanticTokensDeltaParams) (*SemanticTokensDelta, error) {
				return nil, nil
			},
		),
		WithSemanticTokensRangeHandler(
			func(ctx *common.LSPContext, params *SemanticTokensRangeParams) (*SemanticTokens, error) {
				return nil, nil
			},
		),
		WithInlineValueHandler(
			func(ctx *common.LSPContext, params *InlineValueParams) ([]*InlineValue, error) {
				return nil, nil
			},
		),
		WithInlayHintHandler(
			func(ctx *common.LSPContext, params *InlayHintParams) ([]*InlayHint, error) {
				return nil, nil
			},
		),
		WithMonikerHandler(
			func(ctx *common.LSPContext, params *MonikerParams) ([]Moniker, error) {
				return nil, nil
			},
		),
		WithDocumentDiagnosticsHandler(
			func(ctx *common.LSPContext, params *DocumentDiagnosticParams) (any, error) {
				return nil, nil
			},
		),
		WithCodeActionHandler(
			func(ctx *common.LSPContext, params *CodeActionParams) ([]*CodeActionOrCommand, error) {
				return nil, nil
			},
		),
		WithDocumentColorHandler(
			func(ctx *common.LSPContext, params *DocumentColorParams) ([]ColorInformation, error) {
				return nil, nil
			},
		),
		WithDocumentFormattingHandler(
			func(ctx *common.LSPContext, params *DocumentFormattingParams) ([]TextEdit, error) {
				return nil, nil
			},
		),
		WithDocumentRangeFormattingHandler(
			func(ctx *common.LSPContext, params *DocumentRangeFormattingParams) ([]TextEdit, error) {
				return nil, nil
			},
		),
		WithDocumentOnTypeFormattingHandler(
			func(ctx *common.LSPContext, params *DocumentOnTypeFormattingParams) ([]TextEdit, error) {
				return nil, nil
			},
		),
		WithDocumentRenameHandler(
			func(ctx *common.LSPContext, params *RenameParams) (*WorkspaceEdit, error) {
				return nil, nil
			},
		),
		WithDocumentLinkedEditingRangeHandler(
			func(ctx *common.LSPContext, params *LinkedEditingRangeParams) (*LinkedEditingRanges, error) {
				return nil, nil
			},
		),
		WithWorkspaceSymbolHandler(
			func(ctx *common.LSPContext, params *WorkspaceSymbolParams) (any, error) {
				return nil, nil
			},
		),
		WithWorkspaceDidCreateFilesHandler(
			func(ctx *common.LSPContext, params *CreateFilesParams) error {
				return nil
			},
		),
		WithWorkspaceWillCreateFilesHandler(
			func(ctx *common.LSPContext, params *CreateFilesParams) (*WorkspaceEdit, error) {
				return nil, nil
			},
		),
		WithWorkspaceDidRenameFilesHandler(
			func(ctx *common.LSPContext, params *RenameFilesParams) error {
				return nil
			},
		),
		WithWorkspaceWillRenameFilesHandler(
			func(ctx *common.LSPContext, params *RenameFilesParams) (*WorkspaceEdit, error) {
				return nil, nil
			},
		),
		WithWorkspaceDidDeleteFilesHandler(
			func(ctx *common.LSPContext, params *DeleteFilesParams) error {
				return nil
			},
		),
		WithWorkspaceWillDeleteFilesHandler(
			func(ctx *common.LSPContext, params *DeleteFilesParams) (*WorkspaceEdit, error) {
				return nil, nil
			},
		),
		WithWorkspaceExecuteCommandHandler(
			func(ctx *common.LSPContext, params *ExecuteCommandParams) (any, error) {
				return nil, nil
			},
		),
	)
	capabilities := h.CreateServerCapabilities()

	incremental := TextDocumentSyncKindIncremental
	expectedCapabilities := ServerCapabilities{
		TextDocumentSync: &TextDocumentSyncOptions{
			OpenClose:         &True,
			Change:            &incremental,
			WillSave:          &True,
			WillSaveWaitUntil: &True,
			Save:              &True,
		},
		NotebookDocumentSync: &NotebookDocumentSyncOptions{
			Save: &True,
		},
		CompletionProvider:        &CompletionOptions{},
		HoverProvider:             true,
		SignatureHelpProvider:     &SignatureHelpOptions{},
		DeclarationProvider:       true,
		DefinitionProvider:        true,
		TypeDefinitionProvider:    true,
		ImplementationProvider:    true,
		ReferencesProvider:        true,
		CallHierarchyProvider:     true,
		TypeHierarchyProvider:     true,
		DocumentHighlightProvider: true,
		DocumentLinkProvider:      &DocumentLinkOptions{},
		CodeLensProvider:          &CodeLensOptions{},
		FoldingRangeProvider:      true,
		SelectionRangeProvider:    true,
		DocumentSymbolProvider:    true,
		SemanticTokensProvider: &SemanticTokensOptions{
			Full: SemanticDelta{
				Delta: &True,
			},
			Range: true,
		},
		InlineValueProvider:              true,
		InlayHintProvider:                true,
		MonikerProvider:                  true,
		DiagnosticProvider:               true,
		CodeActionProvider:               true,
		ColorProvider:                    true,
		DocumentFormattingProvider:       true,
		DocumentRangeFormattingProvider:  true,
		DocumentOnTypeFormattingProvider: &DocumentOnTypeFormattingOptions{},
		RenameProvider:                   true,
		LinkedEditingRangeProvider:       true,
		WorkspaceSymbolProvider:          true,
		Workspace: &ServerWorkspaceCapabilities{
			FileOperations: &WorkspaceFileOperationServerCapabilities{
				DidCreate: &FileOperationRegistrationOptions{
					Filters: []FileOperationFilter{},
				},
				WillCreate: &FileOperationRegistrationOptions{
					Filters: []FileOperationFilter{},
				},
				DidRename: &FileOperationRegistrationOptions{
					Filters: []FileOperationFilter{},
				},
				WillRename: &FileOperationRegistrationOptions{
					Filters: []FileOperationFilter{},
				},
				DidDelete: &FileOperationRegistrationOptions{
					Filters: []FileOperationFilter{},
				},
				WillDelete: &FileOperationRegistrationOptions{
					Filters: []FileOperationFilter{},
				},
			},
		},
		ExecuteCommandProvider: &ExecuteCommandOptions{},
	}
	s.Require().Equal(expectedCapabilities, capabilities)
}

func TestHandlerServerCapabilitiesTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerServerCapabilitiesTestSuite))
}
