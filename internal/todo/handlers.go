package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yayxela/go-todo/internal/dto"
	"github.com/yayxela/go-todo/internal/validate"
	"github.com/yayxela/go-todo/internal/values"
)

func RegisterHandlers(rg *gin.RouterGroup, todo TODO, validate validate.Validator) {
	group := rg.Group("/todo-list/tasks")
	group.POST("", Create(todo, validate))
	group.PUT("/:id", Update(todo, validate))
	group.DELETE("/:id", Delete(todo, validate))
	group.PUT("/:id/done", Done(todo, validate))
	group.GET("", List(todo, validate))
}

// Create 		...
// @Summary		Create
// @Description	Создание новой задачи
// @Tags 		tasks
// @Accept      json
// @Produce     json
// @Param       input     		body      	dto.CreateRequest		true    "Запрос на создание записи"
// @Success     201          	{object}	dto.CreateResponse		"Новая запись"
// @Failure     404
// @Router /api/v1/todo-list/tasks [post]
func Create(todo TODO, validator validate.Validator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := new(dto.CreateRequest)
		if err := ctx.Bind(request); err != nil {
			_ = ctx.Error(err)
			return
		}
		if err := validator.Validate(request); err != nil {
			_ = ctx.Error(err)
			return
		}
		res, err := todo.Create(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusCreated, res)
	}
}

// Update		...
// @Summary		Update
// @Description	Обновление уже существующей задачи
// @Tags 		tasks
// @Accept      json
// @Produce     json
// @Param 		id		path	string				true	"ID задачи"
// @Param       input	body	dto.UpdateRequest	true    "Запрос на обновление записи"
// @Success     204
// @Failure     404
// @Router /api/v1/todo-list/tasks/{id} [put]
func Update(todo TODO, validator validate.Validator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &dto.UpdateRequest{
			GetByID: dto.GetByID{
				ID: ctx.Param(values.ID),
			},
		}

		if err := ctx.Bind(request); err != nil {
			_ = ctx.Error(err)
			return
		}
		if err := validator.Validate(request); err != nil {
			_ = ctx.Error(err)
			return
		}
		err := todo.Update(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusNoContent, nil)
	}
}

// Delete		...
// @Summary		Delete
// @Description	Удаление задачи
// @Tags 		tasks
// @Accept      json
// @Produce     json
// @Param 		id		path	string	true	"ID задачи"
// @Success     204
// @Failure     404
// @Router /api/v1/todo-list/tasks/{id} [delete]
func Delete(todo TODO, validator validate.Validator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &dto.GetByID{
			ID: ctx.Param(values.ID),
		}
		if err := validator.Validate(request); err != nil {
			_ = ctx.Error(err)
			return
		}
		err := todo.Delete(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusNoContent, nil)
	}
}

// Done		...
// @Summary		Done
// @Description	Пометить задачу выполненной
// @Tags 		tasks
// @Accept      json
// @Produce     json
// @Param 		id		path	string	true	"Document  id"
// @Success     204
// @Failure     404
// @Router /api/v1/todo-list/tasks/{id}/done [put]
func Done(todo TODO, validator validate.Validator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &dto.GetByID{
			ID: ctx.Param(values.ID),
		}
		if err := validator.Validate(request); err != nil {
			_ = ctx.Error(err)
			return
		}
		err := todo.Done(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusNoContent, nil)
	}
}

// List		...
// @Summary		List
// @Description	Список задач по статусу
// @Tags 		tasks
// @Accept      json
// @Produce     json
// @Param 		input	query 		dto.ListRequest true "Запрос на список записей"
// @Success     200		{object}	dto.ListResponse	"Список записей"
// @Router /api/v1/todo-list/tasks [get]
func List(todo TODO, validator validate.Validator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := new(dto.ListRequest)
		if err := ctx.Bind(request); err != nil {
			_ = ctx.Error(err)
			return
		}
		if err := validator.Validate(request); err != nil {
			_ = ctx.Error(err)
			return
		}
		response, err := todo.List(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusOK, response)
	}
}
