package lsp

// PublishDiagnosticsClientCapabilities represents the client capabilities
// specific to diagnostics sent from the server to the client to signal
// results of validation runs.
type PublishDiagnosticsClientCapabilities struct {
	// Whether the clients accepts diagnostics with related information.
	RelatedInformation *bool `json:"relatedInformation,omitempty"`

	// Client supports the tag property to provide meta data about a diagnostic.
	// Clients supporting tags have to handle unknown tags gracefully.
	//
	// @since 3.15.0
	TagSupport *DiagnosticTagSupport `json:"tagSupport,omitempty"`

	// Whether the client interprets the version property of the
	// `textDocument/publishDiagnostics` notification's parameter.
	//
	// @since 3.15.0
	VersionSupport *bool `json:"versionSupport,omitempty"`

	// Client supports a codeDescription property
	//
	// @since 3.16.0
	CodeDescriptionSupport *bool `json:"codeDescriptionSupport,omitempty"`

	// Whether code action supports the `data` property which is
	// preserved between a `textDocument/publishDiagnostics` and
	// `textDocument/codeAction` request.
	//
	// @since 3.16.0
	DataSupport *bool `json:"dataSupport,omitempty"`
}

// DiagnosticTagSupport represents the client capabilities specific to
// diagnostic tags.
type DiagnosticTagSupport struct {
	// The tags supported by the client.
	ValueSet []DiagnosticTag `json:"valueSet"`
}

// DiagnosticClientCapabilities represents the client capabilities
// specific to diagnostic pull requests.
//
// @since 3.17.0
type DiagnosticClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to
	// `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// Whether the clients support related documents for document diagnostic
	// pulls.
	RelatedDocumentSupport *bool `json:"relatedDocumentSupport,omitempty"`
}
