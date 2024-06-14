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
