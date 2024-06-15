package server

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
	"github.com/two-hundred/ls-builder/common"
	"go.uber.org/zap"
)

type Server struct {
	handler      common.Handler
	debug        bool
	logger       *zap.Logger
	conn         *jsonrpc2.Conn
	timeout      time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
}

type ServerOption func(*Server)

func WithServerTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithServerReadTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.readTimeout = timeout
	}
}

func WithServerWriteTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.writeTimeout = timeout
	}
}

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

func (s *Server) GetTimeout() time.Duration {
	return s.timeout
}

func (s *Server) GetReadTimeout() time.Duration {
	return s.readTimeout
}

func (s *Server) GetWriteTimeout() time.Duration {
	return s.writeTimeout
}

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
