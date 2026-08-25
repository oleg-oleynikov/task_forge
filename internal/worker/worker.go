package worker

import (
	"context"
	"fmt"
	"sync"
	"task-forge/internal/task"
	"time"
)

type worker struct {
	id       string
	store    task.TaskStore
	registry *ExecutorRegistry
	lease    time.Duration
}

func New(id string, store task.TaskStore, registry *ExecutorRegistry) *worker {
	return &worker{
		id:       id,
		store:    store,
		registry: registry,
		lease:    5 * time.Minute, // TODO: получать lease извне
	}
}

func (w *worker) Run(ctx context.Context, concurrency int) {
	var wg sync.WaitGroup
	for i := range concurrency {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			w.loop(ctx, fmt.Sprintf("%s-%d", w.id, num))
		}(i)
	}
	wg.Wait()
}

func (w *worker) loop(ctx context.Context, workerID string) {
	ticker := time.NewTicker(time.Second * 2) // Поменять на duration извне
	return
}

func process(ctx context.Context) error {
	return nil
}

// func
