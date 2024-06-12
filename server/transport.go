package server

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"go.uber.org/zap"
)

func newNetworkListener(network string, address string, logger *zap.Logger) (*net.Listener, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		logger.Error(fmt.Sprintf("could not bind to address %s: %v", address, err))
		return nil, err
	}

	cert := os.Getenv("TLS_CERT")
	key := os.Getenv("TLS_KEY")
	if (cert != "") && (key != "") {
		cert, err := tls.X509KeyPair([]byte(cert), []byte(key))
		if err != nil {
			return nil, err
		}
		listener = tls.NewListener(listener, &tls.Config{
			Certificates: []tls.Certificate{cert},
		})
	}

	return &listener, nil
}
