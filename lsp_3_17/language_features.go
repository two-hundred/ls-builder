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

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_completion

const MethodCompletion = Method("textDocument/completion")

// CompletionHandlerFunc is the function signature for the textDocument/completion
// request handler that can be registered for a language server.
//
// Returns: *CompletionList | []*CompletionItem | nil
type CompletionHandlerFunc func(
	ctx *common.LSPContext,
	params *CompletionParams,
) (any, error)

// CompletionParams contains the textDocument/completion request parameters.
type CompletionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams

	// The completion context. This is only available if the client specifies
	// to send this using the client capability
	// `completion.contextSupport === true`
	Context *CompletionContext `json:"context,omitempty"`
}

// CompletionContext contains additional information abou the context in which a
// completion request is triggered.
type CompletionContext struct {
	// How the completion was triggered.
	TriggerKind CompletionTriggerKind `json:"triggerKind"`

	// The trigger character (a single character) that has trigger code
	// complete. Is undefined if
	// `triggerKind !== CompletionTriggerKind.TriggerCharacter`
	TriggerCharacter *string `json:"triggerCharacter,omitempty"`
}

// CompletionTriggerKind defines how a completion request was triggered.
type CompletionTriggerKind = Integer

const (
	// CompletionTriggerKindInvoked means the completion was triggered
	// by the user typing an identifier (24x7 code complete),
	// manual invocation (e.g. Ctrl+Space) or via API.
	CompletionTriggerKindInvoked CompletionTriggerKind = 1

	// CompletionTriggerKindTriggerCharacter means the completion was
	// triggered by a trigger character specified by the `triggerCharacters`
	// properties of the `CompletionRegistrationOptions`.
	CompletionTriggerKindTriggerCharacter CompletionTriggerKind = 2

	// CompletionTriggerKindTriggerForIncompleteCompletions means the
	// completion was re-triggered as the current completion list is
	// incomplete.
	CompletionTriggerKindTriggerForIncompleteCompletions CompletionTriggerKind = 3
)

// CompletionList represents a collection of [completion items](#CompletionItem)
// to be presented in the editor.
type CompletionList struct {
	// This list is not complete. Further typing should result in recomputing
	// this list.
	//
	// Recomputed lists have all their items replaced (not appended) in the
	// incomplete completion sessions.
	IsIncomplete bool `json:"isIncomplete"`

	// In many cases the items of an actual completion result share the same
	// value for properties like `commitCharacters` or the range of a text
	// edit. A completion list can therefore define item defaults which will
	// be used if a completion item itself doesn't specify the value.
	//
	// If a completion list specifies a default value and a completion item
	// also specifies a corresponding value the one from the item is used.
	//
	// Servers are only allowed to return default values if the client
	// signals support for this via the `completionList.itemDefaults`
	// capability.
	//
	// @since 3.17.0
	ItemDefaults *CompletionItemDefaults `json:"itemDefaults,omitempty"`

	// The completion items.
	Items []*CompletionItem `json:"items"`
}

// CompletionItemDefaults provides default values for each kind of completion
// item.
type CompletionItemDefaults struct {
	// A default commit character set.
	//
	// @since 3.17.0
	CommitCharacters []string `json:"commitCharacters,omitempty"`

	// A default edit range.
	//
	// @since 3.17.0
	//
	// Range | InsertReplaceRange | nil
	EditRange any `json:"editRange,omitempty"`

	// A default insert text format.
	//
	// @since 3.17.0
	InsertTextFormat *InsertTextFormat `json:"insertTextFormat,omitempty"`

	// A default insert text mode.
	//
	// @since 3.17.0
	InsertTextMode *InsertTextMode `json:"insertTextMode,omitempty"`

	// A default data value.
	//
	// @since 3.17.0
	Data LSPAny `json:"data,omitempty"`
}

type completionItemDefaultsIntermediary struct {
	CommitCharacters []string `json:"commitCharacters,omitempty"`
	// Range | InsertReplaceRange | nil
	EditRange        json.RawMessage   `json:"editRange,omitempty"`
	InsertTextFormat *InsertTextFormat `json:"insertTextFormat,omitempty"`
	InsertTextMode   *InsertTextMode   `json:"insertTextMode,omitempty"`
	Data             LSPAny            `json:"data,omitempty"`
}

// Fulfils the json.Unmarshaler interface.
func (d *CompletionItemDefaults) UnmarshalJSON(data []byte) error {

	var value completionItemDefaultsIntermediary

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	d.CommitCharacters = value.CommitCharacters
	d.InsertTextFormat = value.InsertTextFormat
	d.InsertTextMode = value.InsertTextMode
	d.Data = value.Data

	d.unmarshalEditRange(&value)
	// On failure to unmarhsal edit range, we can fallback to a nil valuable
	// as the field is optional.
	return nil

}

func (d *CompletionItemDefaults) unmarshalEditRange(value *completionItemDefaultsIntermediary) error {
	var insertReplaceRangeVal InsertReplaceRange
	err := json.Unmarshal(value.EditRange, &insertReplaceRangeVal)
	if err == nil && (insertReplaceRangeVal.Insert != nil || insertReplaceRangeVal.Replace != nil) {
		d.EditRange = insertReplaceRangeVal
		return nil
	}

	var rangeVal Range
	if err := json.Unmarshal(value.EditRange, &rangeVal); err == nil {
		d.EditRange = rangeVal
		return nil
	}

	return err
}

// InsertReplaceRange contains both insert and replace ranges
// for an edit.
type InsertReplaceRange struct {
	Insert  *Range `json:"insert"`
	Replace *Range `json:"replace"`
}

// InsertTextFormat defins whether the insert text in a completion item should be
// interpreted as plain text or a snippet.
type InsertTextFormat = Integer

var (
	// InsertTextFormatPlainText means the primary text
	// to be inserted is treated as a plain string.
	InsertTextFormatPlainText InsertTextFormat = 1

	// InsertTextFormatSnippet means the primary text
	// to be instered is to be treated as a snippet.
	//
	// A snippet can define tab stops and placeholders with `$1`, `$2`
	// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
	// the end of the snippet. Placeholders with equal identifiers are linked,
	// that is typing in one will update others too.
	InsertTextFormatSnippet InsertTextFormat = 2
)

// InsertTextMode determines how whitespace and indentations are handled
// during completion item insertion.
//
// @since 3.16.0
type InsertTextMode = Integer

var (
	// InsertTextModeAsIs means the insertion or replace strings are taken as
	// they are.
	// If the
	// value is multi line the lines below the cursor will be
	// inserted using the indentation defined in the string value.
	// The client will not apply any kind of adjustments to the
	// string.
	InsertTextModeAsIs InsertTextMode = 1

	// InsertTextModeAdjustIndentation means the editor adjusts leading
	// whitespace of new lines so that they match the indentation up to the
	// cursor of the line for which the item is accepted.
	//
	// Consider a line like this: <2tabs><cursor><3tabs>foo. Accepting a
	// multi line completion item is indented using 2 tabs and all
	// following lines inserted will be indented using 2 tabs as well.
	InsertTextModeAdjustIndentation InsertTextMode = 2
)

// CompletionItem represents a completion item that is presented in a completion
// list in a client editor.
type CompletionItem struct {
	// The label of this completion item.
	//
	// The label property is also by default the text that
	// is inserted when selecting this completion.
	//
	// If label details are provided the label itself should
	// be an unqualified name of the completion item.
	Label string `json:"label"`

	// Additional details for the label.
	//
	// @since 3.17.0
	LabelDetails *CompletionItemLabelDetails `json:"labelDetails,omitempty"`

	// The kind of this completion item. Based of the kind
	// an icon is chosen by the editor. The standardized set
	// of available values is defined in `CompletionItemKind`.
	Kind *CompletionItemKind `json:"kind,omitempty"`

	// Tags for this completion item.
	//
	// @since 3.15.0
	Tags []CompletionItemTag `json:"tags,omitempty"`

	// A human-readable string with additional information
	// about this item, like type or symbol information.
	Detail *string `json:"detail,omitempty"`

	// A human-readable string that represents a doc-comment.
	//
	// string | MarkupContent | nil
	Documentation any `json:"documentation,omitempty"`

	// Indicates if this item is deprecated.
	//
	// @deprecated Use `tags` instead if supported.
	Deprecated *bool `json:"deprecated,omitempty"`

	// Select this item when showing.
	//
	// *Note* that only one completion item can be selected and that the
	// tool / client decides which item that is. The rule is that the *first*
	// item of those that match best is selected.
	Preselect *bool `json:"preselect,omitempty"`

	// A string that should be used when comparing this item
	// with other items. When omitted the label is used
	// as the sort text for this item.
	SortText *string `json:"sortText,omitempty"`

	// A string that should be used when filtering a set of
	// completion items. When omitted the label is used as the
	// filter text for this item.
	FilterText *string `json:"filterText,omitempty"`

	// A string that should be inserted into a document when selecting
	// this completion. When omitted the label is used as the insert text
	// for this item.
	//
	// The `insertText` is subject to interpretation by the client side.
	// Some tools might not take the string literally. For example
	// VS Code when code complete is requested in this example
	// `con<cursor position>` and a completion item with an `insertText` of
	// `console` is provided it will only insert `sole`. Therefore it is
	// recommended to use `textEdit` instead since it avoids additional client
	// side interpretation.
	InsertText *string `json:"insertText,omitempty"`

	// The format of the insert text. The format applies to both the
	// `insertText` property and the `newText` property of a provided
	// `textEdit`. If omitted defaults to `InsertTextFormat.PlainText`.
	//
	// Please note that the insertTextFormat doesn't apply to
	// `additionalTextEdits`.
	InsertTextFormat *InsertTextFormat `json:"insertTextFormat,omitempty"`

	// How whitespace and indentation is handled during completion
	// item insertion. If not provided the client's default value depends on
	// the `textDocument.completion.insertTextMode` client capability.
	//
	// @since 3.16.0
	// @since 3.17.0 - support for `textDocument.completion.insertTextMode`
	InsertTextMode *InsertTextMode `json:"insertTextMode,omitempty"`

	// An edit which is applied to a document when selecting this completion.
	// When an edit is provided the value of `insertText` is ignored.
	//
	// *Note:* The range of the edit must be a single line range and it must
	// contain the position at which completion has been requested.
	//
	// Most editors support two different operations when accepting a completion
	// item. One is to insert a completion text and the other is to replace an
	// existing text with a completion text. Since this can usually not be
	// predetermined by a server it can report both ranges. Clients need to
	// signal support for `InsertReplaceEdit`s via the
	// `textDocument.completion.completionItem.insertReplaceSupport` client
	// capability property.
	//
	// *Note 1:* The text edit's range as well as both ranges from an insert
	// replace edit must be a [single line] and they must contain the position
	// at which completion has been requested.
	// *Note 2:* If an `InsertReplaceEdit` is returned the edit's insert range
	// must be a prefix of the edit's replace range, that means it must be
	// contained and starting at the same position.
	//
	// @since 3.16.0 additional type `InsertReplaceEdit`
	//
	// TextEdit | InsertReplaceEdit | nil
	TextEdit any `json:"textEdit,omitempty"`

	// The edit text used if the completion item is part of a CompletionList and
	// CompletionList defines an item default for the text edit range.
	//
	// Clients will only honor this property if they opt into completion list
	// item defaults using the capability `completionList.itemDefaults`.
	//
	// If not provided and a list's default range is provided the label
	// property is used as a text.
	//
	// @since 3.17.0
	TextEditText *string `json:"textEditText,omitempty"`

	// An optional array of additional text edits that are applied when
	// selecting this completion. Edits must not overlap (including the same
	// insert position) with the main edit nor with themselves.
	//
	// Additional text edits should be used to change text unrelated to the
	// current cursor position (for example adding an import statement at the
	// top of the file if the completion item will insert an unqualified type).
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`

	// An optional set of characters that when pressed while this completion is
	// active will accept it first and then type that character. Note that all
	// commit characters should have `length=1` and that superfluous characters
	// will be ignored.
	CommitCharacters []string `json:"commitCharacters,omitempty"`

	// An optional command that is executed *after* inserting this completion.
	// *Note* that additional modifications to the current document should be
	// described with the additionalTextEdits-property.
	Command *Command `json:"command,omitempty"`

	// A data entry field that is preserved on a completion item between
	// a completion and a completion resolve request.
	Data LSPAny `json:"data,omitempty"`
}

type completionItemIntermediary struct {
	Label        string                      `json:"label"`
	LabelDetails *CompletionItemLabelDetails `json:"labelDetails,omitempty"`
	Kind         *CompletionItemKind         `json:"kind,omitempty"`
	Tags         []CompletionItemTag         `json:"tags,omitempty"`
	Detail       *string                     `json:"detail,omitempty"`
	// string | MarkupContent | nil
	Documentation    json.RawMessage   `json:"documentation,omitempty"`
	Deprecated       *bool             `json:"deprecated,omitempty"`
	Preselect        *bool             `json:"preselect,omitempty"`
	SortText         *string           `json:"sortText,omitempty"`
	FilterText       *string           `json:"filterText,omitempty"`
	InsertText       *string           `json:"insertText,omitempty"`
	InsertTextFormat *InsertTextFormat `json:"insertTextFormat,omitempty"`
	InsertTextMode   *InsertTextMode   `json:"insertTextMode,omitempty"`
	// TextEdit | InsertReplaceEdit | nil
	TextEdit            json.RawMessage `json:"textEdit,omitempty"`
	TextEditText        *string         `json:"textEditText,omitempty"`
	AdditionalTextEdits []TextEdit      `json:"additionalTextEdits,omitempty"`
	CommitCharacters    []string        `json:"commitCharacters,omitempty"`
	Command             *Command        `json:"command,omitempty"`
	Data                LSPAny          `json:"data,omitempty"`
}

// Fulfils the json.Unmarhaller interface.
func (i *CompletionItem) UnmarshalJSON(data []byte) error {

	var value completionItemIntermediary

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	i.Label = value.Label
	i.LabelDetails = value.LabelDetails
	i.Kind = value.Kind
	i.Tags = value.Tags
	i.Detail = value.Detail
	i.Deprecated = value.Deprecated
	i.Preselect = value.Preselect
	i.SortText = value.SortText
	i.FilterText = value.FilterText
	i.InsertText = value.InsertText
	i.InsertTextFormat = value.InsertTextFormat
	i.InsertTextMode = value.InsertTextMode
	i.TextEditText = value.TextEditText
	i.AdditionalTextEdits = value.AdditionalTextEdits
	i.CommitCharacters = value.CommitCharacters
	i.Command = value.Command
	i.Data = value.Data

	// Documentation and text edit fields are optional,
	// so on failure to unmarshal, nil values will be set.
	i.unmarshalDocumentation(&value)
	i.unmarshalTextEdit(&value)

	return nil
}

func (i *CompletionItem) unmarshalDocumentation(value *completionItemIntermediary) error {
	var strVal string
	if err := json.Unmarshal(value.Documentation, &strVal); err == nil {
		i.Documentation = strVal
		return nil
	}

	var markupContent MarkupContent
	if err := json.Unmarshal(value.Documentation, &markupContent); err == nil && markupContent.Kind != "" {
		i.Documentation = markupContent
	} else {
		return err
	}

	return nil
}

func (i *CompletionItem) unmarshalTextEdit(value *completionItemIntermediary) error {
	var textEditVal TextEdit
	if err := json.Unmarshal(value.TextEdit, &textEditVal); err == nil && textEditVal.Range != nil {
		i.TextEdit = textEditVal
		return nil
	}

	var insertReplaceEdit InsertReplaceEdit
	err := json.Unmarshal(value.TextEdit, &insertReplaceEdit)
	if err == nil && (insertReplaceEdit.Replace != nil || insertReplaceEdit.Insert != nil) {
		i.TextEdit = insertReplaceEdit
	}

	return err
}

// InsertReplaceEdit is a special text edit to provide an insert and a replace operation.
//
// @since 3.16.0
type InsertReplaceEdit struct {
	// The string to be inserted.
	NewText string `json:"newText"`

	// The range if the insert is requested.
	Insert *Range `json:"insert"`

	// The range if the replace is requested.
	Replace *Range `json:"replace"`
}

// CompletionItemLabelDetails provides additional details for a completion
// item label.
//
// @since 3.17.0
type CompletionItemLabelDetails struct {
	// An optional string which is rendered less prominently directly after
	// {@link CompletionItem.label label}, without any spacing. Should be
	// used for function signatures or type annotations.
	Detail *string `json:"detail,omitempty"`

	// An optional string which is rendered less prominently after
	// {@link CompletionItemLabelDetails.detail}. Should be used for fully qualified
	// names or file path.
	Description *string `json:"description,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItem_resolve

const MethodCompletionItemResolve = Method("completionItem/resolve")

// CompletionItemResolveHandlerFunc is the function signature for the
// completionItem/resolve request handler that can be registered for a
// language server.
type CompletionItemResolveHandlerFunc func(
	ctx *common.LSPContext,
	params *CompletionItem,
) (*CompletionItem, error)
