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

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_typeDefinition

const MethodGotoTypeDefinition = Method("textDocument/typeDefinition")

// GoToTypeDefinitionHandlerFunc is the function signature for the textDocument/typeDefinition
// request handler that can be registered for a language server.
//
// Returns: Location | []Location | []LocationLink | nil
type GotoTypeDefinitionHandlerFunc func(ctx *common.LSPContext, params *TypeDefinitionParams) (any, error)

// TypeDefinitionParams contains the textDocument/typeDefinition request parameters.
type TypeDefinitionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_implementation

const MethodGotoImplementation = Method("textDocument/implementation")

// GoToImplementationHandlerFunc is the function signature for the textDocument/implementation
// request handler that can be registered for a language server.
//
// Returns: Location | []Location | []LocationLink | nil
type GotoImplementationHandlerFunc func(ctx *common.LSPContext, params *ImplementationParams) (any, error)

// ImplementationParams contains the textDocument/implementation request parameters.
type ImplementationParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_references

const MethodFindReferences = Method("textDocument/references")

// FindReferencesHandlerFunc is the function signature for the textDocument/references
// request handler that can be registered for a language server.
type FindReferencesHandlerFunc func(ctx *common.LSPContext, params *ReferencesParams) ([]Location, error)

// ReferencesParams contains the textDocument/references request parameters.
type ReferencesParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams

	Context ReferenceContext `json:"context"`
}

// ReferenceContext contains additional information for the textDocument/references request.
type ReferenceContext struct {
	// Include the declaration of the current symbol.
	IncludeDeclaration bool `json:"includeDeclaration"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_prepareCallHierarchy

const MethodPrepareCallHierarchy = Method("textDocument/prepareCallHierarchy")

// PrepareCallHierarchyHandlerFunc is the function signature for the textDocument/prepareCallHierarchy
// request handler that can be registered for a language server.
type PrepareCallHierarchyHandlerFunc func(ctx *common.LSPContext, params *CallHierarchyPrepareParams) ([]CallHierarchyItem, error)

// CallHierarchyPrepareParams contains the textDocument/prepareCallHierarchy request parameters.
type CallHierarchyPrepareParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

// CallHierarchyItem represents an item within the call hierarchy.
type CallHierarchyItem struct {
	// The name of this item.
	Name string `json:"name"`

	// The kind of this item.
	Kind SymbolKind `json:"kind"`

	// Tags for this item.
	Tags []SymbolTag `json:"tags,omitempty"`

	// More detail for this item, e.g. the signature of a function.
	Detail *string `json:"detail,omitempty"`

	// The resource identifier of this item.
	URI DocumentURI `json:"uri"`

	// The range enclosing this symbol not including leading/trailing whitespace
	// but everything else, e.g. comments and code.
	Range Range `json:"range"`

	// The range that should be selected and revealed when this symbol is being
	// picked, e.g. the name of a function. Must be contained by the
	// [`range`](#CallHierarchyItem.range).
	SelectionRange Range `json:"selectionRange"`

	// A data entry field that is preserved between a call hierarchy prepare and
	// incoming calls or outgoing calls requests.
	Data any `json:"data,omitempty"`
}
