package lsp

import "github.com/two-hundred/ls-builder/common"

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#setTrace

// MethodSetTrace is the method name for the setTrace notification
// as defined in the language server protocol.
const MethodSetTrace = Method("$/setTrace")

// SetTraceHandlerFunc is the function signature for the setTrace notification
// handler that can be registered for a language server.
type SetTraceHandlerFunc func(ctx *common.LSPContext, params *SetTraceParams) error

// SetTraceParams contains the setTrace notification parameters.
type SetTraceParams struct {
	// The new value that should be assigned to the trace setting.
	Value TraceValue `json:"value"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#logTrace

// MethodLogTrace is the method name for the logTrace notification
// as defined in the language server protocol.
const MethodLogTrace = Method("$/logTrace")

// LogTraceParams contains the logTrace notification parameters.
type LogTraceParams struct {
	// The message to be logged.
	Message string `json:"message"`

	// Additional information that can be computed if the `trace` configuration
	// is set to `verbose`.
	Verbose *string `json:"verbose,omitempty"`
}
