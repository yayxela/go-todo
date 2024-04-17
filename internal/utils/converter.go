package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// GetActiveAt ,,,
// Получение даты задачи в таймзоне
func GetActiveAt(activeAt string) time.Time {
	t, _ := time.Parse(time.DateOnly, activeAt)
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

// GetOID ...
// Получение primitive.ObjectID из строки
func GetOID(id string) primitive.ObjectID {
	oid, _ := primitive.ObjectIDFromHex(id)
	return oid
}
