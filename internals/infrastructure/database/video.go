package database

import (
	"context"
	"errors"
	"github.com/koba1108/go-mongodb/internals/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type videoRepository struct {
	db  *mongo.Database
	tbl *mongo.Collection
}

func NewVideoRepository(db *mongo.Database) *videoRepository {
	tbl := db.Collection("videos")
	return &videoRepository{
		db:  db,
		tbl: tbl,
	}
}

func (vr *videoRepository) Create(ctx context.Context, video *model.Video) (*model.Video, error) {
	_, err := vr.tbl.InsertOne(ctx, video)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (vr *videoRepository) Update(ctx context.Context, video *model.Video) (*model.Video, error) {
	_, err := vr.tbl.UpdateByID(ctx, video.ID, video)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (vr *videoRepository) Delete(ctx context.Context, id string) error {
	_, err := vr.tbl.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (vr *videoRepository) FindByID(ctx context.Context, id string) (*model.Video, error) {
	var video model.Video
	err := vr.tbl.FindOne(ctx, bson.M{"id": id}).Decode(&video)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &video, nil
}

func (vr *videoRepository) FindAll(ctx context.Context, keyword string) ([]*model.Video, error) {
	var opt bson.D
	var cur *mongo.Cursor
	var err error
	if keyword != "" {
		opt = bson.D{{
			Key: "$search", Value: bson.D{
				{Key: "index", Value: "full-text-index"},
				{Key: "text", Value: bson.D{
					{Key: "query", Value: keyword},
					{Key: "path", Value: bson.A{"title", "description"}},
				}},
			},
		}}
		if cur, err = vr.tbl.Aggregate(ctx, mongo.Pipeline{opt}); err != nil {
			return nil, err
		}
	} else {
		if cur, err = vr.tbl.Find(ctx, bson.M{}); err != nil {
			return nil, err
		}
	}
	var videos []*model.Video
	if err = cur.All(ctx, &videos); err != nil {
		return nil, err
	}
	return videos, nil
}
