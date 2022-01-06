package internal

import (
	"github.com/Kichiyaki/grpcplayground/internal"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, internal.Wrap(err, "zap.NewDevelopment")
	}

	zap.ReplaceGlobals(logger)

	return logger, nil
}
