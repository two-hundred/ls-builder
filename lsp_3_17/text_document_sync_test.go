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
	tests := []struct {
		name     string
		input    string
		expected *ServerCapabilities
	}{
		{
			name:  "unmarhsals text document sync kind",
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

	for _, test := range tests {
		s.Run(test.name, func() {
			capabilities := &ServerCapabilities{}
			err := json.Unmarshal([]byte(test.input), capabilities)
			s.Require().NoError(err)
			s.Require().Equal(test.expected, capabilities)
		})
	}
}

func TestTextDocumentSyncTestSuite(t *testing.T) {
	suite.Run(t, new(TextDocumentSyncTestSuite))
}
