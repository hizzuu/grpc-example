package logger

import "go.uber.org/zap"

var Log *zap.SugaredLogger

func NewLogger() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	Log = logger.Sugar()
	return nil
}
