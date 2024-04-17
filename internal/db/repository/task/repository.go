package task

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/yayxela/go-todo/internal/db"
	"github.com/yayxela/go-todo/internal/db/models"
	"github.com/yayxela/go-todo/internal/dto"
	"github.com/yayxela/go-todo/internal/utils"
	"github.com/yayxela/go-todo/internal/values"
)

type Task interface {
	Create(ctx context.Context, model *models.Task) error
	Update(ctx context.Context, model *models.Task) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, req *dto.ListRequest) ([]*models.Task, error)
}

type task struct {
	collection *mongo.Collection
}

func New(idb db.IDB) Task {
	return &task{
		collection: idb.GetDB().Collection(db.TaskCollection),
	}
}

func (r *task) Create(ctx context.Context, model *models.Task) error {
	filter := bson.D{
		{Key: "title", Value: model.Title},
		{Key: "activeAt", Value: primitive.NewDateTimeFromTime(model.ActiveAt)},
	}
	err := r.collection.FindOne(ctx, filter).Decode(bson.M{})
	if !errors.Is(err, mongo.ErrNoDocuments) {
		if err != nil {
			return err
		}
		return values.ExistsError
	}

	model.CreatedAt = time.Now()
	model.Status = values.Active
	res, err := r.collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}
	model.ID, _ = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *task) Update(ctx context.Context, model *models.Task) error {
	filter := bson.D{{Key: "_id", Value: model.ID}}
	var update bson.D
	if model.Title != "" {
		update = append(update, bson.E{Key: "title", Value: model.Title})
	}
	if !model.ActiveAt.IsZero() {
		update = append(update, bson.E{Key: "activeAt", Value: model.ActiveAt})
	}
	if model.Status != values.None {
		update = append(update, bson.E{Key: "status", Value: model.Status})
	}
	err := r.collection.FindOneAndUpdate(ctx, filter, bson.D{{Key: "$set", Value: update}}).Err()
	return err
}

func (r *task) Delete(ctx context.Context, id string) error {
	filter := bson.D{{Key: "_id", Value: utils.GetOID(id)}}
	return r.collection.FindOneAndDelete(ctx, filter).Err()
}

func (r *task) List(ctx context.Context, req *dto.ListRequest) ([]*models.Task, error) {
	filter := make(bson.M)
	if req.Status != "" {
		filter["status"] = values.TaskStatus(req.Status)
	} else {
		filter["status"] = values.Active
	}
	if filter["status"] == values.Active {
		filter["activeAt"] = bson.M{"$lte": primitive.NewDateTimeFromTime(time.Now())}
	}
	opts := []*options.FindOptions{
		options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}),
	}
	cursor, err := r.collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	var list []*models.Task
	if err = cursor.All(ctx, &list); err != nil {
		return nil, err
	}
	return list, nil
}
