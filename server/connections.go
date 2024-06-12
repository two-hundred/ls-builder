package server

import (
	"context"
	"io"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
	wsjsonrpc2 "github.com/sourcegraph/jsonrpc2/websocket"
)

var DefaultTimeout = time.Minute

type connWrapper struct {
	timeout            time.Duration
	readTimeout        time.Duration
	writeTimeout       time.Duration
	streamTimeout      time.Duration
	webSocketTimeout   time.Duration
	wsConn             *websocket.Conn
	jsonRPCConnOptions []jsonrpc2.ConnOpt
}

type ConnOption func(*connWrapper)

func WithTimeout(timeout time.Duration) ConnOption {
	return func(c *connWrapper) {
		c.timeout = timeout
	}
}

func WithReadTimeout(timeout time.Duration) ConnOption {

	return func(c *connWrapper) {
		c.readTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) ConnOption {
	return func(c *connWrapper) {
		c.writeTimeout = timeout
	}
}

func WithStreamTimeout(timeout time.Duration) ConnOption {
	return func(c *connWrapper) {
		c.streamTimeout = timeout
	}
}

func WithWebSocketTimeout(timeout time.Duration) ConnOption {
	return func(c *connWrapper) {
		c.webSocketTimeout = timeout
	}
}

func WithWebSocket(wsConn *websocket.Conn) ConnOption {
	return func(c *connWrapper) {
		c.wsConn = wsConn
	}
}

func WithJSONRPCConnOptions(opts ...jsonrpc2.ConnOpt) ConnOption {
	return func(c *connWrapper) {
		c.jsonRPCConnOptions = opts
	}
}

func NewStreamConnection(handler jsonrpc2.Handler, stream io.ReadWriteCloser, opts ...ConnOption) *jsonrpc2.Conn {
	c := &connWrapper{
		timeout:       DefaultTimeout,
		readTimeout:   DefaultTimeout,
		writeTimeout:  DefaultTimeout,
		streamTimeout: DefaultTimeout,
	}

	for _, opt := range opts {
		opt(c)
	}

	// Context with deadline for establishing the connection.
	ctx, cancel := context.WithTimeout(context.Background(), c.streamTimeout)
	defer cancel()

	return jsonrpc2.NewConn(
		ctx,
		jsonrpc2.NewBufferedStream(stream, jsonrpc2.VSCodeObjectCodec{}),
		handler,
		c.jsonRPCConnOptions...,
	)
}

func NewWebSocketConnection(handler jsonrpc2.Handler, wsConn *websocket.Conn, opts ...ConnOption) *jsonrpc2.Conn {
	c := &connWrapper{
		timeout:       DefaultTimeout,
		readTimeout:   DefaultTimeout,
		writeTimeout:  DefaultTimeout,
		streamTimeout: DefaultTimeout,
	}

	for _, opt := range opts {
		opt(c)
	}

	// Context with deadline for establishing the connection.
	ctx, cancel := context.WithTimeout(context.Background(), c.webSocketTimeout)
	defer cancel()

	return jsonrpc2.NewConn(
		ctx,
		wsjsonrpc2.NewObjectStream(wsConn),
		handler,
		c.jsonRPCConnOptions...,
	)
}
