package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/yayxela/go-todo/internal/values"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Status    values.TaskStatus  `bson:"status"`
	ActiveAt  time.Time          `bson:"activeAt"`
	CreatedAt time.Time          `bson:"createdAt"`
}

// GetTitle ...
// Получение заголовка задачи
func (t *Task) GetTitle() string {
	title := t.Title
	if t.ActiveAt.Weekday() == time.Saturday ||
		t.ActiveAt.Weekday() == time.Sunday {
		title = fmt.Sprintf("%s - %s", values.Weekend, title)
	}
	return title
}
