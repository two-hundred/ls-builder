package lsp

import (
	"encoding/json"
)

// CompletionClientCapabilities describes the capabilities of a client
// for completion requests.
type CompletionClientCapabilities struct {
	// Whether completion supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// The client supports the following `CompletionItem` specific capabilities.
	CompletionItem *CompletionItemCapabilities `json:"completionItem,omitempty"`

	// The completion item kind values the client support
	CompletionItemKind *CompletionItemKindCapabilities `json:"completionItemKind,omitempty"`

	// The client supports to send additional context information for a
	// `textDocument/completion` request.
	ContextSupport *bool `json:"contextSupport,omitempty"`

	// The client's default when the completion item doesn't provide a
	// `insertTextMode` property.
	//
	// @since 3.17.0
	InsertTextMode *CompletionItemInsertTextMode `json:"insertTextMode,omitempty"`

	// The client supports the following `CompletionList` specific
	// capabilities.
	//
	// @since 3.17.0
	CompletionList *CompletionListCapabilities `json:"completionList,omitempty"`
}

// CompletionListCapabilities describes the capabilities of a client
// for completion lists.
type CompletionListCapabilities struct {
	// The client supports the following itemDefaults on
	// a completion list.
	//
	// The value lists the supported property names of the
	// `CompletionList.itemDefaults` object. If omitted
	// no properties are supported.
	//
	// @since 3.17.0
	ItemDefaults []string `json:"itemDefaults,omitempty"`
}

// CompletionItemKindCapabilities describes the capabilities of a client
// for completion item kinds.
type CompletionItemKindCapabilities struct {
	// The completion item kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	//
	// If this property is not present the client only supports
	// the completion items kinds from `Text` to `Reference` as defined in
	// the initial version of the protocol.
	ValueSet []CompletionItemKind `json:"valueSet,omitempty"`
}

// CompletionItemKind represents a kind of completion item.
type CompletionItemKind = Integer

const (
	// CompletionItemKindText is a plain text completion item.
	CompletionItemKindText CompletionItemKind = 1

	// CompletionItemKindMethod is a method completion item.
	CompletionItemKindMethod CompletionItemKind = 2

	// CompletionItemKindFunction is a function completion item.
	CompletionItemKindFunction CompletionItemKind = 3

	// CompletionItemKindConstructor is a constructor completion item.
	CompletionItemKindConstructor CompletionItemKind = 4

	// CompletionItemKindField is a field completion item.
	CompletionItemKindField CompletionItemKind = 5

	// CompletionItemKindVariable is a variable completion item.
	CompletionItemKindVariable CompletionItemKind = 6

	// CompletionItemKindClass is a class completion item.
	CompletionItemKindClass CompletionItemKind = 7

	// CompletionItemKindInterface is an interface completion item.
	CompletionItemKindInterface CompletionItemKind = 8

	// CompletionItemKindModule is a module completion item.
	CompletionItemKindModule CompletionItemKind = 9

	// CompletionItemKindProperty is a property completion item.
	CompletionItemKindProperty CompletionItemKind = 10

	// CompletionItemKindUnit is a unit completion item.
	CompletionItemKindUnit CompletionItemKind = 11

	// CompletionItemKindValue is a value completion item.
	CompletionItemKindValue CompletionItemKind = 12

	// CompletionItemKindEnum is an enum completion item.
	CompletionItemKindEnum CompletionItemKind = 13

	// CompletionItemKindKeyword is a keyword completion item.
	CompletionItemKindKeyword CompletionItemKind = 14

	// CompletionItemKindSnippet is a snippet completion item.
	CompletionItemKindSnippet CompletionItemKind = 15

	// CompletionItemKindColor is a color completion item.
	CompletionItemKindColor CompletionItemKind = 16

	// CompletionItemKindFile is a file completion item.
	CompletionItemKindFile CompletionItemKind = 17

	// CompletionItemKindReference is a reference completion item.
	CompletionItemKindReference CompletionItemKind = 18

	// CompletionItemKindFolder is a folder completion item.
	CompletionItemKindFolder CompletionItemKind = 19

	// CompletionItemKindEnumMember is an enum member completion item.
	CompletionItemKindEnumMember CompletionItemKind = 20

	// CompletionItemKindConstant is a constant completion item.
	CompletionItemKindConstant CompletionItemKind = 21

	// CompletionItemKindStruct is a struct completion item.
	CompletionItemKindStruct CompletionItemKind = 22

	// CompletionItemKindEvent is an event completion item.
	CompletionItemKindEvent CompletionItemKind = 23

	// CompletionItemKindOperator is an operator completion item.
	CompletionItemKindOperator CompletionItemKind = 24

	// CompletionItemKindTypeParameter is a type parameter completion item.
	CompletionItemKindTypeParameter CompletionItemKind = 25
)

// CompletionItemCapabilities describes the capabilities of a client
// for completion items.
type CompletionItemCapabilities struct {
	// Client supports snippets as insert text.
	//
	// A snippet can define tab stops and placeholders with `$1`, `$2`
	// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
	// the end of the snippet. Placeholders with equal identifiers are
	// linked, that is typing in one will update others too.
	SnippetSupport *bool `json:"snippetSupport,omitempty"`

	// Client supports commit characters on a completion item.
	CommitCharactersSupport *bool `json:"commitCharactersSupport,omitempty"`

	// Client supports the follow content formats for the documentation
	// property. The order describes the preferred format of the client.
	DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

	// Client supports the deprecated property on a completion item.
	DeprecatedSupport *bool `json:"deprecatedSupport,omitempty"`

	// Client supports the preselect property on a completion item.
	PreselectSupport *bool `json:"preselectSupport,omitempty"`

	// Client supports the tag property on a completion item. Clients
	// supporting tags have to handle unknown tags gracefully. Clients
	// especially need to preserve unknown tags when sending a completion
	// item back to the server in a resolve call.
	//
	// @since 3.15.0
	TagSupport *CompletionItemTagSupport `json:"tagSupport,omitempty"`

	// Client supports insert replace edit to control different behavior if
	// a completion item is inserted in the text or should replace text.
	//
	// @since 3.16.0
	InsertReplaceSupport *bool `json:"insertReplaceSupport,omitempty"`

	// Indicates which properties a client can resolve lazily on a
	// completion item. Before version 3.16.0 only the predefined properties
	// `documentation` and `detail` could be resolved lazily.
	//
	// @since 3.16.0
	ResolveSupport *CompletionItemResolveSupport `json:"resolveSupport,omitempty"`

	// The client supports the `insertTextMode` property on
	// a completion item to override the whitespace handling
	// mode as defined by the client (see `insertTextMode`).
	//
	// @since 3.16.0
	InsertTextModeSupport *CompletionItemInsertTextModeSupport `json:"insertTextModeSupport,omitempty"`

	// The client has support for completion item label
	// details (see also `CompletionItemLabelDetails`).
	//
	// @since 3.17.0
	LabelDetailsSupport *bool `json:"labelDetailsSupport,omitempty"`
}

// CompletionItemCapabilities describes the capabilities of a client
// for completion items with respect to insert text mode.
type CompletionItemInsertTextModeSupport struct {
	ValueSet []CompletionItemInsertTextMode `json:"valueSet"`
}

// CompletionItemInsertTextMode determines how whitespace
// and indentation is handled during completion item insertion.
//
// @since 3.16.0
type CompletionItemInsertTextMode = Integer

const (
	// CompletionItemInsertTextModeAsIs means that the
	// insertion or string replacement is taken as it is.
	// If the value is multi line, the lines below the cursor
	// will be inserted using the indentation defined in the string value.
	// The client will not apply any kind of adjustments to the string.
	CompletionItemInsertTextModeAsIs CompletionItemInsertTextMode = 1

	// CompletionItemInsertTextModeAdjustIndentation means that
	// the editor adjusts leading whitespace of new lines so that tehy match
	// the indentation up to the cursor of the line for which the item is accepted.
	//
	// Consider a line lik this: <2tabs><cursor><3tabs>foo. Accepting a
	// multi line completion item is indentned using 2 tabs and all
	// following lines inserted will be indented using 2 tabs as well.
	CompletionItemInsertTextModeAdjustIndentation CompletionItemInsertTextMode = 2
)

// CompletionItemResolveSupport describes the capabilities of a client
// for resolving completion items lazily.
type CompletionItemResolveSupport struct {
	// The properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}

// CompletionItemTagSupport describes the capabilities of a client
// for completion item tags.
type CompletionItemTagSupport struct {
	// The tags supported by the client.
	ValueSet []CompletionItemTag `json:"valueSet"`
}

// CompletionItemTags are extra annotations that tweak the rendering
// of a completion item.
//
// @since 3.15.0
type CompletionItemTag = Integer

const (
	// CompletionItemTagDeprecated renders a completion
	// as obsolete, usually using a strike-out.
	CompletionItemTagDeprecated CompletionItemTag = 1
)

// HoverClientCapabilities describes the capabilities of a client
// for hover requests.
type HoverClientCapabilities struct {
	// Whether hover supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// SignatureHelpClientCapabilities describes the capabilities of a client
// for signature help requests.
type SignatureHelpClientCapabilities struct {
	// Whether signature help supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// The client supports the following `SignatureInformation`
	// specific properties.
	SignatureInformation *SignatureInformationCapabilities `json:"signatureInformation,omitempty"`

	// The client supports to send additional context information for a
	// `textDocument/signatureHelp` request. A client that opts into
	// contextSupport will also support the `retriggerCharacters` on
	// `SignatureHelpOptions`.
	//
	// @since 3.15.0
	ContextSupport *bool `json:"contextSupport,omitempty"`
}

// SignatureInformationCapabilities describes the capabilities of a client
// for signature information.
type SignatureInformationCapabilities struct {
	// Client supports the follow content formats for the documentation
	// property. The order describes the preferred format of the client.
	DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

	// Client capabilities specific to parameter information.
	ParameterInformation *ParameterInformationCapabilities `json:"parameterInformation,omitempty"`

	// The client supports the `activeParameter` property on
	// a `SignatureInformation` literal.
	ActiveParameterSupport *bool `json:"activeParameterSupport,omitempty"`
}

// ParameterInformationCapabilities describes the capabilities of a client
// for parameter information.
type ParameterInformationCapabilities struct {
	// The client supports processing label offsets instead of a
	// simple label string.
	//
	// @since 3.14.0
	LabelOffsetSupport *bool `json:"labelOffsetSupport,omitempty"`
}

// DeclarationClientCapabilities describes the capabilities of a client
// for goto declaration requests.
type DeclarationClientCapabilities struct {
	// Whether declaration supports dynamic registration. If this is set to
	// `true` the client supports the new `DeclarationRegistrationOptions`
	// return value for the corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// The client supports additional metadata in the form of declaration links.
	LinkSupport *bool `json:"linkSupport,omitempty"`
}

// DefinitionClientCapabilities describes the capabilities of a client
// for goto definition requests.
type DefinitionClientCapabilities struct {
	// Whether definition supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// The client supports additional metadata in the form of definition links.
	LinkSupport *bool `json:"linkSupport,omitempty"`
}

// TypeDefinitionClientCapabilities describes the capabilities of a client
// for goto type definition requests.
type TypeDefinitionClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to
	// `true` the client supports the new `TypeDefinitionRegistrationOptions`
	// return value for the corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// The client supports additional metadata in the form of definition links.
	//
	// @since 3.14.0
	LinkSupport *bool `json:"linkSupport,omitempty"`
}

// ImplementationClientCapabilities describes the capabilities of a client
// for goto implementation requests.
type ImplementationClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to
	// `true` the client supports the new `ImplementationRegistrationOptions`
	// return value for the corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// The client supports additional metadata in the form of definition links.
	//
	// @since 3.14.0
	LinkSupport *bool `json:"linkSupport,omitempty"`
}

// ReferencesClientCapabilities describes the capabilities of a client
// for find references requests.
type ReferenceClientCapabilities struct {
	// Whether references supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// DocumentHighlightClientCapabilities describes the capabilities of a client
// for document highlight requests.
type DocumentHighlightClientCapabilities struct {
	// Whether document highlight supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// DocumentSymbolClientCapabilities describes the capabilities of a client
// for document symbol requests.
type DocumentSymbolClientCapabilities struct {
	// Whether document symbol supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// Specific capabilities for the `SymbolKind` in the
	// `textDocument/documentSymbol` request.
	SymbolKind *SymbolKindCapabilities `json:"symbolKind,omitempty"`

	// The client supports hierarchical document symbols.
	HierarchicalDocumentSymbolSupport *bool `json:"hierarchicalDocumentSymbolSupport,omitempty"`

	// The client supports tags on `SymbolInformation`. Tags are supported on
	// `DocumentSymbol` if `hierarchicalDocumentSymbolSupport` is set to true.
	// Clients supporting tags have to handle unknown tags gracefully.
	//
	// @since 3.16.0
	TagSupport *SymbolTagSupport `json:"tagSupport,omitempty"`

	// The client supports an additional label presented in the UI when
	// registering a document symbol provider.
	//
	// @since 3.16.0
	LabelSupport *bool `json:"labelSupport,omitempty"`
}

// CodeActionClientCapabilities describes the capabilities of a client
// for code action requests.
type CodeActionClientCapabilities struct {
	// Whether code action supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// The client supports code action literals as a valid
	// response of the `textDocument/codeAction` request.
	//
	// @since 3.8.0
	CodeActionLiteralSupport *CodeActionLiteralSupport `json:"codeActionLiteralSupport,omitempty"`

	// Whether code action supports the `isPreferred` property.
	//
	// @since 3.15.0
	IsPreferredSupport *bool `json:"isPreferredSupport,omitempty"`

	// Whether code action supports the `disabled` property.
	//
	// @since 3.16.0
	DisabledSupport *bool `json:"disabledSupport,omitempty"`

	// Whether code action supports the `data` property which is
	// preserved between a `textDocument/codeAction` and a
	// `codeAction/resolve` request.
	//
	// @since 3.16.0
	DataSupport *bool `json:"dataSupport,omitempty"`

	// Whether the client supports resolving additional code action
	// properties via a separate `codeAction/resolve` request.
	//
	// @since 3.16.0
	ResolveSupport *CodeActionResolveSupport `json:"resolveSupport,omitempty"`

	// Whether the client honors the change annotations in
	// text edits and resource operations returned via the
	// `CodeAction#edit` property by for example presenting
	// the workspace edit in the user interface and asking
	// for confirmation.
	//
	// @since 3.16.0
	HonorsChangeAnnotations *bool `json:"honorsChangeAnnotations,omitempty"`
}

// CodeActionResolveSupport describes the capabilities of a client
// for resolving code actions lazily.
type CodeActionResolveSupport struct {
	// The properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}

// CodeActionLiteralSupport describes the capabilities of a client
// for code action literals.
type CodeActionLiteralSupport struct {
	// The code action kind is supported with the following value
	// set.
	CodeActionKind *CodeActionKindCapabilities `json:"codeActionKind,omitempty"`
}

// CodeActionKindCapabilities describes the capabilities of a client
// for code action kinds.
type CodeActionKindCapabilities struct {
	// The code action kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	ValueSet []CodeActionKind `json:"valueSet,omitempty"`
}

// CodeLensClientCapabilities describes the capabilities of a client
// for code lens requests.
type CodeLensClientCapabilities struct {
	// Whether code lens supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// DocumentLinkClientCapabilities describes the capabilities of a client
// for document link requests.
type DocumentLinkClientCapabilities struct {
	// Whether document link supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// Whether the client supports the `tooltip` property on `DocumentLink`.
	//
	// @since 3.15.0
	TooltipSupport *bool `json:"tooltipSupport,omitempty"`
}

// DocumentColorClientCapabilities describes the capabilities of a client
// for document color requests.
type DocumentColorClientCapabilities struct {
	//  Whether document color supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// DocumentFormattingClientCapabilities describes the capabilities of a client
// for document formatting requests.
type DocumentFormattingClientCapabilities struct {
	// Whether formatting supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// DocumentRangeFormattingClientCapabilities describes the capabilities of a client
// for document range formatting requests.
type DocumentRangeFormattingClientCapabilities struct {
	// Whether formatting supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// DocumentOnTypeFormattingClientCapabilities describes the capabilities of a client
// for document on type formatting requests.
type DocumentOnTypeFormattingClientCapabilities struct {
	// Whether on type formatting supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// RenameClientCapabilities describes the capabilities of a client
// for rename requests.
type RenameClientCapabilities struct {
	// Whether rename supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// Client supports testing for validity of rename operations
	// before execution.
	//
	// @since version 3.12.0
	PrepareSupport *bool `json:"prepareSupport,omitempty"`

	// Client supports the default behavior result
	// (`{ defaultBehavior: boolean }`).
	//
	// The value indicates the default behavior used by the
	// client.
	//
	// @since version 3.16.0
	PrepareSupportDefaultBehavior *PrepareSupportDefaultBehavior `json:"prepareSupportDefaultBehavior,omitempty"`

	// Whether the client honors the change annotations in
	// text edits and resource operations returned via the
	// rename request's workspace edit by for example presenting
	// the workspace edit in the user interface and asking
	// for confirmation.
	//
	// @since 3.16.0
	HonorsChangeAnnotations *bool `json:"honorsChangeAnnotations,omitempty"`
}

// PrepareSupportDefaultBehavior is the default behavior used by the client
// for testing validity of operations such as rename.
type PrepareSupportDefaultBehavior = Integer

const (
	// PrepareSupportDefaultBehaviorIdentifier
	// means the client's default behavior is to select the identifier
	// according to the language's syntax rule.
	PrepareSupportDefaultBehaviorIdentifier PrepareSupportDefaultBehavior = 1
)

// FoldingRangeClientCapabilities describes the capabilities of a client
// for folding range requests.
type FoldingRangeClientCapabilities struct {
	// Whether implementation supports dynamic registration for folding range
	// providers. If this is set to `true` the client supports the new
	// `FoldingRangeRegistrationOptions` return value for the corresponding
	// server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// The maximum number of folding ranges that the client prefers to receive
	// per document. The value serves as a hint, servers are free to follow the
	// limit.
	RangeLimit *UInteger `json:"rangeLimit,omitempty"`

	// If set, the client signals that it only supports folding complete lines.
	// If set, client will ignore specified `startCharacter` and `endCharacter`
	// properties in a FoldingRange.
	LineFoldingOnly *bool `json:"lineFoldingOnly,omitempty"`

	// Specific options for the folding range kind.
	//
	// @since 3.17.0
	FoldingRangeKind *FoldingRangeKindCapabilities `json:"foldingRangeKind,omitempty"`

	// Specific options for the folding range.
	// @since 3.17.0
	FoldingRange *FoldingRangeCapabilities `json:"foldingRange,omitempty"`
}

// FoldingRangeCapabilities describes the additiona options
// for the capabilities of a client for folding ranges.
type FoldingRangeCapabilities struct {
	// If set, the client signals that it supports setting collapsedText on
	// folding ranges to display custom labels instead of the default text.
	//
	// @since 3.17.0
	CollapsedText *bool `json:"collapsedText,omitempty"`
}

// FoldingRangeKindCapabilities describes the capabilities of a client
// for folding range kinds.
type FoldingRangeKindCapabilities struct {
	// The folding range kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	ValueSet []FoldingRangeKind `json:"valueSet,omitempty"`
}

// FoldingRangeKind represents predefined
// range kinds.
type FoldingRangeKind = string

const (
	// FoldingRangeKindComment represents a folding range
	// for a comment.
	FoldingRangeKindComment FoldingRangeKind = "comment"

	// FoldingRangeKindImports represents a folding range
	// for imports.
	FoldingRangeKindImports FoldingRangeKind = "imports"

	// FoldingRangeKindRegion represents a folding range
	// for a region (e.g. `#region`).
	FoldingRangeKindRegion FoldingRangeKind = "region"
)

// SelectionRangeClientCapabilities describes the capabilities of a client
// for selection range requests.
type SelectionRangeClientCapabilities struct {
	// Whether implementation supports dynamic registration for selection range
	// providers. If this is set to `true` the client supports the new
	// `SelectionRangeRegistrationOptions` return value for the corresponding
	// server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// LinkedEditingRangeClientCapabilities describes the capabilities of a client
// for linked editing range requests.
type LinkedEditingRangeClientCapabilities struct {
	// Whether the implementation supports dynamic registration.
	// If this is set to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// CallHierarchyClientCapabilities describes the capabilities of a client
// for call hierarchy requests.
type CallHierarchyClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to
	// `true` the client supports the new `(TextDocumentRegistrationOptions &
	// StaticRegistrationOptions)` return value for the corresponding server
	// capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// SemanticTokensClientCapabilities describes the capabilities of a client
// for semantic tokens requests.
type SemanticTokensClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to
	// `true` the client supports the new `(TextDocumentRegistrationOptions &
	// StaticRegistrationOptions)` return value for the corresponding server
	// capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// Which requests the client supports and might send to the server
	// depending on the server's capability. Please note that clients might not
	// show semantic tokens or degrade some of the user experience if a range
	// or full request is advertised by the client but not provided by the
	// server. If for example the client capability `requests.full` and
	// `request.range` are both set to true but the server only provides a
	// range provider the client might not render a minimap correctly or might
	// even decide to not show any semantic tokens at all.
	Requests SemanticTokensRequests `json:"requests"`

	// The token types that the client supports.
	TokenTypes []string `json:"tokenTypes"`

	// The token modifiers that the client supports.
	TokenModifiers []string `json:"tokenModifiers"`

	// The formats the clients supports.
	Formats []TokenFormat `json:"formats"`

	// Whether the client supports tokens that can overlap each other.
	OverlappingTokenSupport *bool `json:"overlappingTokenSupport,omitempty"`

	// Whether the client supports tokens that can span multiple lines.
	MultilineTokenSupport *bool `json:"multilineTokenSupport,omitempty"`

	// Whether the client allows the server to actively cancel a
	// semantic token request, e.g. supports returning
	// ErrorCodes.ServerCancelled. If a server does the client
	// needs to retrigger the request.
	//
	// @since 3.17.0
	ServerCancelSupport *bool `json:"serverCancelSupport,omitempty"`

	// Whether the client uses semantic tokens to augment existing
	// syntax tokens. If set to `true` client side created syntax
	// tokens and semantic tokens are both used for colorization. If
	// set to `false` the client only uses the returned semantic tokens
	// for colorization.
	//
	// If the value is `undefined` then the client behavior is not
	// specified.
	//
	// @since 3.17.0
	AugmentsSyntaxTokens *bool `json:"augmentsSyntaxTokens,omitempty"`
}

type semanticTokenRequestsIntermediate struct {
	Range json.RawMessage `json:"range,omitempty"`
	Full  json.RawMessage `json:"full,omitempty"`
}

type semanticTokensClientCapabilitiesIntermediate struct {
	DynamicRegistration     *bool                             `json:"dynamicRegistration,omitempty"`
	Requests                semanticTokenRequestsIntermediate `json:"requests"`
	TokenTypes              []string                          `json:"tokenTypes"`
	TokenModifiers          []string                          `json:"tokenModifiers"`
	Formats                 []TokenFormat                     `json:"formats"`
	OverlappingTokenSupport *bool                             `json:"overlappingTokenSupport,omitempty"`
	MultilineTokenSupport   *bool                             `json:"multilineTokenSupport,omitempty"`
	ServerCancelSupport     *bool                             `json:"serverCancelSupport,omitempty"`
	AugmentsSyntaxTokens    *bool                             `json:"augmentsSyntaxTokens,omitempty"`
}

// Fulfils the json.Unmarshaler interface.
func (c *SemanticTokensClientCapabilities) UnmarshalJSON(data []byte) error {
	var intermediate semanticTokensClientCapabilitiesIntermediate
	if err := json.Unmarshal(data, &intermediate); err != nil {
		return err
	}

	c.DynamicRegistration = intermediate.DynamicRegistration
	c.TokenTypes = intermediate.TokenTypes
	c.TokenModifiers = intermediate.TokenModifiers
	c.Formats = intermediate.Formats
	c.OverlappingTokenSupport = intermediate.OverlappingTokenSupport
	c.MultilineTokenSupport = intermediate.MultilineTokenSupport
	c.ServerCancelSupport = intermediate.ServerCancelSupport
	c.AugmentsSyntaxTokens = intermediate.AugmentsSyntaxTokens

	if intermediate.Requests.Range != nil {
		var rangeBool bool
		if err := json.Unmarshal(intermediate.Requests.Range, &rangeBool); err == nil {
			c.Requests.Range = rangeBool
		} else {
			var rangeStruct struct{}
			if err := json.Unmarshal(intermediate.Requests.Range, &rangeStruct); err == nil {
				c.Requests.Range = rangeStruct
			} else {
				return err
			}
		}
	}
	err := c.unmarshalSemanticTokensRequestsRange(intermediate)
	if err != nil {
		return err
	}

	err = c.unmarshalSemanticTokensRequestsFull(intermediate)
	return err
}

func (c *SemanticTokensClientCapabilities) unmarshalSemanticTokensRequestsRange(
	intermediate semanticTokensClientCapabilitiesIntermediate,
) error {
	if intermediate.Requests.Range != nil {
		var rangeBool bool
		if err := json.Unmarshal(intermediate.Requests.Range, &rangeBool); err == nil {
			c.Requests.Range = rangeBool
		} else {
			var rangeStruct struct{}
			if err := json.Unmarshal(intermediate.Requests.Range, &rangeStruct); err == nil {
				c.Requests.Range = rangeStruct
			} else {
				return err
			}
		}
	}

	return nil
}

func (c *SemanticTokensClientCapabilities) unmarshalSemanticTokensRequestsFull(
	intermediate semanticTokensClientCapabilitiesIntermediate,
) error {
	if intermediate.Requests.Full != nil {
		var fullBool bool
		if err := json.Unmarshal(intermediate.Requests.Full, &fullBool); err == nil {
			c.Requests.Full = fullBool
		} else {
			var fullStruct SemanticDelta
			if err := json.Unmarshal(intermediate.Requests.Full, &fullStruct); err == nil {
				c.Requests.Full = fullStruct
			} else {
				return err
			}
		}
	}

	return nil
}

// TokenFormat represents additional token format capabilities
// to allow future extension of the format.
type TokenFormat = string

const (
	// TokenFormatRelative represents a relative token format.
	TokenFormatRelative TokenFormat = "relative"
)

// SemanticTokensRequests describes requests that a client
// can send to a server for semantic tokens.
type SemanticTokensRequests struct {
	// The client will send the `textDocument/semanticTokens/range` request
	// if the server provides a corresponding handler.
	//
	// boolean | struct{} | nil
	Range any `json:"range,omitempty"`

	// The client will send the `textDocument/semanticTokens/full` request
	// if the server provides a corresponding handler.
	//
	// boolean | SemanticDelta | nil
	Full any `json:"full,omitempty"`
}

// MonikerClientCapabilities describes the capabilities of a client
// for moniker requests.
type MonikerClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to
	// `true` the client supports the new `(TextDocumentRegistrationOptions &
	// StaticRegistrationOptions)` return value for the corresponding server
	// capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// InlayHintClientCapabilities describes the capabilities of a client
// for inlay hint requests.
type InlayHintClientCapabilities struct {
	// Whether inlay hints support dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// Indicates which properties a client can resolve lazily on an inlay
	// hint.
	ResolveSupport *InlayHintResolveSupport `json:"resolveSupport,omitempty"`
}

// InlayHintResolveSupport describes the capabilities of a client
// for resolving inlay hints lazily.
type InlayHintResolveSupport struct {
	// The properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}

// InlineValueClientCapabilities describes the capabilities of a client
// for inline value requests.
type InlineValueClientCapabilities struct {
	// Whether implementation supports dynamic registration for inline
	// value providers.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// TypeHierarchyClientCapabilities describes the capabilities of a client
// for type hierarchy requests.
type TypeHierarchyClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to
	// `true` the client supports the new `(TextDocumentRegistrationOptions &
	// StaticRegistrationOptions)` return value for the corresponding server
	// capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// CompletionOptions provides server capability options for completion requests.
type CompletionOptions struct {
	WorkDoneProgressOptions

	// The additional characters, beyond the defaults provided by the client (typically
	// [a-zA-Z]), that should automatically trigger a completion request. For example
	// `.` in JavaScript represents the beginning of an object property or method and is
	// thus a good candidate for triggering a completion request.
	//
	// Most tools trigger a completion request automatically without explicitly
	// requesting it using a keyboard shortcut (e.g. Ctrl+Space). Typically they
	// do so when the user starts to type an identifier. For example if the user
	// types `c` in a JavaScript file code complete will automatically pop up
	// present `console` besides others as a completion item. Characters that
	// make up identifiers don't need to be listed here.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	// The list of all possible characters that commit a completion. This field
	// can be used if clients don't support individual commit characters per
	// completion item. See client capability
	// `completion.completionItem.commitCharactersSupport`.
	//
	// If a server provides both `allCommitCharacters` and commit characters on
	// an individual completion item the ones on the completion item win.
	//
	// @since 3.2.0
	AllCommitCharacters []string `json:"allCommitCharacters,omitempty"`

	// The server provides support to resolve additional
	// information for a completion item.
	ResolveProvider *bool `json:"resolveProvider,omitempty"`

	// The server supports the following `CompletionItem` specific
	// capabilities.
	//
	// @since 3.17.0
	CompletionItem *CompletionOptionsItem `json:"completionItem,omitempty"`
}

// CompletionItemOptions provides server capability options for completion items.
type CompletionOptionsItem struct {
	// The server has support for completion item label
	// details (see also `CompletionItemLabelDetails`) when receiving
	// a completion item in a resolve call.
	//
	// @since 3.17.0
	LabelDetailsSupport *bool `json:"labelDetailsSupport,omitempty"`
}

// SignatureHelpOptions provides server capability options for signature help requests.
type SignatureHelpOptions struct {
	WorkDoneProgressOptions

	// The characters that trigger signature help
	// automatically.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	// List of characters that re-trigger signature help.
	//
	// These trigger characters are only active when signature help is already
	// showing. All trigger characters are also counted as re-trigger
	// characters.
	//
	// @since 3.15.0
	RetriggerCharacters []string `json:"retriggerCharacters,omitempty"`
}

// CodeLensOptions provides server capability options for code lens requests.
type CodeLensOptions struct {
	WorkDoneProgressOptions

	// Code lens has a resolve provider as well.
	ResolveProvider *bool `json:"resolveProvider,omitempty"`
}

// DocumentLinkOptions provides server capability options for document link requests.
type DocumentLinkOptions struct {
	WorkDoneProgressOptions

	// Code lens has a resolve provider as well.
	ResolveProvider *bool `json:"resolveProvider,omitempty"`
}

// DocumentOnTypeFormattingOptions provides server capability options for
// document on type formatting requests.
type DocumentOnTypeFormattingOptions struct {
	// A character on which formatting should be triggered, like `}`.
	FirstTriggerCharacter string `json:"firstTriggerCharacter"`

	// More trigger characters.
	MoreTriggerCharacter []string `json:"moreTriggerCharacter,omitempty"`
}

func unmarshalLanguageFeatureServerCapabilities(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if err := unmarshalLanguageFeatureSet1ServerCapabilities(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalLanguageFeatureSet2ServerCapabilities(serverCapabilities, intermediate); err != nil {
		return err
	}

	return nil
}

func unmarshalLanguageFeatureSet1ServerCapabilities(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if err := unmarshalServerCapabilityHoverProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityDeclarationProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityDefinitionProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityTypeDefinitionProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityImplementationProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityReferencesProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityDocumentHighlightProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityDocumentSymbolProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityCodeActionProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityColorProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	return nil
}

func unmarshalLanguageFeatureSet2ServerCapabilities(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if err := unmarshalServerCapabilityDocumentFormattingProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityDocumentRangeFormattingProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityRenameProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityFoldingRangeProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilitySelectionRangeProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityLinkedEditingRangeProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityCallHierarchyProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilitySemanticTokensProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityMonikerProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityTypeHierarchyProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityInlineValueProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityInlayHintProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityDiagnosticProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityWorkspaceSymbolProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	return nil
}

// HoverOptions provides server capability options for hover requests.
type HoverOptions struct {
	WorkDoneProgressOptions
}

// unmarshals the HoverProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityHoverProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.HoverProvider == nil {
		return nil
	}

	var optVal HoverOptions
	if err := json.Unmarshal(intermediate.HoverProvider, &optVal); err == nil {
		serverCapabilities.HoverProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.HoverProvider, &boolVal); err == nil {
			serverCapabilities.HoverProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DeclarationOptions provides server capability options for goto declaration requests.
type DeclarationOptions struct {
	WorkDoneProgressOptions
}

// DeclarationRegistrationOptions provides server capability registration
// options for goto declaration requests.
type DeclarationRegistrationOptions struct {
	DeclarationOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the DeclarationProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// DeclarationRegistrationOptions is a superset of DeclarationOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityDeclarationProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.DeclarationProvider == nil {
		return nil
	}

	var optRegVal DeclarationRegistrationOptions
	if err := json.Unmarshal(intermediate.DeclarationProvider, &optRegVal); err == nil {
		serverCapabilities.DeclarationProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.DeclarationProvider, &boolVal); err == nil {
			serverCapabilities.DeclarationProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DefinitionOptions provides server capability options for goto definition requests.
type DefinitionOptions struct {
	WorkDoneProgressOptions
}

// unmarshals the DefinitionProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityDefinitionProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.DefinitionProvider == nil {
		return nil
	}

	var optVal DefinitionOptions
	if err := json.Unmarshal(intermediate.DefinitionProvider, &optVal); err == nil {
		serverCapabilities.DefinitionProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.DefinitionProvider, &boolVal); err == nil {
			serverCapabilities.DefinitionProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// TypeDefinitionOptions provides server capability options for goto type definition requests.
type TypeDefinitionOptions struct {
	WorkDoneProgressOptions
}

// TypeDefinitionRegistrationOptions provides server capability registration
// options for goto type definition requests.
type TypeDefinitionRegistrationOptions struct {
	TypeDefinitionOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the TypeDefinitionProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// TypeDefinitionRegistrationOptions is a superset of TypeDefinitionOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityTypeDefinitionProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.TypeDefinitionProvider == nil {
		return nil
	}

	var optRegVal TypeDefinitionRegistrationOptions
	if err := json.Unmarshal(intermediate.TypeDefinitionProvider, &optRegVal); err == nil {
		serverCapabilities.TypeDefinitionProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.TypeDefinitionProvider, &boolVal); err == nil {
			serverCapabilities.TypeDefinitionProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// ImplementationOptions provides server capability options for goto implementation requests.
type ImplementationOptions struct {
	WorkDoneProgressOptions
}

// ImplementationRegistrationOptions provides server capability registration
// options for goto implementation requests.
type ImplementationRegistrationOptions struct {
	ImplementationOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the ImplementationProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// ImplementationRegistrationOptions is a superset of ImplementationOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityImplementationProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.ImplementationProvider == nil {
		return nil
	}

	var optRegVal ImplementationRegistrationOptions
	if err := json.Unmarshal(intermediate.ImplementationProvider, &optRegVal); err == nil {
		serverCapabilities.ImplementationProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.ImplementationProvider, &boolVal); err == nil {
			serverCapabilities.ImplementationProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// ReferenceOptions provides server capability options for find references requests.
type ReferenceOptions struct {
	WorkDoneProgressOptions
}

// unmarshals the ReferencesProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityReferencesProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.ReferencesProvider == nil {
		return nil
	}

	var optVal ReferenceOptions
	if err := json.Unmarshal(intermediate.ReferencesProvider, &optVal); err == nil {
		serverCapabilities.ReferencesProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.ReferencesProvider, &boolVal); err == nil {
			serverCapabilities.ReferencesProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DocumentHighlightOptions provides server capability options for document highlight requests.
type DocumentHighlightOptions struct {
	WorkDoneProgressOptions
}

// unmarshals the DocumentHighlightProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityDocumentHighlightProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.DocumentHighlightProvider == nil {
		return nil
	}

	var optVal DocumentHighlightOptions
	if err := json.Unmarshal(intermediate.DocumentHighlightProvider, &optVal); err == nil {
		serverCapabilities.DocumentHighlightProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.DocumentHighlightProvider, &boolVal); err == nil {
			serverCapabilities.DocumentHighlightProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DocumentSymbolOptions provides server capability options for document symbol requests.
type DocumentSymbolOptions struct {
	WorkDoneProgressOptions

	// A human-readable string that is shown when multiple outline trees
	// are shown for the same document.
	//
	// @since 3.16.0
	Label *string `json:"label,omitempty"`
}

// unmarshals the DocumentSymbolProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityDocumentSymbolProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.DocumentSymbolProvider == nil {
		return nil
	}

	var optVal DocumentSymbolOptions
	if err := json.Unmarshal(intermediate.DocumentSymbolProvider, &optVal); err == nil {
		serverCapabilities.DocumentSymbolProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.DocumentSymbolProvider, &boolVal); err == nil {
			serverCapabilities.DocumentSymbolProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// CodeActionOptions provides server capability options for code action requests.
type CodeActionOptions struct {
	WorkDoneProgressOptions

	// CodeActionKinds that this server may return.
	//
	// The list of kinds may be generic, such as `CodeActionKind.Refactor`,
	// or the server may list out every specific kind they provide.
	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitempty"`

	// The server provides support to resolve additional
	// information for a code action.
	//
	// @since 3.16.0
	ResolveProvider *bool `json:"resolveProvider,omitempty"`
}

// The kind of a code action.
//
// Kinds are a hierarchical list of identifiers separated by `.`,
// e.g. `"refactor.extract.function"`.
//
// The set of kinds is open and client needs to announce the kinds it
// supports to the server during initialization.
type CodeActionKind = string

const (
	// CodeActionKindEmpty for the empty kind.
	CodeActionKindEmpty CodeActionKind = ""

	// CodeActionKindQuickFix is the base kind
	// for quickfix actions: 'quickfix'.
	CodeActionKindQuickFix CodeActionKind = "quickfix"

	// CodeActionKindRefactor is the base kind
	// for refactoring actions: 'refactor'.
	CodeActionKindRefactor CodeActionKind = "refactor"

	// CodeActionKindRefactorExtract is the kind for
	// refactoring extraction actions: 'refactor.extract'.
	//
	// Example extract actions:
	//
	// - Extract method
	// - Extract function
	// - Extract variable
	// - Extract interface from class
	// - ...
	CodeActionKindRefactorExtract CodeActionKind = "refactor.extract"

	// CodeActionKindRefactorInline is the base kind for
	// refactoring inline actions: 'refactor.inline'.
	//
	// Example inline actions:
	//
	// - Inline function
	// - Inline variable
	// - Inline constant
	// - ...
	CodeActionKindRefactorInline CodeActionKind = "refactor.inline"

	// CodeActionKindRefactorRewrite is the base kind for
	// refactoring rewrite actions: 'refactor.rewrite'.
	//
	// Example rewrite actions:
	//
	// - Convert JavaScript function to class
	// - Add or remove parameter
	// - Encapsulate field
	// - Make method static
	// - Move method to base class
	// - ...
	CodeActionKindRefactorRewrite CodeActionKind = "refactor.rewrite"

	// CodeActionKindSource is the base kind for
	// source actions: `source`.
	//
	// Source code actions apply to the entire file.
	CodeActionKindSource CodeActionKind = "source"

	// CodeActionKindSourceOrganizeImports is the base kind for
	// an organize imports source action: `source.organizeImports`.
	CodeActionKindSourceOrganizeImports CodeActionKind = "source.organizeImports"

	// CodeActionKindSourceFixAll is the base kind for
	// a 'fix all' source action: `source.fixAll`.
	//
	// 'Fix all' actions automatically fix errors that hae a clear fix that
	// do not require user input. They should not suppress errors or perform
	// unsafe fixes such as generating new types or classes.
	//
	// @since 3.17.0
	CodeActionKindSourceFixAll CodeActionKind = "source.fixAll"
)

// unmarshals the CodeActionProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityCodeActionProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.CodeActionProvider == nil {
		return nil
	}

	var optVal CodeActionOptions
	if err := json.Unmarshal(intermediate.CodeActionProvider, &optVal); err == nil {
		serverCapabilities.CodeActionProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.CodeActionProvider, &boolVal); err == nil {
			serverCapabilities.CodeActionProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DocumentColorOptions provides server capability options for document color requests.
type DocumentColorOptions struct {
	WorkDoneProgressOptions
}

// DocumentRegistrationOptions provides server capability registration
// options for document color requests.
type DocumentColorRegistrationOptions struct {
	DocumentColorOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the ColorProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// DocumentColorRegistrationOptions is a superset of DocumentColorOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityColorProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.ColorProvider == nil {
		return nil
	}

	var optRegVal DocumentColorRegistrationOptions
	if err := json.Unmarshal(intermediate.ColorProvider, &optRegVal); err == nil {
		serverCapabilities.ColorProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.ColorProvider, &boolVal); err == nil {
			serverCapabilities.ColorProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DocumentFormattingOptions provides server capability options for document formatting requests.
type DocumentFormattingOptions struct {
	WorkDoneProgressOptions
}

// unmarshals the DocumentFormattingProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityDocumentFormattingProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.DocumentFormattingProvider == nil {
		return nil
	}

	var optVal DocumentFormattingOptions
	if err := json.Unmarshal(intermediate.DocumentFormattingProvider, &optVal); err == nil {
		serverCapabilities.DocumentFormattingProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.DocumentFormattingProvider, &boolVal); err == nil {
			serverCapabilities.DocumentFormattingProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DocumentRangeFormattingOptions provides server capability options for
// document range formatting requests.
type DocumentRangeFormattingOptions struct {
	WorkDoneProgressOptions
}

// unmarshals the DocumentRangeFormattingProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityDocumentRangeFormattingProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.DocumentRangeFormattingProvider == nil {
		return nil
	}

	var optVal DocumentRangeFormattingOptions
	if err := json.Unmarshal(intermediate.DocumentRangeFormattingProvider, &optVal); err == nil {
		serverCapabilities.DocumentRangeFormattingProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.DocumentRangeFormattingProvider, &boolVal); err == nil {
			serverCapabilities.DocumentRangeFormattingProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// RenameOptions provides server capability options for rename requests.
type RenameOptions struct {
	WorkDoneProgressOptions

	// Renames should be checked and tested before being executed.
	PrepareProvider *bool `json:"prepareProvider,omitempty"`
}

// unmarshals the RenameProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityRenameProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.RenameProvider == nil {
		return nil
	}

	var optVal RenameOptions
	if err := json.Unmarshal(intermediate.RenameProvider, &optVal); err == nil {
		serverCapabilities.RenameProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.RenameProvider, &boolVal); err == nil {
			serverCapabilities.RenameProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// FoldingRangeOptions provides server capability options for folding range requests.
type FoldingRangeOptions struct {
	WorkDoneProgressOptions
}

// FoldingRangeRegistrationOptions provides server capability registration
// options for folding range requests.
type FoldingRangeRegistrationOptions struct {
	FoldingRangeOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the FoldingRangeProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// FoldingRangeRegistrationOptions is a superset of FoldingRangeOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityFoldingRangeProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.FoldingRangeProvider == nil {
		return nil
	}

	var optRegVal FoldingRangeRegistrationOptions
	if err := json.Unmarshal(intermediate.FoldingRangeProvider, &optRegVal); err == nil {
		serverCapabilities.FoldingRangeProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.FoldingRangeProvider, &boolVal); err == nil {
			serverCapabilities.FoldingRangeProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// SelectionRangeOptions provides server capability options for select range requests.
type SelectionRangeOptions struct {
	WorkDoneProgressOptions
}

// SelectionRangeRegistrationOptions provides server capability registration
// options for select range requests.
type SelectionRangeRegistrationOptions struct {
	SelectionRangeOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the SelectionRangeProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// SelectionRangeRegistrationOptions is a superset of SelectionRangeOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilitySelectionRangeProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.SelectionRangeProvider == nil {
		return nil
	}

	var optRegVal SelectionRangeRegistrationOptions
	if err := json.Unmarshal(intermediate.SelectionRangeProvider, &optRegVal); err == nil {
		serverCapabilities.SelectionRangeProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.SelectionRangeProvider, &boolVal); err == nil {
			serverCapabilities.SelectionRangeProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// LinkedEditingRangeOptions provides server capability options for linked editing range requests.
type LinkedEditingRangeOptions struct {
	WorkDoneProgressOptions
}

// LinkedEditingRangeRegistrationOptions provides server capability registration
// options for linked editing range requests.
type LinkedEditingRangeRegistrationOptions struct {
	LinkedEditingRangeOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the LinkedEditingRangeProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// LinkedEditingRangeRegistrationOptions is a superset of LinkedEditingRangeOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityLinkedEditingRangeProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.LinkedEditingRangeProvider == nil {
		return nil
	}

	var optRegVal LinkedEditingRangeRegistrationOptions
	if err := json.Unmarshal(intermediate.LinkedEditingRangeProvider, &optRegVal); err == nil {
		serverCapabilities.LinkedEditingRangeProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.LinkedEditingRangeProvider, &boolVal); err == nil {
			serverCapabilities.LinkedEditingRangeProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// CallHierarchyOptions provides server capability options for call hierarchy requests.
type CallHierarchyOptions struct {
	WorkDoneProgressOptions
}

// CallHierarchyRegistrationOptions provides server capability registration
// options for call hierarchy requests.
type CallHierarchyRegistrationOptions struct {
	CallHierarchyOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the CallHierarchyProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// CallHierarchyRegistrationOptions is a superset of CallHierarchyOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityCallHierarchyProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.CallHierarchyProvider == nil {
		return nil
	}

	var optRegVal CallHierarchyRegistrationOptions
	if err := json.Unmarshal(intermediate.CallHierarchyProvider, &optRegVal); err == nil {
		serverCapabilities.CallHierarchyProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.CallHierarchyProvider, &boolVal); err == nil {
			serverCapabilities.CallHierarchyProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// SemanticTokensOptions provides server capability options for semantic tokens requests.
type SemanticTokensOptions struct {
	WorkDoneProgressOptions

	// The legend used by the server.
	Legend SemanticTokensLegend `json:"legend"`

	// Server supports providing semantic tokens for a specific range of a document.
	//
	// bool | struct{} | nil
	Range any `json:"range,omitempty"`

	// Server supports providing semantic token for a full document.
	//
	// bool | SemanticDelta | nil
	Full any `json:"full,omitempty"`
}

type semanticTokenRegistrationOptionsIntermediate struct {
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
	WorkDoneProgressOptions
	Legend SemanticTokensLegend `json:"legend"`
	Range  json.RawMessage      `json:"range,omitempty"` // bool | struct{} | nil
	Full   json.RawMessage      `json:"full,omitempty"`  // bool | SemanticDelta | nil
}

// SemanticTokensRegistrationOptions provides server capability registration
// options for semantic tokens requests.
type SemanticTokensRegistrationOptions struct {
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
	SemanticTokensOptions
}

// Fulfils the json.Unmarshaler interface.
func (s *SemanticTokensRegistrationOptions) UnmarshalJSON(data []byte) error {
	// If a struct has a custom UnmarshalJSON and it is embedded, it will be used
	// to unmarshal the JSON data into the entire struct instead of just the embedded struct,
	// leaving other fields empty.
	// This is why we need to unmarshal from the top-level of the semantic token options type
	// and can't have a clean separate of the SematicTokenRegistrationOptions
	// and the SemanticTokensOptions types when it comes to deserialising JSON.
	var value semanticTokenRegistrationOptionsIntermediate

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	s.TextDocumentRegistrationOptions = value.TextDocumentRegistrationOptions
	s.StaticRegistrationOptions = value.StaticRegistrationOptions
	s.SemanticTokensOptions.WorkDoneProgressOptions = value.WorkDoneProgressOptions
	s.Legend = value.Legend

	err := s.unmarshalSemanticTokenOptionsRange(value)
	if err != nil {
		return err
	}

	err = s.unmarshalSemanticTokenOptionsFull(value)
	return err
}

func (s *SemanticTokensRegistrationOptions) unmarshalSemanticTokenOptionsFull(
	value semanticTokenRegistrationOptionsIntermediate,
) error {
	if value.Full != nil {
		var boolVal bool
		if err := json.Unmarshal(value.Full, &boolVal); err == nil {
			s.Full = boolVal
		} else {
			var deltaVal SemanticDelta
			if err := json.Unmarshal(value.Full, &deltaVal); err == nil {
				s.Full = deltaVal
			} else {
				return err
			}
		}
	}

	return nil
}

func (s *SemanticTokensRegistrationOptions) unmarshalSemanticTokenOptionsRange(
	value semanticTokenRegistrationOptionsIntermediate,
) error {
	if value.Range != nil {
		var boolVal bool
		if err := json.Unmarshal(value.Range, &boolVal); err == nil {
			s.Range = boolVal
		} else {
			var structVal struct{}
			if err := json.Unmarshal(value.Range, &structVal); err == nil {
				s.Range = structVal
			} else {
				return err
			}
		}
	}

	return nil
}

func unmarshalServerCapabilitySemanticTokensProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.SemanticTokensProvider == nil {
		return nil
	}

	var optRegVal SemanticTokensRegistrationOptions
	err := json.Unmarshal(intermediate.SemanticTokensProvider, &optRegVal)
	if err != nil {
		return err
	}

	serverCapabilities.SemanticTokensProvider = optRegVal
	return nil
}

// SemanticDelta represents the server's support for deltas in semantic tokens.
type SemanticDelta struct {
	// The server supports deltas for full documents.
	Delta *bool `json:"delta,omitempty"`
}

// SemanticTokensLenged represent's the server's way of letting the client
// know which numbers it is using for which types and modifiers.
type SemanticTokensLegend struct {
	// The token types a server uses.
	TokenTypes []string `json:"tokenTypes"`

	// The token modifiers a server uses.
	TokenModifiers []string `json:"tokenModifiers"`
}

// MonikerOptions provides server capability options for moniker requests.
type MonikerOptions struct {
	WorkDoneProgressOptions
}

// MonikerRegistrationOptions provides server capability registration
// options for moniker requests.
type MonikerRegistrationOptions struct {
	MonikerOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the MonikerProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// MonikerRegistrationOptions is a superset of MonikerOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityMonikerProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.MonikerProvider == nil {
		return nil
	}

	var optRegVal MonikerRegistrationOptions
	if err := json.Unmarshal(intermediate.MonikerProvider, &optRegVal); err == nil {
		serverCapabilities.MonikerProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.MonikerProvider, &boolVal); err == nil {
			serverCapabilities.MonikerProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// TypeHierarchyOptions provides server capability options for type hierarchy requests.
type TypeHierarchyOptions struct {
	WorkDoneProgressOptions
}

// TypeHierarchyRegistrationOptions provides server capability registration
// options for type hierarchy requests.
type TypeHierarchyRegistrationOptions struct {
	TypeHierarchyOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the TypeHierarchyProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// TypeHierarchyRegistrationOptions is a superset of TypeHierarchyOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityTypeHierarchyProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.TypeHierarchyProvider == nil {
		return nil
	}

	var optRegVal TypeHierarchyRegistrationOptions
	if err := json.Unmarshal(intermediate.TypeHierarchyProvider, &optRegVal); err == nil {
		serverCapabilities.TypeHierarchyProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.TypeHierarchyProvider, &boolVal); err == nil {
			serverCapabilities.TypeHierarchyProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// InlineValueOptions provides server capability options for inline value requests.
//
// @since 3.17.0
type InlineValueOptions struct {
	WorkDoneProgressOptions
}

// InlineValueRegistrationOptions provides server capability registration
// options for inline value requests.
//
// @since 3.17.0
type InlineValueRegistrationOptions struct {
	InlineValueOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the InlineValueProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// InlineValueRegistrationOptions is a superset of InlineValueOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityInlineValueProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.InlineValueProvider == nil {
		return nil
	}

	var optRegVal InlineValueRegistrationOptions
	if err := json.Unmarshal(intermediate.InlineValueProvider, &optRegVal); err == nil {
		serverCapabilities.InlineValueProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.InlineValueProvider, &boolVal); err == nil {
			serverCapabilities.InlineValueProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// InlayHintOptions provides server capability options for inline value requests.
//
// @since 3.17.0
type InlayHintOptions struct {
	WorkDoneProgressOptions

	// The server provides support to resolve additional
	// information for an inlay hint item.
	ResolveProvider *bool `json:"resolveProvider,omitempty"`
}

// InlayHintRegistrationOptions provides server capability registration
// options for inline value requests.
//
// @since 3.17.0
type InlayHintRegistrationOptions struct {
	InlayHintOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the InlayHintProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// InlayHintRegistrationOptions is a superset of InlayHintOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityInlayHintProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.InlayHintProvider == nil {
		return nil
	}

	var optRegVal InlayHintRegistrationOptions
	if err := json.Unmarshal(intermediate.InlayHintProvider, &optRegVal); err == nil {
		serverCapabilities.InlayHintProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.InlayHintProvider, &boolVal); err == nil {
			serverCapabilities.InlayHintProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DiagnosticOptions provides server capability options for pull diagnostics behaviour.
//
// @since 3.17.0
type DiagnosticOptions struct {
	WorkDoneProgressOptions

	// An optional identifier under which the diagnostics are
	// managed by the client.
	Identifier *string `json:"identifier,omitempty"`

	// Whether the language has inter file dependencies meaning that
	// editing code in one file can result in a different diagnostic
	// set in another file. Inter file dependencies are common for
	// most programming languages and typically uncommon for linters.
	InterFileDependencies bool `json:"interFileDependencies"`

	// The server provides support for workspace diagnostics as well.
	WorkspaceDiagnostics bool `json:"workspaceDiagnostics"`
}

// DiagnosticRegistrationOptions provides server capability registration
// options for pull diagnostics behaviour.
//
// @since 3.17.0
type DiagnosticRegistrationOptions struct {
	DiagnosticOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the DiagnosticProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// DiagnosticRegistrationOptions is a superset of DiagnosticOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityDiagnosticProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.DiagnosticProvider == nil {
		return nil
	}

	var optRegVal DiagnosticRegistrationOptions
	if err := json.Unmarshal(intermediate.DiagnosticProvider, &optRegVal); err == nil {
		serverCapabilities.DiagnosticProvider = optRegVal
	} else {
		return err
	}

	return nil
}

// WorkspaceSymbolOptions provides server capability options for workspace symbol requests.
type WorkspaceSymbolOptions struct {
	WorkDoneProgressOptions

	// The server provides support to resolve additional
	// information for a workspace symbol.
	//
	// @since 3.17.0
	ResolveProvider *bool `json:"resolveProvider,omitempty"`
}

// unmarshals the WorkspaceSymbolProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityWorkspaceSymbolProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.WorkspaceSymbolProvider == nil {
		return nil
	}

	var optVal WorkspaceSymbolOptions
	if err := json.Unmarshal(intermediate.WorkspaceSymbolProvider, &optVal); err == nil {
		serverCapabilities.WorkspaceSymbolProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.WorkspaceSymbolProvider, &boolVal); err == nil {
			serverCapabilities.WorkspaceSymbolProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}
