package lsp

import (
	"context"
	"encoding/json"
	"io"
	"sync"

	"github.com/sourcegraph/jsonrpc2"
	"github.com/stretchr/testify/suite"
	"github.com/two-hundred/ls-builder/server"
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

type testConnectionsContainer struct {
	clientReceivedMessages []*json.RawMessage
	clientReceivedMethods  []string
	clientConn             *jsonrpc2.Conn
	serverConn             *jsonrpc2.Conn
	mu                     sync.Mutex
}

type testStream struct {
	in  io.Reader
	out io.Writer
}

// Read reads from the "in" stream (emulating stdin).
// Fulfils the io.Reader interface.
func (s *testStream) Read(p []byte) (int, error) {
	return s.in.Read(p)
}

// Write writes to the "out" stream (emulating stdout).
// Fulfils the io.Writer interface.
func (s *testStream) Write(p []byte) (int, error) {
	return s.out.Write(p)
}

// Close closes "in" and "out" streams.
// Fulfils the io.Closer interface.
func (s *testStream) Close() error {
	return nil
}

func createTestConnectionsContainer(serverHandler jsonrpc2.Handler) *testConnectionsContainer {
	// Wire up the client and server streams
	// to emulate communication over stdin and stdout.
	clientIn, serverOut := io.Pipe()
	serverIn, clientOut := io.Pipe()
	clientStream := &testStream{
		in:  clientIn,
		out: clientOut,
	}
	serverStream := &testStream{
		in:  serverIn,
		out: serverOut,
	}

	container := &testConnectionsContainer{
		clientReceivedMessages: []*json.RawMessage{},
		clientReceivedMethods:  []string{},
	}

	clientHandler := jsonrpc2.HandlerWithError(
		func(
			ctx context.Context,
			conn *jsonrpc2.Conn,
			req *jsonrpc2.Request,
		) (interface{}, error) {
			container.mu.Lock()
			defer container.mu.Unlock()
			container.clientReceivedMessages = append(container.clientReceivedMessages, req.Params)
			container.clientReceivedMethods = append(container.clientReceivedMethods, req.Method)
			return nil, nil
		},
	)
	serverConn := server.NewStreamConnection(serverHandler, serverStream)
	clientConn := server.NewStreamConnection(clientHandler, clientStream)
	container.serverConn = serverConn
	container.clientConn = clientConn
	return container
}

func newTestServerHandler() jsonrpc2.Handler {
	return jsonrpc2.HandlerWithError(
		func(
			ctx context.Context,
			conn *jsonrpc2.Conn,
			req *jsonrpc2.Request,
		) (interface{}, error) {
			return nil, nil
		},
	)
}
