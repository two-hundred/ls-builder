package lsp

import (
	"encoding/json"
)

type NotebookDocumentSyncClientCapabilities struct{}

func unmarshalNotebookDocumentSyncServerCapabilities(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {

	err := unmarshalServerCapabilityNotebookDocumentSync(serverCapabilities, intermediate)
	if err != nil {
		return err
	}

	return nil
}

// NotebookDocumentSyncOptions describes the options
// specific to a notebook plus its cells to be synced
// to the server.
//
// If a selector provides a notebook document filter
// but no cell selector all cells of a matching notebook document
// will be synced.
//
// If a selector provides no notebook document filter
// but only a cell selector all notebook documents
// that contain at least one matching cell will be synced.
//
// @since 3.17.0
type NotebookDocumentSyncOptions struct {
	// The notebooks to be synced.
	NotebookSelector []*NotebookSelectorItem `json:"notebookSelector"`

	// Whether save notification should be forwarded to
	// the server. Will only be honoured if mode === 'notebook'.
	Save *bool `json:"save,omitempty"`
}

type notebookDocumentSyncOptionsIntermediate struct {
	NotebookSelector []*notebookSelectorIntermediate `json:"notebookSelector"`
	Save             *bool                           `json:"save,omitempty"`
}

// NotebookSelectorItem is the selector for a notebook to be synced.
type NotebookSelectorItem struct {
	// The notebook to be synced. If a string
	// value is provided it matches against the
	// notebook type. '*' matches every notebook.
	// NotebookDocumentFilter | string
	Notebook any `json:"notebook"`

	// The cells of the matching notebook to be synced.
	Cells []NotebookCellLanguage `json:"cells,omitempty"`
}

// NotebookCellLanguage denotes the language of a notebook cell.
type NotebookCellLanguage struct {
	Language string `json:"language"`
}

type notebookSelectorIntermediate struct {
	Notebook json.RawMessage        `json:"notebook"`
	Cells    []NotebookCellLanguage `json:"cells,omitempty"`
}

// NotebookDocumentFilter denotes a notebook document by
// different properties.
//
// @since 3.17.0
type NotebookDocumentFilter struct {
	// The type of the enclosing notebook.
	NotebookType string `json:"notebookType"`

	// A URI [scheme](#Uri.scheme) like `file` or `untitled`.
	Scheme *string `json:"scheme,omitempty"`

	// A glob pattern.
	Pattern *string `json:"pattern,omitempty"`
}

// Registration options specific to a notebook.
//
// @since 3.17.0
type NotebookDocumentSyncRegistrationOptions struct {
	NotebookDocumentSyncOptions
	StaticRegistrationOptions
}

type notebookDocumentSyncRegOptionsIntermediate struct {
	notebookDocumentSyncOptionsIntermediate
	StaticRegistrationOptions
}

// unmarshals the NotebookDocumentSync
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityNotebookDocumentSync(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.NotebookDocumentSync == nil {
		return nil
	}

	// Try with static registration options as is a superset of
	// the notebook document sync options.
	// As the ID for a static registration
	// is optional, JSON unmarshalling will succeed to match
	// "notebookDocumentSyncRegOptionsIntermediate" even when it is not provided.
	// For this reason, standard notebook document sync options will be represented by a
	// "NotebookDocumenySyncRegistrationOptions" struct with a nil ID.
	var optValRegIntermediate notebookDocumentSyncRegOptionsIntermediate
	if err := json.Unmarshal(intermediate.NotebookDocumentSync, &optValRegIntermediate); err == nil {
		targetRegOptVal := &NotebookDocumentSyncRegistrationOptions{
			NotebookDocumentSyncOptions: NotebookDocumentSyncOptions{
				NotebookSelector: []*NotebookSelectorItem{},
			},
			StaticRegistrationOptions: StaticRegistrationOptions{
				ID: optValRegIntermediate.ID,
			},
		}
		err = unmarshalNotebookDocumentSyncOptions(
			optValRegIntermediate.notebookDocumentSyncOptionsIntermediate,
			&targetRegOptVal.NotebookDocumentSyncOptions,
		)
		if err != nil {
			return err
		}
		serverCapabilities.NotebookDocumentSync = *targetRegOptVal
	} else {
		return err
	}

	return nil
}

func unmarshalNotebookDocumentSyncOptions(
	intermediate notebookDocumentSyncOptionsIntermediate,
	target *NotebookDocumentSyncOptions,
) error {
	for _, intermediateSelector := range intermediate.NotebookSelector {
		var filterVal NotebookDocumentFilter
		if err := json.Unmarshal(intermediateSelector.Notebook, &filterVal); err == nil {
			item := &NotebookSelectorItem{
				Notebook: filterVal,
				Cells:    intermediateSelector.Cells,
			}
			target.NotebookSelector = append(target.NotebookSelector, item)
		} else {
			var strVal string
			if err := json.Unmarshal(intermediateSelector.Notebook, &strVal); err == nil {
				item := &NotebookSelectorItem{
					Notebook: strVal,
					Cells:    intermediateSelector.Cells,
				}
				target.NotebookSelector = append(target.NotebookSelector, item)
			} else {
				return err
			}
		}
	}

	return nil
}
