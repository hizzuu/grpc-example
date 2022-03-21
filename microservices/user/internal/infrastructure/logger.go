package infrastructure

import "go.uber.org/zap"

type logger struct {
	Log *zap.SugaredLogger
}

func NewLogger() (*logger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &logger{Log: l.Sugar()}, nil
}
