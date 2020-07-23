package server

import (
	"log"
	"packages/config"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Server struct {
	ServerConfig   *config.ServerConfig
	DataBaseConfig *config.DataBaseConfig
	Logger         *zap.Logger
	Router         *mux.Router
}

func NewServer(serverConfig *config.ServerConfig, dbConfig *config.DataBaseConfig) *Server {
	return &Server{
		ServerConfig:   serverConfig,
		DataBaseConfig: dbConfig,
		Logger:         NewLogger(serverConfig.LogLevel),
	}
}

func (s *Server) StartServer() error {
	return nil
}

func NewLogger(logLevel string) *zap.Logger {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:     logLevel,
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.FullCallerEncoder,
		},
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Println(err)
	}
	return logger
}
