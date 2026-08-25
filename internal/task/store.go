package task

import (
	"context"
	"task-forge/internal/models"
)

var (
	ErrOpenDBConnection error = models.NewInternalError(models.DBError, "failed to open db connection")
	ErrPingDBConnection error = models.NewInternalError(models.DBError, "error by ping db")
)

type TaskStore interface {
	Claim(ctx context.Context, workerID, queueName string) (*models.Task, error)
	Complete(ctx context.Context, taskID string) error
	Create(ctx context.Context, task models.Task) error
	Heartbeat(ctx context.Context, taskID string, workerID string) error
	Fail(ctx context.Context, task *models.Task, execErr error) error
}
