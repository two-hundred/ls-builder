package server

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// RunWebSocketServer starts a new web socket server on the provided address.
func RunWebSocketServer(address string, server *Server, logger *zap.Logger, httpServer *http.Server) error {
	mux := http.NewServeMux()
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	var connectionCount uint64

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			logger.Error(fmt.Sprintf("error upgrading HTTP to web socket: %s", err.Error()))
			return
		}

		connLogger := logger.With(zap.Uint64("id", atomic.AddUint64(&connectionCount, 1)))
		defer callAndLog(conn.Close, "wsConn.Close", connLogger)
		server.ServeWebSocket(conn, connLogger)
	})

	listener, err := newNetworkListener("tcp", address, logger)
	if err != nil {
		return err
	}

	if httpServer == nil {
		httpServer = &http.Server{}
	}
	httpServer.Handler = mux

	if httpServer.ReadTimeout == 0 {
		httpServer.ReadTimeout = server.GetReadTimeout()
	}

	if httpServer.WriteTimeout == 0 {
		httpServer.WriteTimeout = server.GetWriteTimeout()
	}

	err = httpServer.Serve(*listener)
	return errors.Wrap(err, "WebSocket")
}
