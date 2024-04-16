package todo

import (
	"context"
	"time"

	"github.com/yayxela/go-todo/internal/db"
	"github.com/yayxela/go-todo/internal/db/models"
	taskRepo "github.com/yayxela/go-todo/internal/db/repository/task"
	"github.com/yayxela/go-todo/internal/dto"
	"github.com/yayxela/go-todo/internal/values"
)

type TODO interface {
	Create(ctx context.Context, request *dto.TaskRequest) (*dto.CreateResponse, error)
	Update(ctx context.Context, request *dto.TaskRequest) error
	Delete(ctx context.Context, id string) error
	Done(ctx context.Context, id string) error
	List(ctx context.Context, request *dto.ListRequest) (dto.ListResponse, error)
}

type service struct {
	taskRepo taskRepo.Task
}

func New(idb db.IDB) TODO {
	return &service{
		taskRepo: taskRepo.New(idb),
	}
}

func (s *service) Create(ctx context.Context, request *dto.TaskRequest) (*dto.CreateResponse, error) {
	task := &models.Task{
		Title:    request.Title,
		ActiveAt: request.GetActiveAt(),
	}
	err := s.taskRepo.Create(ctx, task)
	if err != nil {
		return nil, err
	}
	return &dto.CreateResponse{
		ID: task.ID.Hex(),
	}, nil
}

func (s *service) Update(ctx context.Context, request *dto.TaskRequest) error {
	task := &models.Task{
		ID:       db.GetOID(request.ID),
		Title:    request.Title,
		ActiveAt: request.GetActiveAt(),
	}
	return s.taskRepo.Update(ctx, task)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.taskRepo.Delete(ctx, id)
}

func (s *service) Done(ctx context.Context, id string) error {
	task := &models.Task{
		ID:     db.GetOID(id),
		Status: values.Done,
	}
	return s.taskRepo.Update(ctx, task)
}

func (s *service) List(ctx context.Context, request *dto.ListRequest) (dto.ListResponse, error) {
	list, err := s.taskRepo.List(ctx, request)
	if err != nil {
		return nil, err
	}
	tasks := make([]*dto.Task, 0, len(list))
	for _, task := range list {
		tasks = append(tasks, &dto.Task{
			ID:       task.ID.Hex(),
			Title:    task.GetTitle(),
			ActiveAt: task.ActiveAt.In(time.Local).Format(time.DateOnly),
		})
	}
	return tasks, nil
}