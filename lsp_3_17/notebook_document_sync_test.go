package lsp

import (
	"encoding/json"
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
	tests := []struct {
		name          string
		input         string
		expected      *ServerCapabilities
		expectedError error
	}{
		{
			name: "unmarhsals notebook document registration options",
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
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			capabilities := &ServerCapabilities{}
			err := json.Unmarshal([]byte(test.input), capabilities)
			if test.expectedError != nil {
				s.Require().Error(err)
				s.Require().Equal(test.expectedError, err)
			} else {
				s.Require().NoError(err)
				s.Require().Equal(test.expected, capabilities)
			}
		})
	}
}

func TestNotebookDocumentSyncTestSuite(t *testing.T) {
	suite.Run(t, new(NotebookDocumentSyncTestSuite))
}
