package lsp

import (
	"encoding/json"

	"github.com/two-hundred/ls-builder/common"
)

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

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_hover

const MethodHover = Method("textDocument/hover")

// HoverHandlerFunc is the function signature for the textDocument/hover
// request handler that can be registered for a language server.
type HoverHandlerFunc func(
	ctx *common.LSPContext,
	params *HoverParams,
) (*Hover, error)

// HoverParams contains the textDocument/hover request parameters.
type HoverParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

// Hover represents the result of a hover request.
type Hover struct {
	// The hover's content.
	//
	// MarkedString | []MarkedString | MarkupContent
	Contents any `json:"contents"`

	// An optional range is a range inside a text document
	// that is used to visualize a hover, e.g. by changing
	// the background color.
	Range *Range `json:"range,omitempty"`
}

type hoverIntermediary struct {
	// MarkedString | []MarkedString | MarkupContent
	Contents json.RawMessage `json:"contents"`
	Range    *Range          `json:"range,omitempty"`
}

// Fulfils the json.Unmarshaler interface.
func (h *Hover) UnmarshalJSON(data []byte) error {

	var value hoverIntermediary

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	h.Range = value.Range

	var markupContentVal MarkupContent
	err := json.Unmarshal(value.Contents, &markupContentVal)
	if err == nil && markupContentVal.Kind != "" {
		h.Contents = markupContentVal
		return nil
	}

	var markedStringVal MarkedString
	if err = json.Unmarshal(value.Contents, &markedStringVal); err == nil {
		h.Contents = markedStringVal
		return nil
	}

	var markedStringArrayVal []MarkedString
	if err = json.Unmarshal(value.Contents, &markedStringArrayVal); err == nil {
		h.Contents = markedStringArrayVal
	}

	return err
}

// MarkedString can be used to render human readable text. It is either a
// markdown string or a code-block that provides a language and a code snippet.
// The language identifier is semantically equal to the optional language
// identifier in fenced code blocks in GitHub issues.
//
// The pair of a language and a value is an equivalent to markdown:
// ```${language}
// ${value}
// ```
//
// Note that markdown strings will be sanitized - that means html will be
// escaped.
//
// @deprecated use MarkupContent instead.
type MarkedString struct {
	Value any // string | MarkedStringLanguage
}

// Fulfils the json.Marshaler interface.
func (s MarkedString) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Value)
}

// Fulfils the json.Unmarshaler interface.
func (s *MarkedString) UnmarshalJSON(data []byte) error {
	var strVal string
	if err := json.Unmarshal(data, &strVal); err == nil {
		s.Value = strVal
		return nil
	} else {
		var markedStringLanguageVal MarkedStringLanguage
		if err := json.Unmarshal(data, &markedStringLanguageVal); err == nil {
			s.Value = markedStringLanguageVal
			return nil
		}
	}

	return nil
}

// MarkedStringLanguage is a pair of a language and a value for a MarkedString.
type MarkedStringLanguage struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_codeLens

const MethodCodeLens = Method("textDocument/codeLens")

// CodeLensHandlerFunc is the function signature for the textDocument/codeLens
// request handler that can be registered for a language server.
type CodeLensHandlerFunc func(
	ctx *common.LSPContext,
	params *CodeLensParams,
) ([]CodeLens, error)

// CodeLensParams contains the textDocument/codeLens request parameters.
type CodeLensParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The document to request code lens for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// CodeLens represents a command that should be shown along with
// source text, like the number of references, a way to run tests, etc.
//
// A code lens is _unresolved_ when no command is associated to it. For
// performance reasons the creation of a code lens and resolving should be done
// in two stages.
type CodeLens struct {
	// The range in which this code lens is valid. Should only span a single line.
	Range Range `json:"range"`

	// The command this code lens represents.
	Command *Command `json:"command,omitempty"`

	// A data entry field that is preserved on a code lens item between
	// a code lens and a code lens resolve request.
	Data LSPAny `json:"data,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#codeLens_resolve

const MethodCodeLensResolve = Method("codeLens/resolve")

// CodeLensResolveHandlerFunc is the function signature for the codeLens/resolve
// request handler that can be registered for a language server.
type CodeLensResolveHandlerFunc func(
	ctx *common.LSPContext,
	params *CodeLens,
) (*CodeLens, error)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#codeLens_refresh

const MethodCodeLensRefresh = Method("workspace/codeLens/refresh")

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_foldingRange

const MethodFoldingRange = Method("textDocument/foldingRange")

// FoldingRangeHandlerFunc is the function signature for the textDocument/foldingRange
// request handler that can be registered for a language server.
type FoldingRangeHandlerFunc func(
	ctx *common.LSPContext,
	params *FoldingRangeParams,
) ([]FoldingRange, error)

// FoldingRangeParams contains the textDocument/foldingRange request parameters.
type FoldingRangeParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// FoldingRange represents a folding range.
// To be valid, start and end line must be bigger
// than zero and smaller than the number of lines in the document.
// Clients are free to ignore invalid ranges.
type FoldingRange struct {
	// The zero-based start line of the range to fold. The folded area starts
	// after the line's last character. To be valid, the end must be zero or
	// larger and smaller than the number of lines in the document.
	StartLine UInteger `json:"startLine"`

	// The zero-based character offset from where the folded range starts. If
	// not defined, defaults to the length of the start line.
	StartCharacter *UInteger `json:"startCharacter,omitempty"`

	// The zero-based end line of the range to fold. The folded area ends with
	// the line's last character. To be valid, the end must be zero or larger
	// and smaller than the number of lines in the document.
	EndLine UInteger `json:"endLine"`

	// The zero-based character offset before the folded range ends. If not
	// defined, defaults to the length of the end line.
	EndCharacter *UInteger `json:"endCharacter,omitempty"`

	// Describes the kind of the folding range such as `comment` or `region`.
	// The kind is used to categorize folding ranges and used by commands like
	// 'Fold all comments'. See [FoldingRangeKind](#FoldingRangeKind) for an
	// enumeration of standardized kinds.
	Kind *FoldingRangeKind `json:"kind,omitempty"`

	// The text that the client should show when the specified range is
	// collapsed. If not defined or not supported by the client, a default
	// will be chosen by the client.
	//
	// @since 3.17.0 - proposed
	CollapsedText *string `json:"collapsedText,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_selectionRange

const MethodSelectionRange = Method("textDocument/selectionRange")

// SelectionRangeHandlerFunc is the function signature for the textDocument/selectionRange
// request handler that can be registered for a language server.
type SelectionRangeHandlerFunc func(
	ctx *common.LSPContext,
	params *SelectionRangeParams,
) ([]SelectionRange, error)

// SelectionRangeParams contains the textDocument/selectionRange request parameters.
type SelectionRangeParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The positions inside the text document.
	Positions []Position `json:"positions"`
}

// SelectionRange represents a selection range.
type SelectionRange struct {
	// The [range](#Range) of this selection range.
	Range Range `json:"range"`

	// The parent selection range containing this range.
	// Therefore `parent.range` must contain `this.range`.
	Parent *SelectionRange `json:"parent,omitempty"`
}
