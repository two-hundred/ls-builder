package server

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
	"github.com/two-hundred/ls-builder/common"
	"go.uber.org/zap"
)

// Server is an LSP over JSON-RPC 2.0 server
// that works for all versions of LSP that ls-builder
// provides.
type Server struct {
	handler      common.Handler
	debug        bool
	logger       *zap.Logger
	conn         *jsonrpc2.Conn
	timeout      time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
}

// ServerOption is a function that configures a server.
type ServerOption func(*Server)

// WithServerTimeout configures the server with a timeout.
func WithServerTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

// WithServerReadTimeout configures the server with a read timeout.
func WithServerReadTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.readTimeout = timeout
	}
}

// WithServerWriteTimeout configures the server with a write timeout.
func WithServerWriteTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.writeTimeout = timeout
	}
}

// NewServer creates a new LSP server over JSON-RPC 2.0.
func NewServer(
	handler common.Handler,
	debug bool,
	logger *zap.Logger,
	conn *jsonrpc2.Conn,
	opts ...ServerOption,
) *Server {
	server := &Server{
		handler:      handler,
		debug:        debug,
		logger:       logger,
		conn:         conn,
		timeout:      DefaultTimeout,
		readTimeout:  DefaultTimeout,
		writeTimeout: DefaultTimeout,
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

// GetTimeout returns the server's timeout.
func (s *Server) GetTimeout() time.Duration {
	return s.timeout
}

// GetReadTimeout returns the server's read timeout.
func (s *Server) GetReadTimeout() time.Duration {
	return s.readTimeout
}

// GetWriteTimeout returns the server's write timeout.
func (s *Server) GetWriteTimeout() time.Duration {
	return s.writeTimeout
}

// Serve serves a JSON-RPC 2.0 connection. If nil is passed in
// for the connection, the server will use the connection that
// was configured when the server was created.
func (s *Server) Serve(optConn *jsonrpc2.Conn, connLogger *zap.Logger) {
	if optConn == nil && s.conn == nil {
		s.logger.Fatal("no connection passed in or configured for server")
	}

	conn := optConn
	if conn == nil {
		conn = s.conn
	}

	connLogger.Info("new stream connection")
	<-conn.DisconnectNotify()
	connLogger.Info("stream connection closed")
}

// ServeWebSocket serves a JSON-RPC 2.0 connection over a WebSocket connection.
// See `RunWebSocketServer` for a complete example of how to serve JSON-RPC 2.0
// communication over WebSockets.
func (s *Server) ServeWebSocket(conn *websocket.Conn, logger *zap.Logger) {
	s.logger.Info("new web socket connection")
	<-NewWebSocketConnection(
		s.NewHandler(),
		conn,
		WithTimeout(s.timeout),
		WithReadTimeout(s.readTimeout),
		WithWriteTimeout(s.writeTimeout),
	).DisconnectNotify()
	s.logger.Info("web socket connection closed")
}
