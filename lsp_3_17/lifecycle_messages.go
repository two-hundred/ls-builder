package lsp

import "github.com/two-hundred/ls-builder/common"

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialize

const MethodInitialize = Method("initialize")

// InitializeFunc is the function signature for the initialize request
// handler that can be registered for a language server.
type InitializeFunc func(ctx common.LSPContext, params *InitializeParams) (any, error)

// InitializeParams contains the initialize request parameters.
type InitializeParams struct {
	WorkDoneProgressParams

	// The process ID of the parent process that started the server.
	// Is null if the process has not been started by another process.
	// If the parent process is not alive then the server should exit
	// (see exit notification) its process.
	ProcessID *Integer `json:"processId"`

	// Information about the client.
	//
	// @since 3.15.0
	ClientInfo *InitializeClientInfo `json:"clientInfo,omitempty"`

	// The locale that the client is currently showing the
	// user interface in. This must not necessarily be the
	// locale of the operating system.
	//
	// Uses IETF language tags as the value's syntax
	// (See https://en.wikipedia.org/wiki/IETF_language_tag)
	//
	// @since 3.16.0
	Locale *string `json:"locale,omitempty"`

	// The rootPath of the workspace. Is null
	// if no folder is open.
	//
	// @deprecated in favour of `rootUri`.
	RootPath *string `json:"rootPath,omitempty"`

	// The rootUri of the workspace. Is null if no
	// folder is open. If both `rootPath` and `rootUri` are set
	// `rootUri` wins.
	//
	// @deprecated in favour of `workspaceFolders`
	RootURI *DocumentURI `json:"rootUri"`

	// User provided initialization options.
	InitializationOptions LSPAny `json:"initializationOptions,omitempty"`

	// The capabilities provided by the client (editor or tool)
	Capabilities ClientCapabilities `json:"capabilities"`

	// The initial trace setting. If omitted trace is disabled ('off').
	Trace *TraceValue `json:"trace,omitempty"`

	// The workspace folders configured in the client when the server starts.
	// This property is only available if the client supports workspace folders.
	// It can be `null` if the client supports workspace folders but none are
	// configured.
	//
	// @since 3.6.0
	WorkspaceFolders []WorkspaceFolder `json:"workspaceFolders,omitempty"`
}

// InitializeClientInfo represents information about the client.
type InitializeClientInfo struct {

	// The name of the client as defined by the client.
	Name string `json:"name"`

	// The client's version as defined by the client.
	Version string `json:"version,omitempty"`
}

// ClientCapabilities represents the capabilities of the client (editor or tool).
type ClientCapabilities struct {
	// Workspace specific client capabilities.
	Workspace *ClientWorkspaceCapabilities `json:"workspace,omitempty"`

	// Text document specific client capabilities.
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`

	// Capabilities specific to the notebook document support.
	//
	// @since 3.17.0
	NotebookDocument *NotebookDocumentClientCapabilities `json:"notebook,omitempty"`

	// Window specific client capabilities.
	Window *WindowClientCapabilities `json:"window,omitempty"`

	// General client capabilities.
	//
	// @since 3.16.0
	General *GeneralClientCapabilities `json:"general,omitempty"`

	// Experimenal client capabilities.
	Experimental LSPAny `json:"experimental,omitempty"`
}

// ClientWorkspaceCapabilities represents the capabilities of the client
// related to workspaces.
type ClientWorkspaceCapabilities struct {
	// the cliet supports applying batch edits to the workspace
	// by supporting the request `workspace/applyEdit`.
	ApplyEdit *bool `json:"applyEdit,omitempty"`

	// Capabilities specific to `WorkspaceEdit`s
	WorkspaceEdit *WorkspaceEditClientCapabilities `json:"workspaceEdit,omitempty"`

	// Capabilities specific to the `workspace/didChangeConfiguration` notification.
	DidChangeConfiguration *DidChangeConfigurationClientCapabilities `json:"didChangeConfiguration,omitempty"`

	// Capabilities specific to the `workspace/didChangeWatchedFiles` notification.
	DidChangeWatchedFiles *DidChangeWatchedFilesClientCapabilities `json:"didChangeWatchedFiles,omitempty"`

	// Capabilities specific to the `workspace/symbol` request.
	Symbol *WorkspaceSymbolClientCapabilities `json:"symbol,omitempty"`

	// Capabilities specific to the `workspace/executeCommand` request.
	ExecuteCommand *ExecuteCommandClientCapabilities `json:"executeCommand,omitempty"`

	// The client has support for workspace folders.
	//
	// @since 3.6.0
	WorkspaceFolders *bool `json:"workspaceFolders,omitempty"`

	// The client supports `workspace/configuration` requests.
	//
	// @since 3.6.0
	Configuration *bool `json:"configuration,omitempty"`

	// Capabilities specific to the semantic token requiests scoped to the
	// workspace.
	//
	// @since 3.16.0
	SemanticTokens *SemanticTokensWorkspaceClientCapabilities `json:"semanticTokens,omitempty"`

	// Capabilities specific to the code lens requests scoped to the workspace.
	//
	// @since 3.16.0
	CodeLens *CodeLensWorkspaceClientCapabilities `json:"codeLens,omitempty"`

	// The client has support for file requests/notifications.
	//
	// @since 3.16.0
	FileOperations *FileOperationClientCapabilities `json:"fileOperations,omitempty"`

	// Client workspace capabilities specific to inline values.
	//
	// @since 3.17.0
	InlineValue *InlineValueWorkspaceClientCapabilities `json:"inlineValue,omitempty"`

	// Client workspace capabilities specific to inlay hints.
	//
	// @since 3.17.0
	InlayHint *InlayHintWorkspaceClientCapabilities `json:"inlayHint,omitempty"`

	// Client workspace capabilities specific to diagnostics.
	//
	// @since 3.17.0
	Diagnostics *DiagnosticWorkspaceClientCapabilities `json:"diagnostics,omitempty"`
}

// FileOperationClientCapabilities represents the capabilities of the client
// related to file operations.
type FileOperationClientCapabilities struct {

	// Whether the client supports dynamic regristration for file requests/notifications.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// The client has support for sending didCreateFiles notifications.
	DidCreate *bool `json:"didCreate,omitempty"`

	// The client has support for sending willCreateFiles requests.
	WillCreate *bool `json:"willCreate,omitempty"`

	// The client has support for sending didRenameFiles notifications.
	DidRename *bool `json:"didRename,omitempty"`

	// The client has support for sending willRenameFiles requests.
	WillRename *bool `json:"willRename,omitempty"`

	// The client has support for sending didDeleteFiles notifications.
	DidDelete *bool `json:"didDelete,omitempty"`

	// The client has support for sending willDeleteFiles requests.
	WillDelete *bool `json:"willDelete,omitempty"`
}

// TextDocumentClientCapabilities represents the capabilities of the client
// related to text documents.
type TextDocumentClientCapabilities struct {
	Synchronization *TextDocumentSyncClientCapabilities `json:"synchronization,omitempty"`

	// Capabilities specific to the `textDocument/completion` request.
	Completion *CompletionClientCapabilities `json:"completion,omitempty"`

	// Capabilities specific to the `textDocument/hover` request.
	Hover *HoverClientCapabilities `json:"hover,omitempty"`

	// Capabilities specific to the `textDocument/signatureHelp` request.
	SignatureHelp *SignatureHelpClientCapabilities `json:"signatureHelp,omitempty"`

	// Capabilities specific to the `textDocument/declaration` request.
	//
	// @since 3.14.0
	Declaration *DeclarationClientCapabilities `json:"declaration,omitempty"`

	// Capabilities specific to the `textDocument/definition` request.
	Definition *DefinitionClientCapabilities `json:"definition,omitempty"`

	// Capabilities specific to the `textDocument/typeDefinition` request.
	//
	// @since 3.6.0
	TypeDefinition *TypeDefinitionClientCapabilities `json:"typeDefinition,omitempty"`

	// Capabilities specific to the `textDocument/implementation` request.
	//
	// @since 3.6.0
	Implementation *ImplementationClientCapabilities `json:"implementation,omitempty"`

	// Capabilities specific to the `textDocument/references` request.
	//
	// @since 3.6.0
	References *ReferenceClientCapabilities `json:"references,omitempty"`

	// Capabilities specific to the `textDocument/documentHighlight` request.
	DocumentHighlight *DocumentHighlightClientCapabilities `json:"documentHighlight,omitempty"`

	// Capabilities specific to the `textDocument/documentSymbol` request.
	DocumentSymbol *DocumentSymbolClientCapabilities `json:"documentSymbol,omitempty"`

	// Capabilities specific to the `textDocument/codeAction` request.
	CodeAction *CodeActionClientCapabilities `json:"codeAction,omitempty"`

	// Capabilities specific to the `textDocument/codeLens` request.
	CodeLens *CodeLensClientCapabilities `json:"codeLens,omitempty"`

	// Capabilities specific to the `textDocument/documentLink` request.
	DocumentLink *DocumentLinkClientCapabilities `json:"documentLink,omitempty"`

	// Capabilities specific to the `textDocument/documentColor` and the
	// `textDocument/colorPresentation` request.
	//
	// @since 3.6.0
	ColorProvider *DocumentColorClientCapabilities `json:"colorProvider,omitempty"`

	// Capabilities specific to the `textDocument/formatting` request.
	Formatting *DocumentFormattingClientCapabilities `json:"formatting,omitempty"`

	// Capabilities specific to the `textDocument/rangeFormatting` request.
	RangeFormatting *DocumentRangeFormattingClientCapabilities `json:"rangeFormatting,omitempty"`

	// Capabilities specific to the `textDocument/onTypeFormatting` request.
	OnTypeFormatting *DocumentOnTypeFormattingClientCapabilities `json:"onTypeFormatting,omitempty"`

	// Capabilities specific to the `textDocument/rename` request.
	Rename *RenameClientCapabilities `json:"rename,omitempty"`

	// Capabilities specific to the `textDocument/publishDiagnostics` notification.
	PublishDiagnostics *PublishDiagnosticsClientCapabilities `json:"publishDiagnostics,omitempty"`

	// Capabilities specific to the `textDocument/foldingRange` request.
	//
	// @since 3.10.0
	FoldingRange *FoldingRangeClientCapabilities `json:"foldingRange,omitempty"`

	// Capabilities specific to the `textDocument/selectionRange` request.
	//
	// @since 3.15.0
	SelectionRange *SelectionRangeClientCapabilities `json:"selectionRange,omitempty"`

	// Capabilities specific to the `textDocument/linkedEditingRange` request.
	//
	// @since 3.16.0
	LinkedEditingRange *LinkedEditingRangeClientCapabilities `json:"linkedEditingRange,omitempty"`

	// Capabilities specific to the various call hierarchy requests.
	//
	// @since 3.16.0
	CallHierarchy *CallHierarchyClientCapabilities `json:"callHierarchy,omitempty"`

	// Capabilities specific to the `textDocument/semanticTokens` request.
	//
	// @since 3.16.0
	SemanticTokens *SemanticTokensClientCapabilities `json:"semanticTokens,omitempty"`

	// Capabilities specific to the `textDocument/moniker` request.
	//
	// @since 3.16.0
	Moniker *MonikerClientCapabilities `json:"moniker,omitempty"`

	// Capabilities specific to the various type hierarchy requests.
	//
	// @since 3.17.0
	TypeHierarchy *TypeHierarchyClientCapabilities `json:"typeHierarchy,omitempty"`

	// Capabilities specific to the `textDocument/inlineValue` request.
	//
	// @since 3.17.0
	InlineValue *InlineValueClientCapabilities `json:"inlineValue,omitempty"`

	// Capabilities specific to the `textDocument/inlayHint` request.
	//
	// @since 3.17.0
	InlayHint *InlayHintClientCapabilities `json:"inlayHint,omitempty"`

	// Capabilities specific to the diagnostic pull model.
	//
	// @since 3.17.0
	Diagnostics *DiagnosticClientCapabilities `json:"diagnostics,omitempty"`
}

// NotebookDocumentClientCapabilities provides
// capabilities specific to the notebook document support.
//
// @since 3.17.0
type NotebookDocumentClientCapabilities struct {
	// Capabilities specific to notebook document synchronization.
	//
	// @since 3.17.0
	Synchronization *NotebookDocumentSyncClientCapabilities `json:"synchronization,omitempty"`
}

// WindowClientCapabilities represents the capabilities of the client
// related to the window.
type WindowClientCapabilities struct {

	// Indicates whether the client supports server initiated
	// progress using the `window/workDoneProgress/create` request.
	//
	// The capability also controls whether the client supports
	// handling of progress notifications. If set, servers are allowed
	// to report a `workDoneProgress` property in the request specific
	// server capabilities.
	WorkDoneProgress *bool `json:"workDoneProgress,omitempty"`

	// Capabilities specific to the `window/showMessage` request.
	//
	// @since 3.16.0
	ShowMessage *ShowMessageRequestClientCapabilities `json:"showMessage,omitempty"`

	// Client capabilities for the show document request.
	//
	// @since 3.16.0
	ShowDocument *ShowDocumentClientCapabilities `json:"showDocument,omitempty"`
}

// GeneralClientCapabilities represents the general capabilities of the client.
//
// @since 3.16.0
type GeneralClientCapabilities struct {
	// Client capability that signals how the client handles stale requests
	// (e.g. a request for which the client will not process the response anymor
	// since teh information is outdated).
	//
	// @since 3.17.0
	StaleRequestSupport *StaleRequestSupport `json:"staleRequestSupport,omitempty"`

	// Client capabilities specific to regular expressions.
	//
	// @since 3.16.0
	RegularExpressions *RegularExpressionsClientCapabilities `json:"regularExpressions,omitempty"`

	// Client capabilities specific to the client's markdown parser.
	//
	// @since 3.16.0
	Markdown *MarkdownClientCapabilities `json:"markdown,omitempty"`

	// The position encodings supported by the client. Client and server
	// have to agree on the same position encoding to ensure that offsets
	// (e.g. character position in a line) are interpreted the same on both
	// side.
	//
	// To keep the protocol backwards compatible the following applies: if
	// the value 'utf-16' is missing from the array of position encodings
	// servers can assume that the client supports UTF-16. UTF-16 is
	// therefore a mandatory encoding.
	//
	// If omitted it defaults to ['utf-16'].
	//
	// Implementation considerations: since the conversion from one encoding
	// into another requires the content of the file / line the conversion
	// is best done where the file is read which is usually on the server
	// side.
	//
	// @since 3.17.0
	PositionEncodings []PositionEncodingKind `json:"positionEncodings,omitempty"`
}

// StaleRequestSupport represents how the client handles stale requests.
//
// @since 3.17.0
type StaleRequestSupport struct {
	// The client will actively cancel the request.
	Cancel bool `json:"cancel"`

	// This list of requests for which the client will retry the request
	// if it receives a response with error code `ContentModified`.
	RetryOnContentModified []string `json:"retryOnContentModified"`
}

// InitializeResult represents the result of the initialize request.
type InitializeResult struct {
	// The capabilities the language server provides.
	Capabilities ServerCapabilities `json:"capabilities"`

	// Information about the server.
	//
	// @since 3.15.0
	ServerInfo *InitializeResultServerInfo `json:"serverInfo,omitempty"`
}

// InitializeResultServerInfo represents information about the server.
type InitializeResultServerInfo struct {

	// The name of the server as defined by the server.
	Name string `json:"name"`

	// The server's version as defined by the server.
	Version *string `json:"version,omitempty"`
}

// Known error codes for an `InitializeError`;
type InitializeErrorCode Integer

const (
	// If the protocol version provided by the client can't be handled by the
	// server.
	//
	// @deprecated This initialize error got replaced by client capabilities.
	// There is no version handshake in version 3.0x
	InitializeErrorCodeUnknownProtocolVersion = InitializeErrorCode(1)
)

// InitializeError represents an error that occurred during the initialize request.
type InitializeError struct {
	// Indicates whether the client execute the following retry logic:
	// (1) show the message provided by the ResponseError to the user
	// (2) user selects retry or cancel
	// (3) if user selected retry the initialize method is sent again.
	Retry bool `json:"retry"`
}

// ServerCapabilities represents the capabilities of the server
// returned in the initialize result.
type ServerCapabilities struct{}
