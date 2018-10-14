package wei

import (
	"encoding/json"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// MyLogger Init
var (
	once      sync.Once
	ZapLogger *zap.Logger
)

// Logging file
func Logging(logType string, message string) error {

	var err error

	// once ensures the singleton is initialized only once
	once.Do(func() {
		var cfg zap.Config
		rawJSON := []byte(`{
			"level": "debug",
			"encoding": "json",
			"outputPaths": ["stdout", "./logs/logs"],
			"errorOutputPaths": ["stderr", "./logs/errors"],
			"encoderConfig": {
			  "messageKey": "message",
			  "levelKey": "level",
			  "levelEncoder": "lowercase"
			}
		  }`)

		if err = json.Unmarshal(rawJSON, &cfg); err != nil {
			panic(err)
		}

		cfg.EncoderConfig.TimeKey = "time"
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		newLogger, err := cfg.Build()
		if err != nil {
			panic(err)
		}
		ZapLogger = newLogger
	})
	defer ZapLogger.Sync()

	if logType == "info" {
		ZapLogger.Info(message)
	}

	if logType == "error" {
		ZapLogger.Error(message)
	}
	return err
}
