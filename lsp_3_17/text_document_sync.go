package lsp

import "encoding/json"

type TextDocumentSyncClientCapabilities struct{}

// TextDocumentSyncOptions represents the options of a text document sync
// capability.
type TextDocumentSyncOptions struct {
	// Open and close notifications are sent to the server.
	// If omitted open close notifications should not be sent.
	OpenClose *bool `json:"openClose,omitempty"`

	// Change notifications are sent to the server.
	// See TextDocumentSyncKind.None, TextDocumentSyncKind.Full and
	// TextDocumentSyncKind.Incremental. If omitted it defaults to
	// TextDocumentSyncKind.None.
	Change *TextDocumentSyncKind `json:"change,omitempty"`
}

// TextDocumentSyncKind defines how the host (editor) should sync document changes
// to the language server.
type TextDocumentSyncKind Integer

var (
	// TextDocumentSyncKindNone means that documents should
	// not be synced at all.
	TextDocumentSyncKindNone TextDocumentSyncKind = 0

	// TextDocumentSyncKindFull means that documents are synced
	// by always sending the full content of the document.
	TextDocumentSyncKindFull TextDocumentSyncKind = 1

	// TextDocumentSyncKindIncremental means that documents are
	// synced by sending the full content on open. After that only
	// incremental updates to the document are sent.
	TextDocumentSyncKindIncremental TextDocumentSyncKind = 2
)

// UnmarshalServerCapabilityTextDocumentSync unmarshals the
// TextDocumentSync field of a server capabilities object.
// This modifies the serverCapabilities object.
func UnmarshalServerCapabilityTextDocumentSync(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.TextDocumentSync == nil {
		return nil
	}

	var optVal TextDocumentSyncOptions
	if err := json.Unmarshal(intermediate.TextDocumentSync, &optVal); err == nil {
		serverCapabilities.TextDocumentSync = optVal
	} else {
		var kindVal TextDocumentSyncKind
		if err := json.Unmarshal(intermediate.TextDocumentSync, &kindVal); err == nil {
			serverCapabilities.TextDocumentSync = kindVal
		} else {
			return err
		}
	}

	return nil
}
