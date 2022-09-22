package database

import (
	"context"
	"errors"
	"github.com/koba1108/go-mongodb/internals/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type playerRepository struct {
	db  *mongo.Database
	tbl *mongo.Collection
}

func NewPlayerRepository(db *mongo.Database) *playerRepository {
	tbl := db.Collection("players")
	return &playerRepository{
		db:  db,
		tbl: tbl,
	}
}

func (vr *playerRepository) Create(ctx context.Context, player *model.Player) (*model.Player, error) {
	_, err := vr.tbl.InsertOne(ctx, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (vr *playerRepository) Update(ctx context.Context, player *model.Player) (*model.Player, error) {
	_, err := vr.tbl.UpdateByID(ctx, player.ID, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (vr *playerRepository) Delete(ctx context.Context, id string) error {
	_, err := vr.tbl.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (vr *playerRepository) FindByID(ctx context.Context, id string) (*model.Player, error) {
	var player model.Player
	log.Println("id", id)
	err := vr.tbl.FindOne(ctx, bson.M{"id": id}).Decode(&player)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &player, nil
}

func (vr *playerRepository) FindAll(ctx context.Context) ([]*model.Player, error) {
	var players []*model.Player
	cur, err := vr.tbl.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var player model.Player
		if err = cur.Decode(&player); err != nil {
			return nil, err
		}
		players = append(players, &player)
	}
	return players, nil
}
