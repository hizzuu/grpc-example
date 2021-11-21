package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func NewLogger() error {
	log, err := zap.NewProduction()
	if err != nil {
		return err
	}
	Log = log.Sugar()
	return nil
}
