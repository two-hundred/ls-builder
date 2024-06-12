package server

import (
	"github.com/sourcegraph/jsonrpc2"
	"go.uber.org/zap"
)

func CreateConnectionOptions(debug bool, logger *zap.Logger) []jsonrpc2.ConnOpt {
	if debug {
		return []jsonrpc2.ConnOpt{
			jsonrpc2.LogMessages(&JSONRPCLogger{logger}),
		}
	}
	return nil
}
