package lsp

type DidChangeConfigurationClientCapabilities struct{}

type DidChangeWatchedFilesClientCapabilities struct{}

type WorkspaceSymbolClientCapabilities struct{}

type ExecuteCommandClientCapabilities struct{}

type SemanticTokensWorkspaceClientCapabilities struct{}

type CodeLensWorkspaceClientCapabilities struct{}

type InlineValueWorkspaceClientCapabilities struct{}

type InlayHintWorkspaceClientCapabilities struct{}

type DiagnosticWorkspaceClientCapabilities struct{}

type WorkspaceFolder struct{}

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
