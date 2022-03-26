package logs

import "go.uber.org/zap"

var Logger *zap.Logger
var Sugar *zap.SugaredLogger

func Init() {
	Logger, _ = zap.NewProduction()
	Sugar = Logger.Sugar()
}
