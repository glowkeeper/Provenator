package log

import (
    "log"
    "go.uber.org/zap"
    "github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/internal/text"
)

// SLogger - faster but slower
var SLogger *zap.SugaredLogger
// PLogger - faster and faster
var PLogger *zap.Logger

func setLoggers (logger *zap.Logger) {

	defer logger.Sync()
	SLogger = logger.Sugar()
	PLogger = SLogger.Desugar()
}

// Init initialized the logging tool.
func Init(env string) {

	if env == "production" {

		logger, err := zap.NewProduction()
		if err != nil {
		  log.Fatalf(text.ErrorZap, err)
		}
		setLoggers(logger)

	} else {

		logger, err := zap.NewDevelopment()
		if err != nil {
		  log.Fatal(text.ErrorZap, err)
		}
		setLoggers(logger)
	}
}
