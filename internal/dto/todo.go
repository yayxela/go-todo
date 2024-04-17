package dto

type TaskRequest struct {
	Title    string `json:"title" validate:"required,max=200"`
	ActiveAt string `json:"activeAt" validate:"required,date"`
}

type CreateRequest struct {
	TaskRequest
}

type GetByID struct {
	ID string `name:"id" json:"-" validate:"required,object-id"`
}

type UpdateRequest struct {
	GetByID
	TaskRequest
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
