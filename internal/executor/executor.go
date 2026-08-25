package executor

import "task-forge/internal/models"

type Executor interface {
	Execute(task models.Task) error
}
