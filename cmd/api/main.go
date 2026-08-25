package api

import (
	"log/slog"
	"task-forge/internal/config"
	"task-forge/internal/server"
)

func main() {
	cfg := &config.Config{}
	s := server.NewServer(cfg)
	if err := s.Run(); err != nil {
		slog.Error("server down")
	}
}
