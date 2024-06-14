package lsp

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type NotebookDocumentSyncTestSuite struct {
	suite.Suite
}

func (s *NotebookDocumentSyncTestSuite) Test_unmarshal_notebook_document_sync_capability() {
	scheme := "file"
	pattern := "**/*.ipynb"
	testRegisterID := "test-register-id"
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals notebook document registration options",
			input: `
			{
				"notebookDocumentSync": {
					"id": "test-register-id",
					"notebookSelector": [
						{
							"notebook": {
								"notebookType": "test-notebook",
								"scheme": "file",
								"pattern": "**/*.ipynb"
							},
							"cells": [{ "language": "python" }]
						},
						{
							"notebook": "*",
							"cells": [{ "language": "python" }]
						}

					],
					"save": true
				}
			}
			`,
			expected: &ServerCapabilities{
				NotebookDocumentSync: NotebookDocumentSyncRegistrationOptions{
					StaticRegistrationOptions: StaticRegistrationOptions{
						ID: &testRegisterID,
					},
					NotebookDocumentSyncOptions: NotebookDocumentSyncOptions{
						NotebookSelector: []*NotebookSelectorItem{
							{
								Notebook: NotebookDocumentFilter{
									NotebookType: "test-notebook",
									Scheme:       &scheme,
									Pattern:      &pattern,
								},
								Cells: []NotebookCellLanguage{
									{
										Language: "python",
									},
								},
							},
							{
								Notebook: "*",
								Cells: []NotebookCellLanguage{
									{
										Language: "python",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func TestNotebookDocumentSyncTestSuite(t *testing.T) {
	suite.Run(t, new(NotebookDocumentSyncTestSuite))
}
