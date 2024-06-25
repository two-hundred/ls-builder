package lsp

// ShowMessageRequestClientCapabilities represents the client capabilities
// specific to the show message request.
type ShowMessageRequestClientCapabilities struct {
	// Capabilities specific to the `MessageActionItem` type.
	MessageActionItem *MessageActionItemClientCapabilities `json:"messageActionItem,omitempty"`
}

// MessageActionItemClientCapabilities represents the client capabilities
// specific to the message action item.
type MessageActionItemClientCapabilities struct {
	// Whether the client supports additional attributes which
	// are preserved and sent back to the server in the
	// request's response.
	AdditionalPropertiesSupport *bool `json:"additionalPropertiesSupport,omitempty"`
}

// ShowDocumentClientCapabilities represents the client capabilities
// specific to the show document request.
//
// @since 3.16.0
type ShowDocumentClientCapabilities struct {
	// The client has support for the show document request.
	Support *bool `json:"support,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#window_showMessage

const MethodShowMessageNotification = Method("window/showMessage")

// ShowMessageParams represents the parameters of a `window/showMessage` notification.
type ShowMessageParams struct {
	// The message type. See {@link MessageType}.
	Type MessageType `json:"type"`

	// The actual message.
	Message string `json:"message"`
}

// MessageType defins the type of window message.
type MessageType = Integer

const (
	// MessageTypeError represents an error message.
	MessageTypeError MessageType = 1

	// MessageTypeWarning represents a warning message.
	MessageTypeWarning MessageType = 2

	// MessageTypeInfo represents an information message.
	MessageTypeInfo MessageType = 3

	// MessageTypeLog represents a log message.
	MessageTypeLog MessageType = 4

	// MessageTypeDebug represents a debug message.
	//
	// @since 3.18.0
	// @proposed
	MessageTypeDebug MessageType = 5
)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#window_showMessageRequest

const MethodShowMessageRequest = Method("window/showMessageRequest")

// ShowMessageRequestParams represents the parameters of a `window/showMessageRequest` notification.
type ShowMessageRequestParams struct {
	// The message type. See {@link MessageType}.
	Type MessageType `json:"type"`

	// The actual message.
	Message string `json:"message"`

	// The message action items to present.
	Actions []MessageActionItem `json:"actions,omitempty"`
}

// MessageActionItem represents an additional message action
// includes in a `window/showMessageRequest` requests.
type MessageActionItem struct {
	// A short title like 'Retry', 'Open Log' etc.
	Title string `json:"title"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#window_logMessage

const MethodLogMessage = Method("window/logMessage")

// LogMessageParams represents the parameters of a `window/logMessage` notification.
type LogMessageParams struct {
	// The message type. See {@link MessageType}.
	Type MessageType `json:"type"`

	// The actual message.
	Message string `json:"message"`
}
