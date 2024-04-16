package values

type TaskStatus string

// статусы задачи
const (
	Active TaskStatus = "active"
	Done   TaskStatus = "done"
	None   TaskStatus = ""
)

var (
	TaskStatuses = []TaskStatus{Active, Done}
)

// ключи
const (
	ID      = "id"
	Weekend = "ВЫХОДНОЙ"
)
