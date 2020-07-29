package server

import (
	"log"
	"net/http"
	"packages/config"
	"packages/database"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Server struct {
	ServerConfig   *config.ServerConfig
	DataBaseConfig *config.DataBaseConfig
	Logger         *zap.Logger
	Router         *mux.Router
	ConnectionDB   *database.Connection
}

func NewServer(serverConfig *config.ServerConfig, dbConfig *config.DataBaseConfig) *Server {
	return &Server{
		ServerConfig:   serverConfig,
		DataBaseConfig: dbConfig,
		Logger:         NewLogger(serverConfig.LogLevel),
		Router:         mux.NewRouter(),
		ConnectionDB:   database.NewConnectionDB(dbConfig.Name, dbConfig.User, dbConfig.Password),
	}
}

func (s *Server) StartServer() error {
	s.NewAuthController()
	s.Logger.Info("Server started",
		zap.String("Name", "OAuth2"),
		zap.String("Port", s.ServerConfig.Port),
	)
	return http.ListenAndServe(s.ServerConfig.Port, s.Router)
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
