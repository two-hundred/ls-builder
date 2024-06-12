package lsp

import (
	"strings"
	"unicode/utf8"
)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#uri

// DocumentURI is a string representing a URI to a document.
type DocumentURI = string

// URI is a string that is used for tagging normal
// non-document URIs.
type URI = string

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#regExp

// RegularExpressionsClientCapabilities represents the capabilities of the client
// for regular expressions.
type RegularExpressionsClientCapabilities struct {
	// The engine's name.
	Engine string `json:"engine"`
	// The engine's version.
	Version *string `json:"version,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocuments

var EOL = []string{"\n", "\r\n", "\r"}

// A Position in a text document expressed as zero-based line and character offset.
type Position struct {
	// Line position in a document (zero-based).
	Line UInteger `json:"line"`

	// Character offset on a line in a document (zero-based).
	// The meaning of this offset is determined by the negotiated
	// `PositionEncodingKind`.
	//
	// If the character value is greater than the line length it
	// defaults back to the line length.
	Character UInteger `json:"character"`
}

func (p Position) IndexIn(text string, posEncodingKind PositionEncodingKind) int {
	// This code is modified from the gopls implementation found:
	// https://cs.opensource.google/go/x/tools/+/refs/tags/v0.1.5:internal/span/utf16.go;l=70

	// In accordance with the LSP Spec:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-3-17#textDocuments
	// p.Character represents utf-8, utf-16 or utf-32 code units depending on the position encoding kind.
	// These are not bytes and to indexes in Go strings are in bytes (as Go strings are made up of byte arrays),
	// so we need to convert the character offset to a byte offset.
	// Strings in Go are UTF-8 encoded, so we can use the `utf8` package to convert the character offset
	// to a byte offset and then based on the agreed position encoding kind, count the characters
	// up to p.Character to get the byte offset.

	// Find the byte offset for the line.
	index := p.getLineIndex(text)

	// The index represents the byte offsets from the beginning of the line counting
	// p.Character code units (utf-8, utf-16 or utf-32) from the index byte offset.
	byteOffset := index
	remains := text[index:]
	char := int(p.Character)

	for count := 1; count <= char; count += 1 {

		if len(remains) <= 0 {
			// char has gone past the end of the content,
			// this indicates an invalid position.
			return 0
		}

		r, w := utf8.DecodeRuneInString(remains)
		if r == utf8.RuneError {
			// Invalid UTF-8 sequence.
			// This indicates an invalid position.
			return 0
		}

		if r == '\n' {
			// As per the LSP Spec:
			//
			// If the character value is greater than the line length it
			// defaults back to the line length.
			break
		}

		remains = remains[w:]
		if posEncodingKind == PositionEncodingKindUTF16 {
			if r >= utf16_2CodePoints {
				// A rune that holds a code point in the range 0x10000 to 0x10FFFF
				// that is represented by two code units in UTF-16.
				count += 1
				// If we have finished in a two code unit rune, there is no need
				// to go further.
				if count > char {
					break
				}
			}
		} else if posEncodingKind == PositionEncodingKindUTF8 {
			// UTF-8 code units are bytes, so we can just count the bytes
			// as characters until we reach the desired character.
			count += w - 1
		}

		// In the case of UTF-32, each rune is a code point as per an
		// encoding-agnostic representation of character offsets.
		// This means that we can just count each rune as a character.
		byteOffset += w
	}

	return byteOffset
}

func (p Position) EndOfLineIn(content string, posEncodingKind PositionEncodingKind) Position {
	index := p.IndexIn(content, posEncodingKind)
	remains := content[index:]
	if eol := strings.Index(remains, "\n"); eol != -1 {
		return Position{
			Line:      p.Line,
			Character: p.Character + UInteger(eol),
		}
	} else {
		return p
	}
}

func (p Position) getLineIndex(text string) int {
	hasNewLine := true
	index := 0
	row := UInteger(0)
	for hasNewLine && row < p.Line {
		rest := string(text[index:])
		if next := strings.Index(rest, "\n"); next != -1 {
			index += next + 1
		} else {
			hasNewLine = false
		}
		row += 1
	}
	return index
}

const (
	// The threshold for a rune that holds a code point in the range
	// 0x10000 to 0x10FFFF that is represented by two code units in UTF-16.
	// These are code points outside of the Basic Multilingual Plane (BMP).
	utf16_2CodePoints = 0x10000
)

// PositionEncodingKind is a type that indicates
// how positions are encoded,
// specifically what column offsets mean.
//
// @since 3.17.0
type PositionEncodingKind string

const (
	// PositionEncodingKindUTF8 is used when the
	// character offsets count UTF-8 code units (e.g. bytes).
	//
	// @since 3.17.0
	PositionEncodingKindUTF8 PositionEncodingKind = "utf-8"

	// PositionEncodingKindUTF16 is used when the
	// character offsets count UTF-16 code units.
	//
	// This is the default and must always be supported by servers.

	// @since 3.17.0
	PositionEncodingKindUTF16 PositionEncodingKind = "utf-16"

	// PositionEncodingKindUTF32 is used when the
	// character offsets count UTF-32 code units.
	//
	// Note: these are the same as unicode code points,
	// so this `PositionEncodingKind` may also be used for an
	// encoding-agnostic representation of character offsets.
	//
	// @since 3.17.0
	PositionEncodingKindUTF32 PositionEncodingKind = "utf-32"
)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#range

// A Range in the text document expressed as (zero-based) start and end positions.
// A range is comparable to a selection in an editor where the end position is
// exclusive.
type Range struct {
	// The range's start position.
	Start Position `json:"start"`

	// The range's end position.
	End Position `json:"end"`
}

func (r Range) IndexesIn(content string, posEncodingKind PositionEncodingKind) (int, int) {
	return r.Start.IndexIn(content, posEncodingKind), r.End.IndexIn(content, posEncodingKind)
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentItem

// TextDocumentItem is an item to transfer a text document from the client to the server.
type TextDocumentItem struct {
	// The text document's URI.
	URI DocumentURI `json:"uri"`

	// The text document's language identifier.
	LanguageID string `json:"languageId"`

	// The version number of this document (it will strictly increase after each
	// change, including undo/redo).
	Version Integer `json:"version"`

	// The content of the opened text document.
	Text string `json:"text"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentIdentifier

// TextDocumentIdentifier identifies a text document.
type TextDocumentIdentifier struct {
	// The text document's URI.
	URI DocumentURI `json:"uri"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#versionedTextDocumentIdentifier

// VersionedTextDocumentIdentifier is an identifier to denote a specific version of a text document.
type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier

	// The version number of this document.
	// The version number of this document. If an optional versioned text document
	// identifier is sent from the server to the client and the file is not
	// open in the editor (the server has not received an open notification
	// before) the server can send `null` to indicate that the version is
	// known and the content on disk is the master (as specified with document
	// content ownership).
	//
	// The version number of a document will increase after each change,
	// including undo/redo. The number doesn't need to be consecutive.
	Version Integer `json:"version"`
}

// VersionedTextDocumentIdentifier is an identifier which optionally denotes a specific version of a text document.
// This information usually flows from the server to the client.
type OptionalVersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier

	// The version number of this document.
	// The version number of this document. If an optional versioned text document
	// identifier is sent from the server to the client and the file is not
	// open in the editor (the server has not received an open notification
	// before) the server can send `null` to indicate that the version is
	// known and the content on disk is the master (as specified with document
	// content ownership).
	//
	// The version number of a document will increase after each change,
	// including undo/redo. The number doesn't need to be consecutive.
	Version *Integer `json:"version,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentPositionParams

// TextDocumentPositionParams is the parameters of a request that
// requires a text document and a position.
// A parameter literal used in requests to pass a text document and a position inside that document.
// It is up to the client to decide how a selection is converted into a position
// when issuing a request for a text document.
// The client can for example honor or ignore the selection direction to make LSP request
// consistent with features implemented internally.
type TextDocumentPositionParams struct {
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The position inside the text document.
	Position Position `json:"position"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#documentFilter

// A DocumentFilter denotes a document through properties like
// language, scheme or pattern.
// An example is a filter that applies to TypeScript files on disk.
// Another example is a filter that applies to JSON files with name package.json:
// { language: 'typescript', scheme: 'file' }
// { language: 'json', pattern: '**/package.json' }
type DocumentFilter struct {
	// A language id, like `typescript`.
	Language *string `json:"language,omitempty"`

	// A Uri [scheme](#Uri.scheme), like `file` or `untitled`.
	Scheme *string `json:"scheme,omitempty"`

	// A glob pattern, like `*.{ts,js}`.
	// - `*` to match one or more characters in a path segment
	// - `?` to match on one character in a path segment
	// - `**` to match any number of path segments, including none
	// - `{}` to group sub patterns into an OR expression. (e.g. `**​/*.{ts,js}`
	//   matches all TypeScript and JavaScript files)
	// - `[]` to declare a range of characters to match in a path segment
	//   (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
	// - `[!...]` to negate a range of characters to match in a path segment
	//   (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but
	//   not `example.0`)
	Pattern *string `json:"pattern,omitempty"`
}

// A DocumentSelector is the combination of one or more document filters.
type DocumentSelector = []DocumentFilter

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentEdit

// TextEdit is a textual edit applicable to a text document.
type TextEdit struct {
	// The range of the text document to be manipulated. To insert
	// text into a document create a range where start === end.
	Range Range `json:"range"`

	// The string to be inserted. For delete operations use an
	// empty string.
	NewText string `json:"newText"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentContentChangeEvent

// ChangeAnnotation provides additinoa information that
// describes document changes.
//
// @since 3.16.0
type ChangeAnnotation struct {
	// A human-readable string describing the change.
	// This is rendered prominently in the user interface.
	Label string `json:"label"`

	// A flag which indicates that user confirmation is needed
	// before applying the change.
	NeedsConfirmation *bool `json:"needsConfirmation,omitempty"`

	// A human-readable string which is rendered less prominent
	// in the user interface.
	Description *string `json:"description,omitempty"`
}

// ChangeAnnotationIdentifier is a string that identifies a change annotation
// managed by a workspace edit.
//
// @since 3.16.0
type ChangeAnnotationIdentifier = string

// A special text edit with additional change annotation.
//
// @since 3.16.0
type AnnotatedTextEdit struct {
	TextEdit

	// The actual annotation identifier.
	AnnotationID ChangeAnnotationIdentifier `json:"annotationId"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentEdit

type TextDocumentEdit struct {
	// The text document to change.
	TextDocument OptionalVersionedTextDocumentIdentifier `json:"textDocument"`

	// The edits to be applied.
	//
	// @since  3.16.0 - support for AnnotatedTextEdit. This
	// is guarded by the client capability `workspace.workspaceEdit.changeAnnotationSupport`.
	Edits []any `json:"edits"` // TextEdit | AnnotatedTextEdit (checked at runtime)
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#location

// Location represents a location inside a resource, such as a line
// inside a text file.
type Location struct {
	// The URI of the document.
	URI DocumentURI `json:"uri"`

	// The range in the document.
	Range Range `json:"range"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#locationLink

// LocationLink represents a link between a source and a target location.
type LocationLink struct {
	// Span of the origin of this link.
	// Used as the underlined span for mouse interaction.
	// Defaults to the word range at the mouse position.
	OriginSelectionRange *Range `json:"originSelectionRange,omitempty"`

	// The target resource identifier of this link.
	TargetURI DocumentURI `json:"targetUri"`

	// The full target range of this link. If the target for example is a symbol then target range is the
	// range enclosing this symbol not including leading/trailing whitespace but everything else
	// like comments. This information is typically used to highlight the range in the editor.
	TargetRange Range `json:"targetRange"`

	// The range that should be selected and revealed when this link is
	// being followed, e.g the name of a function.
	// Must be contained by the the `targetRange`. See also `DocumentSymbol#range`
	TargetSelectionRange Range `json:"targetSelectionRange"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#diagnostic

// A Diagnostic, such as a compiler error or warning.
// Diagnostic objects are only valid in the scope of a resource.
type Diagnostic struct {
	// The range at which the message applies.
	Range Range `json:"range"`

	// The diagnostic's severity. Can be omitted. If omitted it is up to the
	// client to interpret diagnostics as error, warning, info or hint.
	Severity *DiagnosticSeverity `json:"severity,omitempty"`

	// The diagnostic's code, which might appear in the user interface.
	Code *IntOrString `json:"code,omitempty"`

	// An optional property to describe the error code.
	// @since 3.16.0
	CodeDescription *CodeDescription `json:"codeDescription,omitempty"`

	// A human-readable string describing the source of this
	// diagnostic, e.g. 'typescript' or 'super lint'.
	Source *string `json:"source,omitempty"`

	// The diagnostic's message.
	Message string `json:"message"`

	// Additional metadata about the diagnostic.
	//
	// @since 3.15.0
	Tags []DiagnosticTag `json:"tags,omitempty"`

	// An array of related diagnostic information, e.g. when symbol-names within
	// a scope collide all definitions can be marked via this property.
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation,omitempty"`

	// A data entry field that is preserved between a `textDocument/publishDiagnostics`
	// notification and `textDocument/codeAction` request.
	//
	// @since 3.16.0
	Data any `json:"data,omitempty"`
}

type DiagnosticSeverity = UInteger

const (
	// Reports an error.
	DiagnosticSeverityError DiagnosticSeverity = 1
	// Reports a warning.
	DiagnosticSeverityWarning DiagnosticSeverity = 2
	// Reports an information.
	DiagnosticSeverityInformation DiagnosticSeverity = 3
	// Reports a hint.
	DiagnosticSeverityHint DiagnosticSeverity = 4
)

type DiagnosticTag = UInteger

const (
	// Unused or unnecessary code.
	// Clients are allowed to render diagnostics with this tag faded out instead of having
	// an error squiggle.
	DiagnosticTagUnnecessary DiagnosticTag = 1

	// Deprecated or obsolete code.
	// Clients are allowed to rendered diagnostics with this tag strike through.
	DiagnosticTagDeprecated DiagnosticTag = 2
)

// DiagnosticRelatedInformation represents a related message and
// source code location for a diagnostic.
// This should be used to point to code locations that cause or
// related to a diagnostics, e.g when duplicating a symbol in a scope.
type DiagnosticRelatedInformation struct {
	// The location of this related diagnostic information.
	Location Location `json:"location"`

	// The message of this related diagnostic information.
	Message string `json:"message"`
}

// Structure to capture a description for an error code.
// @since 3.16.0
type CodeDescription struct {
	// A URI to open with more information about the diagnostic error.
	Href URI `json:"href"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#command

// Command represents a reference to a command.
// Provides a title which will be used to represente a command in the UI.
// Commands are identified by a string identifier.
// The recommended way to handle commands
// is to implement their execution on the server side if the client and server provides
// the corresponding capabilities.
// Alternatively, the toll extension code could handle the command.
// The protocol currently doesn't specify a set of well-known commands.
type Command struct {
	// Title of the command, like `save`.
	Title string `json:"title"`

	// The identifier of the actual command handler.
	Command string `json:"command"`

	// Arguments that the command handler should be
	// invoked with.
	Arguments []LSPAny `json:"arguments,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#markupContent

// MarkupKind describes the content type that client supports
// in various result literals like `Hover`, `ParameterInfo` or `CompletionItem`.
// Please note that `MarkupKinds` must not start with a `$`. These kinds are
// reserved for internal usage.
type MarkupKind = string

const (
	// Plain text is supported as a content format
	MarkupKindPlainText MarkupKind = "plaintext"
	// Markdown is supported as a content format
	MarkupKindMarkdown MarkupKind = "markdown"
)

// A MarkupContent literal represents a string value which content is
// interpreted base on its kind flag. Currently the protocol supports
// `plaintext` and `markdown` as markup kinds.
//
// If the kind is `markdown` then the value can contain fenced code blocks like
// in GitHub issues.
//
// Here is an example how such a string can be constructed using
// JavaScript / TypeScript:
// ```typescript
//
//	let markdown: MarkdownContent = {
//		kind: MarkupKind.Markdown,
//		value: [
//			'# Header',
//			'Some text',
//			'```typescript',
//			'someCode();',
//			'```'
//		].join('\n')
//	};
//
// ```
//
// *Please Note* that clients might sanitize the return markdown. A client could
// decide to remove HTML from the markdown to avoid script execution.
type MarkupContent struct {
	// The type of the Markup
	Kind MarkupKind `json:"kind"`

	// The content itself
	Value string `json:"value"`
}

// MarkdownClientCapabilities describes the capabilities specific
// to the used markdown parser.
//
// @since 3.16.0
type MarkdownClientCapabilities struct {
	// The name of the parser.
	Parser string `json:"parser"`

	// The version of the parser.
	Version *string `json:"version,omitempty"`

	// A list of HTML tags that the client allows / supports
	// in Markdown.
	//
	// @since 3.17.0
	AllowedTags []string `json:"allowedTags,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#resourceChanges

// CreateFileOptions provides options to create a file.
type CreateFileOptions struct {
	// Overwrite existing file. Overwrite wins over `ignoreIfExists`
	Overwrite *bool `json:"overwrite,omitempty"`

	// Ignore if exists.
	IgnoreIfExists *bool `json:"ignoreIfExists,omitempty"`
}

// CreateFile defines an operation to create a file.
type CreateFile struct {
	// A create operatoin
	Kind string `json:"kind"` // == "create"

	// The resource to create.
	URI DocumentURI `json:"uri"`

	// Additional options
	Options *CreateFileOptions `json:"options,omitempty"`

	// An optional annotation identifier describing the operation.
	//
	// @since 3.16.0
	AnnotationID *ChangeAnnotationIdentifier `json:"annotationId,omitempty"`
}

// RenameFileOptions provides options to rename a file.
type RenameFileOptions struct {
	// Overwrite target if existing. Overwrite wins over `ignoreIfExists`
	Overwrite *bool `json:"overwrite,omitempty"`

	// Ignores if target exists.
	IgnoreIfExists *bool `json:"ignoreIfExists,omitempty"`
}

// RenameFile defines an operation to rename a file.
type RenameFile struct {
	// A rename operation.
	Kind string `json:"kind"` // == "rename"

	// The old (existing) location.
	OldURI DocumentURI `json:"oldUri"`

	// The new location.
	NewURI DocumentURI `json:"newUri"`

	// Rename options.
	Options *RenameFileOptions `json:"options,omitempty"`

	// An optional annotation identifier describing the operation.
	//
	// @since 3.16.0
	AnnotationID *ChangeAnnotationIdentifier `json:"annotationId,omitempty"`
}

// DeleteFileOptions provides options to delete a file.
type DeleteFileOptions struct {
	// Delete the content recursively if a folder is denoted.
	Recursive *bool `json:"recursive,omitempty"`

	// Ignore the operation if the file doesn't exist.
	IgnoreIfNotExists *bool `json:"ignoreIfNotExists,omitempty"`
}

// DeleteFile defines an operation to delete a file.
type DeleteFile struct {
	// A delete operation.
	Kind string `json:"kind"` // == "delete"

	// The file to delete.
	URI DocumentURI `json:"uri"`

	// Delete options.
	Options *DeleteFileOptions `json:"options,omitempty"`

	// An optional annotation identifier describing the operation.
	//
	// @since 3.16.0
	AnnotationID *ChangeAnnotationIdentifier `json:"annotationId,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspaceEdit

// A WorkspaceEdit represents changes to many resources managed in the workspace.
// The edit should either provide changes or documentChanges. If the client can handle
// versioned document edits and if documentChanges are present, the latter are preferred over changes.
type WorkspaceEdit struct {
	// Holds changes to existing resources.
	Changes map[DocumentURI][]TextEdit `json:"changes,omitempty"`

	// Depending on the client capaiblity
	// `workspace.workspaceEdit.resourceOperations` document changes are either
	// an array of `TextDocumentEdit`s to express changes to n different text documents
	// where each text document edit addresses a specific version of a text document.
	// Or it can contain above `TextDocumentEdit`s mixed with create, rename and delete file / folder operations.
	//
	// Whether a client supports versioned document edits is expressed via
	// `workspace.workspaceEdit.documentChanges` client capability.
	//
	// If a client neither supports `documentChanges` nor `workspace.workspaceEdit.resourceOperations`
	// then only plain `TextEdit`s using the `changes` property are supported.
	DocumentChanges []any `json:"documentChanges,omitempty"` // TextDocumentEdit | CreateFile | RenameFile | DeleteFile (checked at runtime)

	// A map of change annotations that can be referenced in
	// `AnnotatedTextEdit`s or create, rename and delete file / folder operations.
	//
	// Whether clients honour this property depends on the client capability
	// `workspace.changeAnnotationSupport`.
	//
	// @since 3.16.0
	ChangeAnnotations map[ChangeAnnotationIdentifier]ChangeAnnotation `json:"changeAnnotations,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspaceEditClientCapabilities

// WorkspaceEditClientCapabilities defines the capabilities the client has for
// workspace edits.
type WorkspaceEditClientCapabilities struct {
	// The client supports versioned document changes in `WorkspaceEdit`s
	DocumentChanges *bool `json:"documentChanges,omitempty"`

	// The resource operations the client supports. Clients should at least
	// support 'create', 'rename' and 'delete' files and folders.
	//
	// @since 3.13.0
	ResourceOperations []ResourceOperationKind `json:"resourceOperations,omitempty"`

	// The failure handling strategy of a client if applying the workspace edit
	// fails.
	//
	// @since 3.13.0
	FailureHandling FailureHandlingKind `json:"failureHandling,omitempty"`

	// Whether the client normalises line endings to the client specific
	// setting.
	// If set to `true` the client will normalize line ending characters
	// in a workspace edit to the client specific new line character(s).
	//
	// @since 3.16.0
	NormalizesLineEndings *bool `json:"normalizesLineEndings,omitempty"`

	// Whether the client in general supports change annotations on text edits,
	// create file, rename file and delete file changes.
	//
	// @since 3.16.0
	ChangeAnnotationSupport *ChangeAnnotationSupport `json:"changeAnnotationSupport,omitempty"`
}

type ChangeAnnotationSupport struct {
	// Whether the client groups edits with equal labels into tree nodes,
	// for instance all edits labelled with "Changes in Strings" would be
	// a tree node.
	GroupsOnLabel *bool `json:"groupsOnLabel,omitempty"`
}

type ResourceOperationKind = string

const (
	// Supports creating new files and folders.
	ResourceOperationKindCreate ResourceOperationKind = "create"
	// Supports renaming existing files and folders.
	ResourceOperationKindRename ResourceOperationKind = "rename"
	// Supports deleting existing files and folders.
	ResourceOperationKindDelete ResourceOperationKind = "delete"
)

type FailureHandlingKind = string

const (
	// Applying the workspace change is simply aborted if one of the changes provided
	// fails. All operations executed before the failing operation stay executed.
	FailureHandlingKindAbort FailureHandlingKind = "abort"
	// All operations are executed transactional. That means they either all
	// succeed or no changes at all are applied to the workspace.
	FailureHandlingKindTransactional FailureHandlingKind = "transactional"
	// If the workspace edit contains only textual file changes they are executed transactional.
	// If resource changes (create, rename or delete file) are part of the change the failure
	// handling strategy is abort.
	FailureHandlingKindTextOnlyTransactional FailureHandlingKind = "textOnlyTransactional"
	// The client tries to undo the operations already executed. But there is no
	// guarantee that this is succeeding.
	FailureHandlingKindUndo FailureHandlingKind = "undo"
)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workDoneProgress

// WorkDoneProgressBegin is the payload sent with a `$/progress` notification
// that signals the begining of some background work.
type WorkDoneProgressBegin struct {
	Kind string `json:"kind"` // == "begin"

	// Mandatory title of the progress operation. Used to briefly inform about
	// the kind of operation being performed.
	//
	// Examples: "Indexing" or "Linking dependencies".
	Title string `json:"title"`

	// Controls if a cancel button should show to allow the user to cancel the
	// long running operation. Clients that don't support cancellation are allowed
	// to ignore the setting.
	Cancellable *bool `json:"cancellable,omitempty"`

	// Optional, more detailed associated progress message. Contains
	// complementary information to the `title`.
	//
	// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
	// If unset, the previous progress message (if any) is still valid.
	Message *string `json:"message,omitempty"`

	// Optional progress percentage to display (value 100 is considered 100%).
	// If not provided infinite progress is assumed and clients are allowed
	// to ignore the `percentage` value in subsequent in report notifications.
	//
	// The value should be steadily rising. Clients are free to ignore values
	// that are not following this rule. The value range is [0, 100].
	Percentage *UInteger `json:"percentage,omitempty"`
}

// WorkDoneProgressReport is the payload that is used when reporting progress.
type WorkDoneProgressReport struct {
	Kind string `json:"kind"` // == "report"

	// Controls enablement state of a cancel button. This property is only valid
	// if a cancel button got requested in the `WorkDoneProgressBegin` payload.
	//
	// Clients that don't support cancellation or don't support control the
	// button's enablement state are allowed to ignore the setting.
	Cancellable *bool `json:"cancellable,omitempty"`

	// Optional, more detailed associated progress message. Contains
	// complementary information to the `title`.
	//
	// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
	// If unset, the previous progress message (if any) is still valid.
	Message *string `json:"message,omitempty"`

	// Optional progress percentage to display (value 100 is considered 100%).
	// If not provided infinite progress is assumed and clients are allowed
	// to ignore the `percentage` value in subsequent in report notifications.
	//
	// The value should be steadily rising. Clients are free to ignore values
	// that are not following this rule. The value range is [0, 100].
	Percentage *UInteger `json:"percentage,omitempty"`
}

// WorkDoneProgressEnd is the payload sent signalling the end of of some work.
type WorkDoneProgressEnd struct {
	Kind string `json:"kind"` // == "end"

	// Optional, a final message indicating to for example indicate the outcome
	// of the operation.
	Message *string `json:"message,omitempty"`
}

// WorkDoneProgressParams provides parameters that should be provided for client-initiated
// progress.
type WorkDoneProgressParams struct {
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}

// WorkDoneProgressOptions is the type for server capabilities.
//
// An example usage would be:
// ```json
//
//	{
//		"referencesProvider": {
//			"workDoneProgress": true
//		}
//	}
//
// ```
type WorkDownProgressOptions struct {
	WorkDoneProgress *bool `json:"workDoneProgress,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#partialResults

// PatialResultParams allow the passing of a partial result token
// in a context that supports reporting progress for partial results.
type PartialResultParams struct {
	// An optional token that a server can use to report partial results (e.g.
	// streaming) to the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#traceValue

// A TraceValue represents the level of verbosity with which the server systematically reports
// its execution trace using `$/logTrace` notifications.
type TraceValue string

const (
	TraceValueOff     = TraceValue("off")
	TraceValueMessage = TraceValue("messages")
	TraceValueVerbose = TraceValue("verbose")
)
