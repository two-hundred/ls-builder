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

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#callHierarchy_incomingCalls

const MethodCallHierarchyIncomingCalls = Method("callHierarchy/incomingCalls")

// CallHierarchyIncomingCallsHandlerFunc is the function signature for the callHierarchy/incomingCalls
// request handler that can be registered for a language server.
type CallHierarchyIncomingCallsHandlerFunc func(
	ctx *common.LSPContext,
	params *CallHierarchyIncomingCallsParams,
) ([]CallHierarchyIncomingCall, error)

// CallHierarchyIncomingCallsParams contains the callHierarchy/incomingCalls request parameters.
type CallHierarchyIncomingCallsParams struct {
	WorkDoneProgressParams
	PartialResultParams

	Item CallHierarchyItem `json:"item"`
}

// CallHierarchyIncomingCall represents an incoming call within the call hierarchy.
type CallHierarchyIncomingCall struct {
	// The item that makes the call.
	From CallHierarchyItem `json:"from"`

	// The range at which at which the calls appears. This is relative to the caller
	// denoted by [`this.from`](#CallHierarchyIncomingCall.from).
	FromRanges []Range `json:"fromRanges"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#callHierarchy_outgoingCalls

const MethodCallHierarchyOutgoingCalls = Method("callHierarchy/outgoingCalls")

// CallHierarchyOutgoingCallsHandlerFunc is the function signature for the callHierarchy/outgoingCalls
// request handler that can be registered for a language server.
type CallHierarchyOutgoingCallsHandlerFunc func(
	ctx *common.LSPContext,
	params *CallHierarchyOutgoingCallsParams,
) ([]CallHierarchyOutgoingCall, error)

// CallHierarchyOutgoingCallsParams contains the callHierarchy/outgoingCalls request parameters.
type CallHierarchyOutgoingCallsParams struct {
	WorkDoneProgressParams
	PartialResultParams

	Item CallHierarchyItem `json:"item"`
}

// CallHierarchyOutgoingCall represents an outgoing call within the call hierarchy.
type CallHierarchyOutgoingCall struct {
	// The item that is called.
	To CallHierarchyItem `json:"to"`

	// The range at which this item is called. This is the range relative to
	// the caller, e.g the item passed to `callHierarchy/outgoingCalls` request.
	FromRanges []Range `json:"fromRanges"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_prepareTypeHierarchy

const MethodPrepareTypeHierarchy = Method("textDocument/prepareTypeHierarchy")

// PrepareTypeHierarchyHandlerFunc is the function signature for the textDocument/prepareTypeHierarchy
// request handler that can be registered for a language server.
type PrepareTypeHierarchyHandlerFunc func(
	ctx *common.LSPContext,
	params *TypeHierarchyPrepareParams,
) ([]TypeHierarchyItem, error)

// TypeHierarchyPrepareParams contains the textDocument/prepareTypeHierarchy request parameters.
type TypeHierarchyPrepareParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

// TypeHierarchyItem represents an item within the type hierarchy.
type TypeHierarchyItem struct {
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
	// [`range`](#TypeHierarchyItem.range).
	SelectionRange Range `json:"selectionRange"`

	// A data entry field that is preserved between a type hierarchy prepare and
	// supertypes or subtypes requests. It could also be used to identify the
	// type hierarchy in the server, helping improve the performance on
	// resolving supertypes and subtypes.
	Data LSPAny `json:"data,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#typeHierarchy_supertypes

const MethodTypeHierarchySupertypes = Method("typeHierarchy/supertypes")

// TypeHierarchySupertypesHandlerFunc is the function signature for the typeHierarchy/supertypes
// request handler that can be registered for a language server.
type TypeHierarchySupertypesHandlerFunc func(
	ctx *common.LSPContext,
	params *TypeHierarchySupertypesParams,
) ([]TypeHierarchyItem, error)

// TypeHierarchySupertypesParams contains the typeHierarchy/supertypes request parameters.
type TypeHierarchySupertypesParams struct {
	WorkDoneProgressParams
	PartialResultParams

	Item TypeHierarchyItem `json:"item"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#typeHierarchy_subtypes

const MethodTypeHierarchySubtypes = Method("typeHierarchy/subtypes")

// TypeHierarchySubtypesHandlerFunc is the function signature for the typeHierarchy/subtypes
// request handler that can be registered for a language server.
type TypeHierarchySubtypesHandlerFunc func(
	ctx *common.LSPContext,
	params *TypeHierarchySubtypesParams,
) ([]TypeHierarchyItem, error)

// TypeHierarchySubtypesParams contains the typeHierarchy/subtypes request parameters.
type TypeHierarchySubtypesParams struct {
	WorkDoneProgressParams
	PartialResultParams

	Item TypeHierarchyItem `json:"item"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_documentHighlight

const MethodDocumentHighlight = Method("textDocument/documentHighlight")

// DocumentHighlightHandlerFunc is the function signature for the textDocument/documentHighlight
// request handler that can be registered for a language server.
type DocumentHighlightHandlerFunc func(
	ctx *common.LSPContext,
	params *DocumentHighlightParams,
) ([]DocumentHighlight, error)

// DocumentHighlightParams contains the textDocument/documentHighlight request parameters.
type DocumentHighlightParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// DocumentHighlight represents a document highlight.
type DocumentHighlight struct {
	// The range this highlight applies to.
	Range Range `json:"range"`

	// The highlight kind, default is DocumentHighlightKind.Text.
	Kind *DocumentHighlightKind `json:"kind,omitempty"`
}

type DocumentHighlightKind = Integer

var (
	// DocumentHighlightKindText is for a textual occurrence.
	DocumentHighlightKindText DocumentHighlightKind = 1

	// DocumentHighlightKindRead is for read-access of a symbol, like reading a variable.
	DocumentHighlightKindRead DocumentHighlightKind = 2

	// DocumentHighlightKindWrite is for write-access of a symbol, like writing to a variable.
	DocumentHighlightKindWrite DocumentHighlightKind = 3
)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_documentLink

const MethodDocumentLink = Method("textDocument/documentLink")

// DocumentLinkHandlerFunc is the function signature for the textDocument/documentLink
// request handler that can be registered for a language server.
type DocumentLinkHandlerFunc func(
	ctx *common.LSPContext,
	params *DocumentLinkParams,
) ([]DocumentLink, error)

// DocumentLinkParams contains the textDocument/documentLink request parameters.
type DocumentLinkParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The document to provide document links for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// DocumentLink represents a document link.
// This is range in a text document that links to an internal or
// external resource, like another text document or a web site.
type DocumentLink struct {
	// The range this link applies to.
	Range Range `json:"range"`

	// The uri this link points to. If missing a resolve request is sent later.
	Target *DocumentURI `json:"target,omitempty"`

	// The tooltip text when you hover over this link.
	//
	// If a tooltip is provided, is will be displayed in a string that includes
	// instructions on how to trigger the link, such as `{0} (ctrl + click)`.
	// The specific instructions vary depending on OS, user settings, and
	// localization.
	//
	// @since 3.15.0
	Tooltip *string `json:"tooltip,omitempty"`

	// A data entry field that is preserved on a document link between a
	// DocumentLinkRequest and a DocumentLinkResolveRequest.
	Data LSPAny `json:"data,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#documentLink_resolve

const MethodDocumentLinkResolve = Method("documentLink/resolve")

// DocumentLinkResolveHandlerFunc is the function signature for the documentLink/resolve
// request handler that can be registered for a language server.
type DocumentLinkResolveHandlerFunc func(
	ctx *common.LSPContext,
	params *DocumentLink,
) (*DocumentLink, error)
