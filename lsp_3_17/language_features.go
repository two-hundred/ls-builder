package lsp

import (
	"encoding/json"
	"errors"

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

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_documentSymbol

const MethodDocumentSymbol = Method("textDocument/documentSymbol")

// DocumentSymbolHandlerFunc is the function signature for the textDocument/documentSymbol
// request handler that can be registered for a language server.
//
// Returns: []DocumentSymbol | []SymbolInformation | nil
type DocumentSymbolHandlerFunc func(
	ctx *common.LSPContext,
	params *DocumentSymbolParams,
) (any, error)

// DocumentSymbolParams contains the textDocument/documentSymbol request parameters.
type DocumentSymbolParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// DocumentSymbol represents programming constructs like variables, classes, interfaces etc.
// that appear in a document. Document symbols can be hierarchical and they
// have two ranges: one that encloses its definition and one that points to its
// most interesting range, e.g. the range of an identifier.
type DocumentSymbol struct {
	// The name of this symbol. Will be displayed in the user interface and
	// therefore must not be an empty string or a string only consisting of
	// white spaces.
	Name string `json:"name"`

	// More detail for this symbol, e.g the signature of a function.
	Detail *string `json:"detail,omitempty"`

	// The kind of this symbol.
	Kind SymbolKind `json:"kind"`

	// Tags for this symbol.
	//
	// @since 3.16.0
	Tags []SymbolTag `json:"tags,omitempty"`

	// Indicates if this symbol is deprecated.
	//
	// @deprecated Use tags instead.
	Deprecated *bool `json:"deprecated,omitempty"`

	// The range enclosing this symbol not including leading/trailing whitespace
	// but everything else like comments. This information is typically used to
	// determine if the clients cursor is inside the symbol to reveal in the
	// symbol in the UI.
	Range Range `json:"range"`

	// The range that should be selected and revealed when this symbol is being
	// picked, e.g. the name of a function. Must be contained by the `range`.
	SelectionRange Range `json:"selectionRange"`

	// Children of this symbol, e.g. properties of a class.
	Children []DocumentSymbol `json:"children,omitempty"`
}

// SymbolInformation represents information about
// programming constructs like variables, classes, interfaces etc.
//
// @deprecated use DocumentSymbol or WorkspaceSymbol instead.
type SymbolInformation struct {
	// The name of this symbol.
	Name string `json:"name"`

	// The kind of this symbol.
	Kind SymbolKind `json:"kind"`

	// Tags for this symbol.
	//
	// @since 3.16.0
	Tags []SymbolTag `json:"tags,omitempty"`

	// Indicates if this symbol is deprecated.
	//
	// @deprecated Use tags instead.
	Deprecated *bool `json:"deprecated,omitempty"`

	// The location of this symbol. The location's range is used by a tool
	// to reveal the location in the editor. If the symbol is selected in the
	// tool the range's start information is used to position the cursor. So
	// the range usually spans more then the actual symbol's name and does
	// normally include things like visibility modifiers.
	//
	// The range doesn't have to denote a node range in the sense of an abstract
	// syntax tree. It can therefore not be used to re-construct a hierarchy of
	// the symbols.
	Location Location `json:"location"`

	// The name of the symbol containing this symbol. This information is for
	// user interface purposes (e.g. to render a qualifier in the user interface
	// if necessary). It can't be used to re-infer a hierarchy for the document
	// symbols.
	ContainerName *string `json:"containerName,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_semanticTokens

const MethodSemanticTokensFull = Method("textDocument/semanticTokens/full")

// SemanticTokensFullHandlerFunc is the function signature for the textDocument/semanticTokens/full
// request handler that can be registered for a language server.
type SemanticTokensFullHandlerFunc func(
	ctx *common.LSPContext,
	params *SemanticTokensParams,
) (*SemanticTokens, error)

// SemanticTokensParams contains the textDocument/semanticTokens/full request parameters.
type SemanticTokensParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// SemanticTokens represents a full set of semantic tokens.
type SemanticTokens struct {
	// An optional result id. If provided and clients support delta updating
	// the client will include the result id in the next semantic token request.
	// A server can then instead of computing all semantic tokens again simply
	// send a delta.
	ResultID *string `json:"resultId,omitempty"`

	// The actual tokens.
	Data []UInteger `json:"data"`
}

const MethodSemanticTokensFullDelta = Method("textDocument/semanticTokens/full/delta")

// SemanticTokensFullDeltaHandlerFunc is the function signature for the textDocument/semanticTokens/full/delta
// request handler that can be registered for a language server.
type SemanticTokensFullDeltaHandlerFunc func(
	ctx *common.LSPContext,
	params *SemanticTokensDeltaParams,
) (*SemanticTokensDelta, error)

// SemanticTokensDeltaParams contains the textDocument/semanticTokens/full/delta request parameters.
type SemanticTokensDeltaParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The result id of a previous response. The result Id can either point to
	// a full response or a delta response depending on what was received last.
	PreviousResultID string `json:"previousResultId"`
}

// SemanticTokensDelta represents a delta set of semantic tokens.
type SemanticTokensDelta struct {
	ResultID *string `json:"resultId,omitempty"`

	// The semantic token edits to transform a previous result into a new result.
	Edits []SemanticTokensEdit `json:"edits"`
}

// SemanticTokensEdit represents a semantic token edit.
type SemanticTokensEdit struct {
	// The start offset of the edit.
	Start UInteger `json:"start"`

	// The number of elements to remove.
	DeleteCount UInteger `json:"deleteCount"`

	// The elements to insert.
	Data []UInteger `json:"data,omitempty"`
}

const MethodSemanticTokensRange = Method("textDocument/semanticTokens/range")

// SemanticTokensRangeHandlerFunc is the function signature for the textDocument/semanticTokens/range
// request handler that can be registered for a language server.
type SemanticTokensRangeHandlerFunc func(
	ctx *common.LSPContext,
	params *SemanticTokensRangeParams,
) (*SemanticTokens, error)

// SemanticTokensRangeParams contains the textDocument/semanticTokens/range request parameters.
type SemanticTokensRangeParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The range the semantic tokens are requested for.
	Range Range `json:"range"`
}

const MethodSemanticTokensRefresh = Method("workspace/semanticTokens/refresh")

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#semanticTokenTypes

// SemanticTokenType is a predefined token types.
type SemanticTokenType = string

const (
	SemanticTokenTypeNamespace = SemanticTokenType("namespace")
	// Represents a generic type. Acts as a fallback for types which
	// can't be mapped to a specific type like class or enum.
	SemanticTokenTypeType          = SemanticTokenType("type")
	SemanticTokenTypeClass         = SemanticTokenType("class")
	SemanticTokenTypeEnum          = SemanticTokenType("enum")
	SemanticTokenTypeInterface     = SemanticTokenType("interface")
	SemanticTokenTypeStruct        = SemanticTokenType("struct")
	SemanticTokenTypeTypeParameter = SemanticTokenType("typeParameter")
	SemanticTokenTypeParameter     = SemanticTokenType("parameter")
	SemanticTokenTypeVariable      = SemanticTokenType("variable")
	SemanticTokenTypeProperty      = SemanticTokenType("property")
	SemanticTokenTypeEnumMember    = SemanticTokenType("enumMember")
	SemanticTokenTypeEvent         = SemanticTokenType("event")
	SemanticTokenTypeFunction      = SemanticTokenType("function")
	SemanticTokenTypeMethod        = SemanticTokenType("method")
	SemanticTokenTypeMacro         = SemanticTokenType("macro")
	SemanticTokenTypeKeyword       = SemanticTokenType("keyword")
	SemanticTokenTypeModifier      = SemanticTokenType("modifier")
	SemanticTokenTypeComment       = SemanticTokenType("comment")
	SemanticTokenTypeString        = SemanticTokenType("string")
	SemanticTokenTypeNumber        = SemanticTokenType("number")
	SemanticTokenTypeRegexp        = SemanticTokenType("regexp")
	SemanticTokenTypeOperator      = SemanticTokenType("operator")
)

// SemanticTokenModifier is a predefined token modifiers.
type SemanticTokenModifier string

const (
	SemanticTokenModifierDeclaration    = SemanticTokenModifier("declaration")
	SemanticTokenModifierDefinition     = SemanticTokenModifier("definition")
	SemanticTokenModifierReadonly       = SemanticTokenModifier("readonly")
	SemanticTokenModifierStatic         = SemanticTokenModifier("static")
	SemanticTokenModifierDeprecated     = SemanticTokenModifier("deprecated")
	SemanticTokenModifierAbstract       = SemanticTokenModifier("abstract")
	SemanticTokenModifierAsync          = SemanticTokenModifier("async")
	SemanticTokenModifierModification   = SemanticTokenModifier("modification")
	SemanticTokenModifierDocumentation  = SemanticTokenModifier("documentation")
	SemanticTokenModifierDefaultLibrary = SemanticTokenModifier("defaultLibrary")
)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_inlayHint

const MethodInlayHint = Method("textDocument/inlayHint")

// InlayHintHandlerFunc is the function signature for the textDocument/inlayHint
// request handler that can be registered for a language server.
type InlayHintHandlerFunc func(
	ctx *common.LSPContext,
	params *InlayHintParams,
) ([]*InlayHint, error)

// InlayHintParams contains the textDocument/inlayHint request parameters.
// A parameter literal used in inlay hints requests.
//
// @since 3.17.0
type InlayHintParams struct {
	WorkDoneProgressParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The visible document range for which inlay hints should be computed.
	Range Range `json:"range"`
}

// InlayHint holds inlay hint information.
//
// @since 3.17.0
type InlayHint struct {
	// The position of this hint.
	//
	// If multiple hints have the same position, they will be shown in the order
	// they appear in the response.
	Position Position `json:"position"`

	// The label of this hint. A human readable string or an array of
	// InlayHintLabelPart label parts.
	//
	// *Note* that neither the string nor the label part can be empty.
	//
	// string | []*InlayHintLabelPart
	Label any `json:"label"`

	// The kind of this hint. Can be omitted in which case the client
	// should fall back to a reasonable default.
	Kind *InlayHintKind `json:"kind,omitempty"`

	// Optional text edits that are performed when accepting this inlay hint.
	//
	// *Note* that edits are expected to change the document so that the inlay
	// hint (or its nearest variant) is now part of the document and the inlay
	// hint itself is now obsolete.
	//
	// Depending on the client capability `inlayHint.resolveSupport` clients
	// might resolve this property late using the resolve request.
	TextEdits []TextEdit `json:"textEdits,omitempty"`

	// The tooltip text when you hover over this item.
	//
	// Depending on the client capability `inlayHint.resolveSupport` clients
	// might resolve this property late using the resolve request.
	//
	// string | MarkupContent | nil
	Tooltip any `json:"tooltip,omitempty"`

	// Render padding before the hint.
	//
	// Note: Padding should use the editor's background color, not the
	// background color of the hint itself. That means padding can be used
	// to visually align/separate an inlay hint.
	PaddingLeft *bool `json:"paddingLeft,omitempty"`

	// Render padding after the hint.
	//
	// Note: Padding should use the editor's background color, not the
	// background color of the hint itself. That means padding can be used
	// to visually align/separate an inlay hint.
	PaddingRight *bool `json:"paddingRight,omitempty"`

	// A data entry field that is preserved on an inlay hint between
	// a `textDocument/inlayHint` and a `inlayHint/resolve` request.
	Data LSPAny `json:"data,omitempty"`
}

type inlayHintIntermediary struct {
	Position Position `json:"position"`
	// string | []*InlayHintLabelPart
	Label     json.RawMessage `json:"label"`
	Kind      *InlayHintKind  `json:"kind,omitempty"`
	TextEdits []TextEdit      `json:"textEdits,omitempty"`
	// string | MarkupContent | nil
	Tooltip      json.RawMessage `json:"tooltip,omitempty"`
	PaddingLeft  *bool           `json:"paddingLeft,omitempty"`
	PaddingRight *bool           `json:"paddingRight,omitempty"`
	Data         LSPAny          `json:"data,omitempty"`
}

// Fulfils the json.Unmarshaler interface.
func (h *InlayHint) UnmarshalJSON(data []byte) error {

	var value inlayHintIntermediary

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	h.Position = value.Position
	h.Kind = value.Kind
	h.TextEdits = value.TextEdits
	h.PaddingLeft = value.PaddingLeft
	h.PaddingRight = value.PaddingRight
	h.Data = value.Data

	err := h.unmarshalInlayHintLabel(&value)
	if err != nil {
		return err
	}

	err = h.unmarshalInlayHintTooltip(&value)
	return err
}

func (h *InlayHint) unmarshalInlayHintLabel(value *inlayHintIntermediary) error {
	var strVal string
	if err := json.Unmarshal(value.Label, &strVal); err == nil {
		h.Label = strVal
		return nil
	}

	inlayHintLabelPartArrayVal := []*InlayHintLabelPart{}
	err := json.Unmarshal(value.Label, &inlayHintLabelPartArrayVal)
	if err == nil {
		h.Label = inlayHintLabelPartArrayVal
	}

	return err
}

func (h *InlayHint) unmarshalInlayHintTooltip(value *inlayHintIntermediary) error {
	var strVal string
	if err := json.Unmarshal(value.Tooltip, &strVal); err == nil {
		h.Tooltip = strVal
		return nil
	}

	var markupContentVal MarkupContent
	err := json.Unmarshal(value.Tooltip, &markupContentVal)
	if err == nil && markupContentVal.Kind != "" {
		h.Tooltip = markupContentVal
	}

	// Ignore the error as this field is optional.
	return nil
}

// InlayHintKind is a kind of inlay hint.
//
// @since 3.17.0
type InlayHintKind = UInteger

var (
	// InlayHintKindType is an inlay hint
	// for a type annotation.
	InlayHintKindType InlayHintKind = 1

	// InlayHintKindParameter is an inlay hint
	// for a parameter annotation.
	InlayHintKindParameter InlayHintKind = 2
)

// InlayHintLabelPart represents a part of a label in an inlay hint.
// An inlay hint label part allows for interactive and composite labels
// of inlay hints.
//
// @since 3.17.0
type InlayHintLabelPart struct {
	// The value of this label part.
	Value string `json:"value"`

	// The tooltip text when you hover over this label part. Depending on
	// the client capability `inlayHint.resolveSupport` clients might resolve
	// this property late using the resolve request.
	//
	// string | MarkupContent | nil
	Tooltip any `json:"tooltip,omitempty"`

	// An optional source code location that represents this
	// label part.
	//
	// The editor will use this location for the hover and for code navigation
	// features: This part will become a clickable link that resolves to the
	// definition of the symbol at the given location (not necessarily the
	// location itself), it shows the hover that shows at the given location,
	// and it shows a context menu with further code navigation commands.
	//
	// Depending on the client capability `inlayHint.resolveSupport` clients
	// might resolve this property late using the resolve request.
	Location *Location `json:"location,omitempty"`

	// An optional command for this label part.
	//
	// Depending on the client capability `inlayHint.resolveSupport` clients
	// might resolve this property late using the resolve request.
	Command *Command `json:"command,omitempty"`
}

type inlayHintLabelPartIntermediary struct {
	Value string `json:"value"`
	// string | MarkupContent | nil
	Tooltip  json.RawMessage `json:"tooltip,omitempty"`
	Location *Location       `json:"location,omitempty"`
	Command  *Command        `json:"command,omitempty"`
}

// Fulfils the json.Unmarshaler interface.
func (lp *InlayHintLabelPart) UnmarshalJSON(data []byte) error {

	var value inlayHintLabelPartIntermediary

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	lp.Value = value.Value
	lp.Location = value.Location
	lp.Command = value.Command

	var strTooltip string
	if err := json.Unmarshal(value.Tooltip, &strTooltip); err == nil {
		lp.Tooltip = strTooltip
		return nil
	}

	var markupContentTooltip MarkupContent
	err := json.Unmarshal(value.Tooltip, &markupContentTooltip)
	if err == nil && markupContentTooltip.Kind != "" {
		lp.Tooltip = markupContentTooltip
	}

	// Ignore the error as this field is optional.
	return nil
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#inlayHint_resolve

const MethodInlayHintResolve = Method("inlayHint/resolve")

// InlayHintResolveHandlerFunc is the function signature for the inlayHint/resolve
// request handler that can be registered for a language server.
type InlayHintResolveHandlerFunc func(
	ctx *common.LSPContext,
	params *InlayHint,
) (*InlayHint, error)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspace_inlayHint_refresh

const MethodInlayHintRefresh = Method("workspace/inlayHint/refresh")

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_inlineValue

const MethodInlineValue = Method("textDocument/inlineValue")

// InlineValueHandlerFunc is the function signature for the textDocument/inlineValue
// request handler that can be registered for a language server.
//
// Inline value information can be provided by different means:
// - directly as a text value (class InlineValueText).
// - as a name to use for a variable lookup (class InlineValueVariableLookup)
// - as an evaluatable expression (class InlineValueEvaluatableExpression)
// The InlineValue types combines all inline value types into one type.
//
// @since 3.17.0
type InlineValueHandlerFunc func(
	ctx *common.LSPContext,
	params *InlineValueParams,
) ([]*InlineValue, error)

// Inline value information can be provided by different means:
// - directly as a text value (class InlineValueText).
// - as a name to use for a variable lookup (class InlineValueVariableLookup)
// - as an evaluatable expression (class InlineValueEvaluatableExpression)
// The InlineValue types combines all inline value types into one type.
//
// @since 3.17.0
type InlineValue struct {
	InlineValueText           *InlineValueText                  `json:"inlineValueText,omitempty"`
	InlineValueVariableLookup *InlineValueVariableLookup        `json:"inlineValueVariableLookup,omitempty"`
	InlineValueEvaluatable    *InlineValueEvaluatableExpression `json:"inlineValueEvaluatable,omitempty"`
}

// Fulfils the json.Marshaler interface.
func (iv *InlineValue) MarshalJSON() ([]byte, error) {
	if iv.InlineValueText != nil {
		return json.Marshal(iv.InlineValueText)
	} else if iv.InlineValueVariableLookup != nil {
		return json.Marshal(iv.InlineValueVariableLookup)
	} else if iv.InlineValueEvaluatable != nil {
		return json.Marshal(iv.InlineValueEvaluatable)
	}

	return nil, errors.New("one InlineValue type must be set")
}

// Fulfils the json.Unmarshaler interface.
func (iv *InlineValue) UnmarshalJSON(data []byte) error {
	var ivText InlineValueText
	if err := json.Unmarshal(data, &ivText); err == nil && ivText.Text != "" {
		iv.InlineValueText = &ivText
		return nil
	}

	var ivVariableLookup InlineValueVariableLookup
	if err := json.Unmarshal(data, &ivVariableLookup); err == nil && ivVariableLookup.VariableName != nil {
		iv.InlineValueVariableLookup = &ivVariableLookup
		return nil
	}

	var ivEvaluatable InlineValueEvaluatableExpression
	if err := json.Unmarshal(data, &ivEvaluatable); err == nil && ivEvaluatable.Expression != nil {
		iv.InlineValueEvaluatable = &ivEvaluatable
		return nil
	}

	return errors.New("one InlineValue type must be set")
}

// InlineValueParams contains the textDocument/inlineValue request parameters.
//
// @since 3.17.0
type InlineValueParams struct {
	WorkDoneProgressParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The document range for which inline values should be computed.
	Range Range `json:"range"`

	// Additional information about the context in which inline values were
	// requested.
	Context InlineValueContext `json:"context"`
}

// InlineValueContext provides additional information about the context in
// which inline values were requested.
//
// @since 3.17.0
type InlineValueContext struct {
	// The stack frame (as a DAP Id) where the execution has stopped.
	FrameID Integer `json:"frameId,omitempty"`

	// The document range where execution has stopped.
	// Typically the end position of the range denotes the line where the
	// inline values are shown.
	StoppedLocation Range `json:"stoppedLocation"`
}

// InlineValueText provides inline value as text.
//
// @since 3.17.0
type InlineValueText struct {
	// The document range for which the inline value applies.
	Range Range `json:"range"`

	// The text of the inline value.
	Text string `json:"text"`
}

// InlineValueVariableLookup provides inline value as a variable lookup.
//
// If only a range is specified, the variable name will be extracted from
// the underlying document.
//
// An optional variable name can be used to override the extracted name.
//
// @since 3.17.0
type InlineValueVariableLookup struct {
	// The document range for which the inline value applies.
	// The range is used to extract the variable name from the underlying
	// document.
	Range Range `json:"range"`

	// If specified, the name of the variable to lookup.
	VariableName *string `json:"variableName,omitempty"`

	// How to perform the lookup.
	CaseSensitiveLookup bool `json:"caseSensitiveLookup"`
}

// InlineValueEvaluatableExpression provides inline value as an evaluatable expression.
//
// If only a range is specified, the expression will be extracted from the
// underlying document.
//
// An optional expression can be used to override the extracted expression.
//
// @since 3.17.0
type InlineValueEvaluatableExpression struct {
	// The document range for which the inline value applies.
	// The range is used to extract the evaluatable expression from the
	// underlying document.
	Range Range `json:"range"`

	// If specified, the expression overrides the extracted expression.
	Expression *string `json:"expression,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspace_inlineValue_refresh

const MethodInlineValueRefresh = Method("workspace/inlineValue/refresh")

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_moniker

const MethodMoniker = Method("textDocument/moniker")

// MonikerHandlerFunc is the function signature for the textDocument/moniker
// request handler that can be registered for a language server.
type MonikerHandlerFunc func(
	ctx *common.LSPContext,
	params *MonikerParams,
) ([]Moniker, error)

// MonikerParams contains the textDocument/moniker request parameters.
type MonikerParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// Moniker definition to match the LSIF 0.5 moniker definition.
type Moniker struct {
	// The scheme of the moniker, e.g. tsc or .NET
	Scheme string `json:"scheme"`

	// The identifier of the moniker. The value is opaque in LSIF however
	// schema owners are allowed to define the structure if they want.
	Identifier string `json:"identifier"`

	// The scope in which the moniker is unique.
	Unique UniquenessLevel `json:"unique"`

	// The moniker kind, if known.
	Kind *MonikerKind `json:"kind,omitempty"`
}

type MonikerKind = string

var (
	// The moniker represents a symbol that is imported into a project
	MonikerKindImport MonikerKind = "import"

	// The moniker represents a symbol that is exported from a project
	MonikerKindExport MonikerKind = "export"

	// The moniker represents a symbol that is local to a project (e.g. a local
	// variable of a function, a class not visible outside the project, ...)
	MonikerKindLocal MonikerKind = "local"
)

// UniquenessLevel defines in which scope the monikor is unique.
type UniquenessLevel = string

const (
	// The moniker is only unique inside a document.
	UniquenessLevelDocument UniquenessLevel = "document"

	// The moniker is unique inside a project for which a dump was created.
	UniquenessLevelProject UniquenessLevel = "project"

	// The moniker is unique inside the group to which a project belongs.
	UniquenessLevelGroup UniquenessLevel = "group"

	// The moniker is unique inside the moniker scheme.
	UniquenessLevelScheme UniquenessLevel = "scheme"

	// The moniker is globally unique.
	UniquenessLevelGlobal UniquenessLevel = "global"
)
