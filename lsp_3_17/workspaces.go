package lsp

import (
	"encoding/json"

	"github.com/two-hundred/ls-builder/common"
)

// DidChangeConfigurationClientCapabilities describes the capabilities
// of a client for the `workspace/didChangeConfiguration` request.
type DidChangeConfigurationClientCapabilities struct {
	// Did change configuration notification supports dynamic registration.
	//
	// @since 3.6.0 to support the new pull model.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// DidChangeWatchedFilesClientCapabilities describes the capabilities
// of a client for the `workspace/didChangeWatchedFiles` request.
type DidChangeWatchedFilesClientCapabilities struct {
	// Did change watched files notification supports dynamic registration.
	// Please note that the current protocol doesn't support static
	// configuration for file changes from the server side.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// Whether the client has support for relative patterns
	// or not.
	//
	// @since 3.17.0
	RelativePatternSupport *bool `json:"relativePatternSupport,omitempty"`
}

// WorkspaceSymbolClientCapabilities describes the capabilities of a client
// for the `workspace/symbol` request.
type WorkspaceSymbolClientCapabilities struct {
	// Symbol request supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// Specific capabilities for the `SymbolKind` in the `workspace/symbol` request.
	SymbolKind *SymbolKindCapabilities `json:"symbolKind,omitempty"`

	// The client supports tags on `SymbolInformation` and `WorkspaceSymbol`.
	// Clients supporting tags have to handle unknown tags gracefully.
	//
	// @since 3.16.0
	TagSupport *SymbolTagSupport `json:"tagSupport,omitempty"`

	// The client supports partial workspace symbols.
	// The client will send the request `workspaceSymbol/resolve` to the server
	// to resolve additional properties.
	//
	// @since 3.17.0 - proposedState
	ResolveSupport *WorkspaceSymbolResolveSupport `json:"resolveSupport,omitempty"`
}

// WorkspaceSymbolResolveSupport provides specific capabilities for
// resolving additional properties for workspace symbols.
type WorkspaceSymbolResolveSupport struct {
	// The properties that a client can resolve lazily.
	// Usually `location.range`.
	Properties []string `json:"properties,omitempty"`
}

// SymbolTagSupport provides specific capabilities for tags on
// symbol objects such as `SymbolInformation` and `WorkspaceSymbol`.
type SymbolTagSupport struct {
	// The tags supported by the client.
	ValueSet []SymbolTag `json:"valueSet"`
}

// SymbolKindCapabilities provides specific capabilities for the `SymbolKind` in
// symbol requests.
type SymbolKindCapabilities struct {
	// The symbol kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	//
	// If this property is not present the client only supports
	// the symbol kinds from `File` to `Array` as defined in
	// the initial version of the protocol.
	ValueSet []SymbolKind `json:"valueSet,omitempty"`
}

// SymbolTag is an extra annotation that tweak rendering of a symbol.
//
// @since 3.16.0
type SymbolTag = Integer

const (
	// SymbolTagDeprecated renders a symbol as obsolete, usually using a strike-out.
	SymbolTagDeprecated SymbolTag = 1
)

// SymbolKind represents a symbol kind in the `workspace/symbol` request.
type SymbolKind = Integer

const (
	// SymbolKindFile represents a file symbol kind.
	SymbolKindFile SymbolKind = 1

	// SymbolKindModule represents a module symbol kind.
	SymbolKindModule SymbolKind = 2

	// SymbolKindNamespace represents a namespace symbol kind.
	SymbolKindNamespace SymbolKind = 3

	// SymbolKindPackage represents a package symbol kind.
	SymbolKindPackage SymbolKind = 4

	// SymbolKindClass represents a class symbol kind.
	SymbolKindClass SymbolKind = 5

	// SymbolKindMethod represents a method symbol kind.
	SymbolKindMethod SymbolKind = 6

	// SymbolKindProperty represents a property symbol kind.
	SymbolKindProperty SymbolKind = 7

	// SymbolKindField represents a field symbol kind.
	SymbolKindField SymbolKind = 8

	// SymbolKindConstructor represents a constructor symbol kind.
	SymbolKindConstructor SymbolKind = 9

	// SymbolKindEnum represents an enum symbol kind.
	SymbolKindEnum SymbolKind = 10

	// SymbolKindInterface represents an interface symbol kind.
	SymbolKindInterface SymbolKind = 11

	// SymbolKindFunction represents a function symbol kind.
	SymbolKindFunction SymbolKind = 12

	// SymbolKindVariable represents a variable symbol kind.
	SymbolKindVariable SymbolKind = 13

	// SymbolKindConstant represents a constant symbol kind.
	SymbolKindConstant SymbolKind = 14

	// SymbolKindString represents a string symbol kind.
	SymbolKindString SymbolKind = 15

	// SymbolKindNumber represents a number symbol kind.
	SymbolKindNumber SymbolKind = 16

	// SymbolKindBoolean represents a boolean symbol kind.
	SymbolKindBoolean SymbolKind = 17

	// SymbolKindArray represents an array symbol kind.
	SymbolKindArray SymbolKind = 18

	// SymbolKindObject represents an object symbol kind.
	SymbolKindObject SymbolKind = 19

	// SymbolKindKey represents a key symbol kind.
	SymbolKindKey SymbolKind = 20

	// SymbolKindNull represents a null symbol kind.
	SymbolKindNull SymbolKind = 21

	// SymbolKindEnumMember represents an enum member symbol kind.
	SymbolKindEnumMember SymbolKind = 22

	// SymbolKindStruct represents a struct symbol kind.
	SymbolKindStruct SymbolKind = 23

	// SymbolKindEvent represents an event symbol kind.
	SymbolKindEvent SymbolKind = 24

	// SymbolKindOperator represents an operator symbol kind.
	SymbolKindOperator SymbolKind = 25

	// SymbolKindTypeParameter represents a type parameter symbol kind.
	SymbolKindTypeParameter SymbolKind = 26
)

// ExecuteCommandClientCapabilities describes the capabilities of a client
// for the `workspace/executeCommand` request.
type ExecuteCommandClientCapabilities struct {
	// Execute command supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

// SemanticTokensWorkspaceClientCapabilities describes the capabilities of a client
// for semantic tokens in workspaces.
type SemanticTokensWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from
	// the server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// semantic tokens currently shown. It should be used with absolute care
	// and is useful for situation where a server for example detect a project
	// wide change that requires such a calculation.
	RefreshSupport *bool `json:"refreshSupport,omitempty"`
}

// CodeLensWorkspaceClientCapabilities describes the capabilities of a client
// for code lenses in workspaces.
type CodeLensWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from the
	// server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// code lenses currently shown. It should be used with absolute care and is
	// useful for situation where a server for example detect a project wide
	// change that requires such a calculation.
	RefreshSupport *bool `json:"refreshSupport,omitempty"`
}

// InlineValueWorkspaceClientCapabilities describes the capabilities of a client
// specific to inline values.
//
// @since 3.17.0
type InlineValueWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from
	// the server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// inline values currently shown. It should be used with absolute care and
	// is useful for situation where a server for example detect a project wide
	// change that requires such a calculation.
	RefreshSupport *bool `json:"refreshSupport,omitempty"`
}

// InlayHintWorkspaceClientCapabilities describes the capabilities of a client
// for inlay hints in workspaces.
//
// @since 3.17.0
type InlayHintWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from
	// the server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// inlay hints currently shown. It should be used with absolute care and
	// is useful for situation where a server for example detects a project wide
	// change that requires such a calculation.
	RefreshSupport *bool `json:"refreshSupport,omitempty"`
}

// DiagnosticWorkspaceClientCapabilities describes the capabilities of a client
// specific to diagnostic pull requests.
//
// @since 3.17.0
type DiagnosticWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from
	// the server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// pulled diagnostics currently shown. It should be used with absolute care
	// and is useful for situation where a server for example detects a project
	// wide change that requires such a calculation.
	RefreshSupport *bool `json:"refreshSupport,omitempty"`
}

// WorkspaceFolder represents a location and name
// of a workspace folder.
type WorkspaceFolder struct {
	// The associated URI for this workspace folder.
	URI URI `json:"uri"`

	// The name of the workspace folder. Used to refer to this
	// workspace folder in the user interface.
	Name string `json:"name"`
}

// ExecuteCommandOptions options to be used for server capabilities
// for executing commands in workspaces.
type ExecuteCommandOptions struct {
	WorkDoneProgressOptions

	// The commands to be executed on the server.
	Commands []string `json:"commands"`
}

// WorkspaceFoldersServerCapabilities describes the capabilities of a server
// for workspace folders.
type WorkspaceFoldersServerCapabilities struct {
	// The server has support for workspace folders.
	Supported *bool `json:"supported,omitempty"`

	// Whether the server wants to receive workspace folder
	// change notifications.
	//
	// If a string is provided, the string is treated as an ID
	// under which the notification is registered on the client
	// side. The ID can be used to unregister for these events
	// using the `client/unregisterCapability` request.
	ChangeNotifications *BoolOrString `json:"changeNotifications,omitempty"`
}

// WorkspaceFoldersClientCapabilities describes the options
// to register for file operations.
type FileOperationRegistrationOptions struct {
	// The actual filters.
	Filters []FileOperationFilter `json:"filters"`
}

// FileOperationFilter dseecribes in which file operation requests
// or notification the server is interested in.
//
// @since 3.16.0
type FileOperationFilter struct {
	// A Uri like `file` or `untitled`.
	Scheme *string `json:"scheme,omitempty"`

	// The actual file operation pattern.
	Pattern FileOperationPattern `json:"pattern"`
}

// FileOperationPattern describes a file operation pattern
// to describe in which file operation requests or notifications
// the server is interested in.
//
// @since 3.16.0
type FileOperationPattern struct {
	// The glob pattern to match. Glob patterns can have the following syntax:
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
	Glob string `json:"glob"`

	// Whether to match files or folders with this pattern.
	//
	// Matches both if undefined.
	Matches *FileOperationPatternKind `json:"matches,omitempty"`

	// Additional options used during matching.
	Options *FileOperationPatternOptions `json:"options,omitempty"`
}

// FileOperationPatternKind describes if a glob pattern matches
// a file or folder.
type FileOperationPatternKind = string

const (
	// FileOperationPatternKindFile matches files only.
	FileOperationPatternKindFile FileOperationPatternKind = "file"

	// FileOperationPatternKindFolder matches folders only.
	FileOperationPatternKindFolder FileOperationPatternKind = "folder"
)

// FileOperationPatternOptions provides matching options for file operation
// patterns.
//
// @since 3.16.0
type FileOperationPatternOptions struct {
	// The pattern should be matched ignoring casing.
	IgnoreCase *bool `json:"ignoreCase,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspace_symbol

const MethodWorkspaceSymbol = Method("workspace/symbol")

// WorkspaceSymbolHandlerFunc is the function signature for the `workspace/symbol`
// method.
//
// Returns: []SymbolInformation | []WorkspaceSymbol | nil
type WorkspaceSymbolHandlerFunc func(
	context *common.LSPContext,
	params *WorkspaceSymbolParams,
) (any, error)

// WorkspaceSymbolParams contains the parameters for the `workspace/symbol` request.
type WorkspaceSymbolParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// A query string to filter symbols by. Clients may send an empty
	// string here to request all symbols.
	Query string `json:"query"`
}

// WorkspaceSymbol represents a special workspace symbol
// that supports locations without a range.
//
// @since 3.17.0
type WorkspaceSymbol struct {
	// The name of this symbol.
	Name string `json:"name"`

	// The kind of this symbol.
	Kind SymbolKind `json:"kind"`

	// Tags for this symbol.
	Tags []SymbolTag `json:"tags,omitempty"`

	// The name of the symbol containing this symbol. This information is for
	// user interface purposes (e.g. to render a qualifier in the user interface
	// if necessary). It can't be used to re-infer a hierarchy for the document
	// symbols.
	ContainerName *string `json:"containerName,omitempty"`

	// The location of this symbol. Whether a server is allowed to
	// return a location without a range depends on the client
	// capability `workspace.symbol.resolveSupport`.
	//
	// See also `SymbolInformation.location`.
	//
	// Location | DocumentURIObject
	Location any `json:"location"`

	// A data entry field that is preserved on a workspace symbol between a
	// workspace symbol request and a workspace symbol resolve request.
	Data LSPAny `json:"data,omitempty"`
}

type workspaceSymbolIntermediary struct {
	Name          string      `json:"name"`
	Kind          SymbolKind  `json:"kind"`
	Tags          []SymbolTag `json:"tags,omitempty"`
	ContainerName *string     `json:"containerName,omitempty"`
	// Location | DocumentURIObject
	Location json.RawMessage `json:"location"`
	Data     LSPAny          `json:"data,omitempty"`
}

// Fulfils the `json.Unmarshaler` interface.
func (w *WorkspaceSymbol) UnmarshalJSON(data []byte) error {
	var intermediary workspaceSymbolIntermediary
	if err := json.Unmarshal(data, &intermediary); err != nil {
		return err
	}

	w.Name = intermediary.Name
	w.Kind = intermediary.Kind
	w.Tags = intermediary.Tags
	w.ContainerName = intermediary.ContainerName
	w.Data = intermediary.Data

	err := w.unmarshalLocation(intermediary.Location)
	return err
}

func (w *WorkspaceSymbol) unmarshalLocation(data json.RawMessage) error {
	var location Location
	if err := json.Unmarshal(data, &location); err == nil && location.Range != nil {
		w.Location = location
		return nil
	}

	var uriObject DocumentURIObject
	err := json.Unmarshal(data, &uriObject)
	if err == nil {
		w.Location = uriObject
		return nil
	}

	return err
}

// DocumentURIObject represents a document URI object
// used in the `WorkspaceSymbol` struct.
type DocumentURIObject struct {
	URI DocumentURI `json:"uri"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspace_symbolResolve

const MethodWorkspaceSymbolResolve = Method("workspaceSymbol/resolve")

// WorkspaceSymbolResolveHandlerFunc is the function signature for the
// `workspaceSymbol/resolve` method.
type WorkspaceSymbolResolveHandlerFunc func(
	context *common.LSPContext,
	params *WorkspaceSymbol,
) (*WorkspaceSymbol, error)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspace_configuration

const MethodWorkspaceConfiguration = Method("workspace/configuration")

// ConfigurationParams contains the parameters for the `workspace/configuration` request
// made from the server to the client.
type ConfigurationParams struct {
	Items []ConfigurationItem `json:"items"`
}

// ConfigurationItem is a workspace configuration item to be fetched.
type ConfigurationItem struct {
	// The scope to get the configuration section for.
	ScopeURI *URI `json:"scopeUri,omitempty"`

	// The configuration section asked for.
	Section *string `json:"section,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspace_didChangeConfiguration

const MethodWorkspaceDidChangeConfiguration = Method("workspace/didChangeConfiguration")

// WorkspaceDidChangeConfigurationHandlerFunc is the function signature for the
// `workspace/didChangeConfiguration` method.
type WorkspaceDidChangeConfigurationHandlerFunc func(
	context *common.LSPContext,
	params *DidChangeConfigurationParams,
) error

// DidChangeConfigurationParams contains the parameters for the `workspace/didChangeConfiguration`
// notification.
type DidChangeConfigurationParams struct {
	// The actual changed settings.
	Settings LSPAny `json:"settings"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspace_workspaceFolders

const MethodWorkspaceFolders = Method("workspace/workspaceFolders")

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspace_didChangeWorkspaceFolders

const MethodWorkspaceDidChangeFolders = Method("workspace/didChangeWorkspaceFolders")

// WorkspaceDidChangeFoldersHandlerFunc is the function signature for the
// `workspace/didChangeWorkspaceFolders` method.
type WorkspaceDidChangeFoldersHandlerFunc func(
	context *common.LSPContext,
	params *DidChangeWorkspaceFoldersParams,
) error

// DidChangeWorkspaceFoldersParams contains the parameters for the `workspace/didChangeWorkspaceFolders`
// notification.
type DidChangeWorkspaceFoldersParams struct {
	// The actual workspace folder change event.
	Event WorkspaceFoldersChangeEvent `json:"event"`
}

// WorkspaceFoldersChangeEvent represents a change event for workspace folders.
type WorkspaceFoldersChangeEvent struct {
	// The array of added workspace folders.
	Added []WorkspaceFolder `json:"added"`

	// The array of removed workspace folders.
	Removed []WorkspaceFolder `json:"removed"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspace_willCreateFiles

const MethodWorkspaceWillCreateFiles = Method("workspace/willCreateFiles")

// WorkspaceWillCreateFilesHandlerFunc is the function signature for the
// `workspace/willCreateFiles` method.
type WorkspaceWillCreateFilesHandlerFunc func(
	context *common.LSPContext,
	params *CreateFilesParams,
) (*WorkspaceEdit, error)

// CreateFilesParams contains the parameters sent in notifications/requests
// for user-initiated creation of files.
//
// @since 3.16.0
type CreateFilesParams struct {
	// An array of all files/folders created in this operation.
	Files []FileCreate `json:"files"`
}

// FileCreate represents information on a file/folder creation.
//
// @since 3.16.0
type FileCreate struct {
	// A file:// URI for the location of the file/folder being created.
	URI string `json:"uri"`
}
