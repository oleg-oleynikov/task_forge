package server

import "task-forge/internal/config"

type server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *server {
	return &server{
		cfg: cfg,
	}
}

func (s *server) Run() error {
	return nil
}
