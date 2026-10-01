package task

import (
	"context"
	"database/sql"
	"fmt"
	"task-forge/internal/models"
	"time"
)

type PostgresTaskStore struct {
	db    *sql.DB
	lease time.Duration
}

func NewPostgresTaskStore(dsn string, lease time.Duration) (*PostgresTaskStore, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, ErrOpenDBConnection
	}

	if err := db.Ping(); err != nil {
		return nil, ErrPingDBConnection
	}

	return &PostgresTaskStore{
		db:    db,
		lease: lease,
	}, nil
}

func Claim(ctx context.Context, workerID, queueName string) (*models.Task, error) {
	return nil, fmt.Errorf("Not implemented")
}

func Complete(ctx context.Context, taskID string) error {
	return fmt.Errorf("Not implemented")
}

func Create(ctx context.Context, task models.Task) error {
	return fmt.Errorf("Not implemented")
}

func Heartbeat(ctx context.Context, taskID string, workerID string) error {
	return fmt.Errorf("Not implemented")
}

func Fail(ctx context.Context, task *models.Task, execErr error) error {
	return fmt.Errorf("Not implemented")
}
