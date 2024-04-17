package todo

import (
	"context"
	"time"

	"github.com/yayxela/go-todo/internal/db"
	"github.com/yayxela/go-todo/internal/db/models"
	taskRepo "github.com/yayxela/go-todo/internal/db/repository/task"
	"github.com/yayxela/go-todo/internal/dto"
	"github.com/yayxela/go-todo/internal/utils"
	"github.com/yayxela/go-todo/internal/values"
)

type TODO interface {
	Create(ctx context.Context, request *dto.CreateRequest) (*dto.CreateResponse, error)
	Update(ctx context.Context, request *dto.UpdateRequest) error
	Delete(ctx context.Context, request *dto.GetByID) error
	Done(ctx context.Context, request *dto.GetByID) error
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

func (s *service) Create(ctx context.Context, request *dto.CreateRequest) (*dto.CreateResponse, error) {
	task := &models.Task{
		Title:    request.Title,
		ActiveAt: utils.GetActiveAt(request.ActiveAt),
	}
	err := s.taskRepo.Create(ctx, task)
	if err != nil {
		return nil, err
	}
	return &dto.CreateResponse{
		ID: task.ID.Hex(),
	}, nil
}

func (s *service) Update(ctx context.Context, request *dto.UpdateRequest) error {
	task := &models.Task{
		ID:       utils.GetOID(request.ID),
		Title:    request.Title,
		ActiveAt: utils.GetActiveAt(request.ActiveAt),
	}
	return s.taskRepo.Update(ctx, task)
}

func (s *service) Delete(ctx context.Context, request *dto.GetByID) error {
	return s.taskRepo.Delete(ctx, request.ID)
}

func (s *service) Done(ctx context.Context, request *dto.GetByID) error {
	task := &models.Task{
		ID:     utils.GetOID(request.ID),
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
