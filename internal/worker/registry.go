package worker

import (
	"sync"
	"task-forge/internal/executor"
	"task-forge/internal/models"
)

type ExecutorType string

var (
	ErrExecutorAlreadyExists error = models.NewInternalError(models.RegistryError, "executor already exist")
)

const (
	Shell       ExecutorType = "shell"
	HTTPRequest ExecutorType = "http_request"
)

type ExecutorRegistry struct {
	mu  sync.RWMutex
	reg map[ExecutorType]executor.Executor
}

func NewExecutorRegistry() *ExecutorRegistry {
	return &ExecutorRegistry{
		reg: make(map[ExecutorType]executor.Executor),
	}
}

func (r *ExecutorRegistry) Get(t ExecutorType) (executor.Executor, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	executor, exist := r.reg[t]
	return executor, exist
}

func (r *ExecutorRegistry) Add(t ExecutorType, exc executor.Executor) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.Get(t); ok {
		return ErrExecutorAlreadyExists
	}

	r.reg[t] = exc
	return nil
}
