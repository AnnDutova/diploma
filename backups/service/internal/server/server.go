package server

import (
	"backups/internal/backup"
	"backups/internal/config"
	"backups/internal/db"
	"backups/internal/minio"
	"backups/pkg/model"
	"log"
)

type Server struct {
	Config *model.Config
	Router *Router
}

func NewServer(configPath string) (*Server, error) {
	appConfig, err := config.ParseConfig(configPath)
	if err != nil {
		return nil, err
	}

	provider, err := db.NewDataBaseProvider(appConfig)
	if err != nil {
		return nil, err
	}
	log.Println("Init DB")

	minio, err := minio.NewMinioClient(appConfig)
	if err != nil {
		return nil, err
	}

	log.Println("Init minio")

	backup, err := backup.NewBackupManager(provider, minio)
	if err != nil {
		return nil, err
	}

	log.Println("Init backupManager")

	handler, err := NewHandler(backup)
	if err != nil {
		return nil, err
	}

	router, err := NewRouter(handler)
	if err != nil {
		return nil, err
	}

	return &Server{
		Config: appConfig,
		Router: router,
	}, nil
}

func (s *Server) Start() error {
	log.Println("Start Server")
	if err := s.Router.configurateRouter(s.Config.Server.Port); err != nil {
		return err
	}
	return nil
}
