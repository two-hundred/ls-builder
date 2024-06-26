package server

import (
	"context"

	"go.uber.org/zap"
)

// RunTCP begins listening for TCP connections on the provided address,
// creating a LSP over JSON-RPC 2.0 connection on top of each incoming
// TCP connection with the provided server.
func RunTCP(ctx context.Context, address string, server *Server, logger *zap.Logger) error {
	listener, err := newNetworkListener("tcp", address, logger)
	if err != nil {
		return err
	}

	connLogger := logger.With(zap.String("address", address))
	defer callAndLog((*listener).Close, "listener.Close", connLogger)
	logger.Info("listening for TCP connections")

	var connectionCount uint64

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			connection, err := (*listener).Accept()
			if err != nil {
				return err
			}

			connectionCount += 1
			connectionLogger := logger.With(zap.Uint64("id", connectionCount))

			go server.Serve(NewStreamConnection(server.NewHandler(), connection), connectionLogger)
		}
	}
}
