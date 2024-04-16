package dto

import (
	"time"
)

type TaskRequest struct {
	ID       string `json:"-"`
	Title    string `json:"title" validate:"required,max=200"`
	ActiveAt string `json:"activeAt" validate:"required,date"`
}

// GetActiveAt ,,,
// Получение даты задачи в таймзоне
func (r *TaskRequest) GetActiveAt() time.Time {
	t, _ := time.Parse(time.DateOnly, r.ActiveAt)
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

type CreateResponse struct {
	ID string `json:"id"`
}

type ListRequest struct {
	Status string `form:"status" validate:"omitempty,task-status"`
}

type Task struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	ActiveAt string `json:"activeAt"`
}

type ListResponse []*Task
