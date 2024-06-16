package lsp

import (
	"encoding/json"

	"github.com/two-hundred/ls-builder/common"
)

// TextDocumentSyncClientCapabilities represents the client capabilities
// specific to text document synchronisation.
type TextDocumentSyncClientCapabilities struct {
	// Whether text document synchronization supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	// The client supports sending will save notifications.
	WillSave *bool `json:"willSave,omitempty"`

	// The client supports sending a will save request and
	// waits for a response providing text edits which will
	// be applied to the document before it is saved.
	WillSaveWaitUntil *bool `json:"willSaveWaitUntil,omitempty"`

	// The client supports did save notifications.
	DidSave *bool `json:"didSave,omitempty"`
}

func unmarshalTextDocumentSyncServerCapabilities(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {

	err := unmarshalServerCapabilityTextDocumentSync(serverCapabilities, intermediate)
	if err != nil {
		return err
	}

	return nil
}

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

// unmarshals the TextDocumentSync
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityTextDocumentSync(
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

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_didOpen

const MethodTextDocumentDidOpen = Method("textDocument/didOpen")

// TextDocumentDidOpenHandlerFunc is the function signature for the handler
// of the textDocument/didOpen notification.
type TextDocumentDidOpenHandlerFunc func(ctx *common.LSPContext, params *DidOpenTextDocumentParams) error

// DidOpenTextDocumentParams contains the parameters of the textDocument/didOpen notification.
type DidOpenTextDocumentParams struct {
	// The document that was opened.
	TextDocument TextDocumentItem `json:"textDocument"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_didChange

const MethodTextDocumentDidChange = Method("textDocument/didChange")

// TextDocumentDidChangeHandlerFunc is the function signature for the handler
// of the textDocument/didChange notification.
type TextDocumentDidChangeHandlerFunc func(ctx *common.LSPContext, params *DidChangeTextDocumentParams) error

// DidChangeTextDocumentParams contains the
// parameters of the textDocument/didChange notification.
type DidChangeTextDocumentParams struct {
	// The document that did change. The version number points
	// to the version after all provided content changes have
	// been applied.
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

	// The actual content changes. The content changes describe single state
	// changes to the document. So if there are two content changes c1 (at
	// array index 0) and c2 (at array index 1) for a document in state S then
	// c1 moves the document from S to S' and c2 from S' to S''. So c1 is
	// computed on the state S and c2 is computed on the state S'.
	//
	// To mirror the content of a document using change events use the following
	// approach:
	// - start with the same initial content
	// - apply the 'textDocument/didChange' notifications in the order you
	//   receive them.
	// - apply the `TextDocumentContentChangeEvent`s in a single notification
	//   in the order you receive them.
	//
	// TextDocumentContentChangeEvent | TextDocumentContentChangeEventWhole
	ContentChanges []any `json:"contentChanges"`
}

type didChangeTextDocumentParamsIntermediate struct {
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`
	// TextDocumentContentChangeEvent | TextDocumentContentChangeEventWhole
	ContentChanges []json.RawMessage `json:"contentChanges"`
}

// Fulfils the json.Unmarshaler interface.
func (p *DidChangeTextDocumentParams) UnmarshalJSON(data []byte) error {
	var intermediate didChangeTextDocumentParamsIntermediate
	if err := json.Unmarshal(data, &intermediate); err != nil {
		return err
	}

	p.TextDocument = intermediate.TextDocument

	for _, raw := range intermediate.ContentChanges {
		var changeEvent TextDocumentContentChangeEvent
		if err := json.Unmarshal(raw, &changeEvent); err == nil {
			if changeEvent.Range != nil {
				p.ContentChanges = append(p.ContentChanges, changeEvent)
			} else {
				changeEventWhole := TextDocumentContentChangeEventWhole{
					Text: changeEvent.Text,
				}
				p.ContentChanges = append(p.ContentChanges, changeEventWhole)
			}
		} else {
			return err
		}
	}

	return nil
}

// TextDocumentContentChangeEvent represents
// an event describing a change to a text document. If only a text is provided
// it is considered to be the full content of the document.
type TextDocumentContentChangeEvent struct {
	// The range of the document that changed.
	Range *Range `json:"range"`

	// The optional length of the range that was replaced.
	//
	// @deprecated use range instead.
	RangeLength *UInteger `json:"rangeLength,omitempty"`

	// The new text for the provided range.
	Text string `json:"text"`
}

// TextDocumentContentChangeEventWhole represents
// an event describing a change to a text document where
// the full content of the document is provided.
type TextDocumentContentChangeEventWhole struct {
	// The new text of the whole document.
	Text string `json:"text"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_willSave

const MethodTextDocumentWillSave = Method("textDocument/willSave")

// TextDocumentWillSaveHandlerFunc is the function signature for the handler
// of the textDocument/willSave notification.
type TextDocumentWillSaveHandlerFunc func(
	ctx *common.LSPContext,
	params *WillSaveTextDocumentParams,
) error

// WillSaveTextDocumentParams are the parameters of a will save
// text document notification.
type WillSaveTextDocumentParams struct {
	// The document that will be saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The reason why the document is being saved.
	Reason TextDocumentSaveReason `json:"reason"`
}

// TextDocumentSaveReason represents the reason why a text document is being saved.
type TextDocumentSaveReason Integer

const (
	// TextDocumentSaveReasonManual means that the document save is manually triggered,
	// e.g. by the user pressing save, by starting debugging, or by an API call.
	TextDocumentSaveReasonManual TextDocumentSaveReason = 1

	// TextDocumentSaveReasonAfterDelay means that the document save is triggered
	// automatically after a delay.
	TextDocumentSaveReasonAfterDelay TextDocumentSaveReason = 2

	// TextDocumentSaveReasonFocusOut means that the document save is triggered when
	// the editor loses focus.
	TextDocumentSaveReasonFocusOut TextDocumentSaveReason = 3
)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_willSaveWaitUntil

const MethodTextDocumentWillSaveWaitUntil = Method("textDocument/willSaveWaitUntil")

// TextDocumentWillSaveWaitUntilHandlerFunc is the function signature for the handler
// of the textDocument/willSaveWaitUntil request.
type TextDocumentWillSaveWaitUntilHandlerFunc func(
	ctx *common.LSPContext,
	params *WillSaveTextDocumentParams,
) ([]TextEdit, error)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_didSave

const MethodTextDocumentDidSave = Method("textDocument/didSave")

// TextDocumentDidSaveHandlerFunc is the function signature for the handler
// of the textDocument/didSave notification.
type TextDocumentDidSaveHandlerFunc func(
	ctx *common.LSPContext,
	params *DidSaveTextDocumentParams,
) error

// DidSaveTextDocumentParams are the parameters of a did save
// text document notification.
type DidSaveTextDocumentParams struct {
	// The document that was saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Optional, contains the content when saved.
	// Whether this is present or not depends on the includeText value
	// when the save notification was requested.
	Text *string `json:"text,omitempty"`
}
