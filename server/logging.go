package server

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type JSONRPCLogger struct {
	logger *zap.Logger
}

// Printf implements the jsonrpc2.Logger interface.
func (l *JSONRPCLogger) Printf(format string, v ...any) {
	l.logger.Debug(fmt.Sprintf(strings.TrimSuffix(format, "\n"), v...))
}
