package lsp

import (
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

func TestTextDocumentSyncTestSuite(t *testing.T) {
	suite.Run(t, new(TextDocumentSyncTestSuite))
}
