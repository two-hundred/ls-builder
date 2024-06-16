package lsp

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TextDocumentSyncTestSuite struct {
	suite.Suite
}

func (s *TextDocumentSyncTestSuite) Test_unmarshal_text_document_sync_capability() {
	trueVal := true
	tests := []serverCapabilityFixture{
		{
			name:  "unmarshals text document sync kind",
			input: "{\"textDocumentSync\":1}",
			expected: &ServerCapabilities{
				TextDocumentSync: TextDocumentSyncKindFull,
			},
		},
		{
			name:  "unmarshals text document sync options",
			input: "{\"textDocumentSync\":{\"openClose\":true,\"change\":2}}",
			expected: &ServerCapabilities{
				TextDocumentSync: TextDocumentSyncOptions{
					OpenClose: &trueVal,
					Change:    &TextDocumentSyncKindIncremental,
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

func (s *TextDocumentSyncTestSuite) Test_unmarshal_did_change_text_document_params() {
	tests := []didChangeTextDocumentFixture{
		{
			name: "unmarshals did change text document params for whole document",
			input: `
			{
				"textDocument": {
					"uri": "file:///path/to/file",
					"version": 1
				},
				"contentChanges": [
					{
						"text": "new content"
					}
				]
			}
			`,
			expected: &DidChangeTextDocumentParams{
				TextDocument: VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: TextDocumentIdentifier{
						URI: "file:///path/to/file",
					},
					Version: 1,
				},
				ContentChanges: []interface{}{
					TextDocumentContentChangeEventWhole{
						Text: "new content",
					},
				},
			},
		},
		{
			name: "unmarshals did change text document params for range change event",
			input: `
			{
				"textDocument": {
					"uri": "file:///path/to/file2",
					"version": 1
				},
				"contentChanges": [
					{
						"text": "new content",
						"range": {
							"start": {
								"line": 3,
								"character": 2
							},
							"end": {
								"line": 3,
								"character": 14
							}
						}
					}
				]
			}
			`,
			expected: &DidChangeTextDocumentParams{
				TextDocument: VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: TextDocumentIdentifier{
						URI: "file:///path/to/file2",
					},
					Version: 1,
				},
				ContentChanges: []interface{}{
					TextDocumentContentChangeEvent{
						Range: &Range{
							Start: Position{
								Line:      3,
								Character: 2,
							},
							End: Position{
								Line:      3,
								Character: 14,
							},
						},
						Text: "new content",
					},
				},
			},
		},
	}

	testDidChangeTextDocument(&s.Suite, tests)
}

func TestTextDocumentSyncTestSuite(t *testing.T) {
	suite.Run(t, new(TextDocumentSyncTestSuite))
}

type didChangeTextDocumentFixture struct {
	name     string
	input    string
	expected *DidChangeTextDocumentParams
}

func testDidChangeTextDocument(s *suite.Suite, tests []didChangeTextDocumentFixture) {
	for _, test := range tests {
		s.Run(test.name, func() {
			params := &DidChangeTextDocumentParams{}
			err := json.Unmarshal([]byte(test.input), params)
			s.Require().NoError(err)
			s.Require().Equal(test.expected, params)
		})
	}
}
