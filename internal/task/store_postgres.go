package task

import (
	"context"
	"database/sql"
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
	return nil, nil
}
func Complete(ctx context.Context, taskID string) error {
	return nil
}

func Create(ctx context.Context, task models.Task) error {
	return nil
}

func Heartbeat(ctx context.Context, taskID string, workerID string) error {
	return nil
}

func Fail(ctx context.Context, task *models.Task, execErr error) error {
	return nil
}
