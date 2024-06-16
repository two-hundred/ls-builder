package lsp

import "github.com/two-hundred/ls-builder/common"

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_declaration

const MethodGotoDeclaration = Method("textDocument/declaration")

// GoToDeclarationHandlerFunc is the function signature for the textDocument/declaration
// request handler that can be registered for a language server.
//
// Returns: Location | []Location | []LocationLink | nil
type GotoDeclarationHandlerFunc func(ctx *common.LSPContext, params *DeclarationParams) (any, error)

// DeclarationParams contains the textDocument/declaration request parameters.
type DeclarationParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_definition

const MethodGotoDefinition = Method("textDocument/definition")

// GoToDefinitionHandlerFunc is the function signature for the textDocument/definition
// request handler that can be registered for a language server.
//
// Returns: Location | []Location | []LocationLink | nil
type GotoDefinitionHandlerFunc func(ctx *common.LSPContext, params *DefinitionParams) (any, error)

// DefinitionParams contains the textDocument/definition request parameters.
type DefinitionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}
