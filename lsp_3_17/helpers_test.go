package lsp

import (
	"encoding/json"

	"github.com/stretchr/testify/suite"
)

type serverCapabilityFixture struct {
	name     string
	input    string
	expected *ServerCapabilities
}

func testServerCapabilities(s *suite.Suite, tests []serverCapabilityFixture) {
	for _, test := range tests {
		s.Run(test.name, func() {
			capabilities := &ServerCapabilities{}
			err := json.Unmarshal([]byte(test.input), capabilities)
			s.Require().NoError(err)
			s.Require().Equal(test.expected, capabilities)
		})
	}
}

type clientCapabilityFixture struct {
	name     string
	input    string
	expected *ClientCapabilities
}

func testClientCapabilities(s *suite.Suite, tests []clientCapabilityFixture) {
	for _, test := range tests {
		s.Run(test.name, func() {
			capabilities := &ClientCapabilities{}
			err := json.Unmarshal([]byte(test.input), capabilities)
			s.Require().NoError(err)
			s.Require().Equal(test.expected, capabilities)
		})
	}
}
