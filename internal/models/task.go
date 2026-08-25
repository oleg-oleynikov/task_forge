package models

import "uuid"

type TaskStatus string

const (
	Archived TaskStatus = "archived"
	Pending  TaskStatus = "pending"
	Retry    TaskStatus = "retry"
)

type Task struct {
	taskId   string
	workerId string
	status   TaskStatus
	payload  []byte
}

func NewTask(payload []byte) *Task {
	return &Task{
		taskId:  uuid.NewV7().String(),
		status:  Pending,
		payload: payload,
	}
}
