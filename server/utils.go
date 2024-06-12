package server

import (
	"fmt"

	"go.uber.org/zap"
)

func callAndLog(f func() error, name string, logger *zap.Logger) {
	err := f()
	if err != nil {
		logger.Error(fmt.Sprintf("error calling %s: %s", name, err.Error()))
	}
}
