package server

import (
	"context"
	"io"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
	wsjsonrpc2 "github.com/sourcegraph/jsonrpc2/websocket"
)

// DefaultTimeout is the default timeout used for a JSON-RPC 2.0 connection.
var DefaultTimeout = time.Minute

type connWrapper struct {
	timeout            time.Duration
	readTimeout        time.Duration
	writeTimeout       time.Duration
	wsConn             *websocket.Conn
	jsonRPCConnOptions []jsonrpc2.ConnOpt
}

// ConnOption is a function that configures a connection.
type ConnOption func(*connWrapper)

// WithTimeout configures the connection with a timeout.
func WithTimeout(timeout time.Duration) ConnOption {
	return func(c *connWrapper) {
		c.timeout = timeout
	}
}

// WithReadTimeout configures the connection with a read timeout.
func WithReadTimeout(timeout time.Duration) ConnOption {

	return func(c *connWrapper) {
		c.readTimeout = timeout
	}
}

// WithWriteTimeout configures the connection with a write timeout.
func WithWriteTimeout(timeout time.Duration) ConnOption {
	return func(c *connWrapper) {
		c.writeTimeout = timeout
	}
}

// WithWebSocket configures the connection with a WebSocket connection.
func WithWebSocket(wsConn *websocket.Conn) ConnOption {
	return func(c *connWrapper) {
		c.wsConn = wsConn
	}
}

// WithJSONRPCConnOptions configures the connection with JSON-RPC 2.0 connection options.
func WithJSONRPCConnOptions(opts ...jsonrpc2.ConnOpt) ConnOption {
	return func(c *connWrapper) {
		c.jsonRPCConnOptions = opts
	}
}

// NewStreamConnection creates a new JSON-RPC 2.0 connection over a stream (io.ReadWriterCloser).
func NewStreamConnection(handler jsonrpc2.Handler, stream io.ReadWriteCloser, opts ...ConnOption) *jsonrpc2.Conn {
	c := &connWrapper{
		timeout:      DefaultTimeout,
		readTimeout:  DefaultTimeout,
		writeTimeout: DefaultTimeout,
	}

	for _, opt := range opts {
		opt(c)
	}

	ctx := context.Background()

	return jsonrpc2.NewConn(
		ctx,
		jsonrpc2.NewBufferedStream(stream, jsonrpc2.VSCodeObjectCodec{}),
		handler,
		c.jsonRPCConnOptions...,
	)
}

// NewWebSocketConnection creates a new JSON-RPC 2.0 connection over a WebSocket connection.
func NewWebSocketConnection(handler jsonrpc2.Handler, wsConn *websocket.Conn, opts ...ConnOption) *jsonrpc2.Conn {
	c := &connWrapper{
		timeout:      DefaultTimeout,
		readTimeout:  DefaultTimeout,
		writeTimeout: DefaultTimeout,
	}

	for _, opt := range opts {
		opt(c)
	}

	ctx := context.Background()

	return jsonrpc2.NewConn(
		ctx,
		wsjsonrpc2.NewObjectStream(wsConn),
		handler,
		c.jsonRPCConnOptions...,
	)
}
