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

func (vr *videoRepository) FindAll(ctx context.Context) ([]*model.Video, error) {
	var videos []*model.Video
	cur, err := vr.tbl.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var video model.Video
		if err = cur.Decode(&video); err != nil {
			return nil, err
		}
		videos = append(videos, &video)
	}
	return videos, nil
}
