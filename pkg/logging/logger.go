package logging

import (
	"encoding/json"
	"go.uber.org/zap"
)

func GetLogger(applicationName string) (*zap.SugaredLogger, error) {
	rawJSON := []byte(`{
   "level": "debug",
   "encoding": "json",
   "outputPaths": ["stdout"],
   "errorOutputPaths": ["stderr"],
   "encoderConfig": {
     "messageKey": "message",
     "levelKey": "level",
     "levelEncoder": "lowercase"
   }
 }`)
	var cfg zap.Config

	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		return nil, err
	}

	zapLogger, err := cfg.Build()

	if err != nil {
		return nil, err
	}

	zapLogger = zapLogger.Named(applicationName)

	return zapLogger.Sugar(), nil
}
