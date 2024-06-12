package server

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func RunWebSocketServer(address string, server *Server, logger *zap.Logger) error {
	mux := http.NewServeMux()
	upgrader := websocket.Upgrader{CheckOrigin: func(request *http.Request) bool { return true }}

	var connectionCount uint64

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			logger.Warn(fmt.Sprintf("error upgrading HTTP to web socket: %s", err.Error()))
			http.Error(writer, errors.Wrap(err, "could not upgrade to web socket").Error(), http.StatusBadRequest)
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

	httpServer := http.Server{
		Handler:      http.TimeoutHandler(mux, server.GetTimeout(), ""),
		ReadTimeout:  server.GetReadTimeout(),
		WriteTimeout: server.GetWriteTimeout(),
	}

	err = httpServer.Serve(*listener)
	return errors.Wrap(err, "WebSocket")
}
