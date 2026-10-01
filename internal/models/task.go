package models

import (
	"time"
	"uuid"
)

type TaskStatus string

const (
	Archived TaskStatus = "archived"
	Pending  TaskStatus = "pending"
	Running  TaskStatus = "running"
	Retry    TaskStatus = "retry"
)

type Task struct {
	taskId    string
	workerId  string
	lockedBy  string // workerId
	lockedAt  time.Time
	status    TaskStatus
	payload   []byte
	runAt     time.Time
	createdAt time.Time
	updatedAt time.Time
	attemps   int
}

func NewTask(payload []byte) *Task {
	return &Task{
		taskId:    uuid.NewV7().String(),
		status:    Pending,
		payload:   payload,
		createdAt: time.Now(),
	}
}
