package database

import (
	"context"
	"errors"
	"github.com/koba1108/go-mongodb/internals/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (pr *playerRepository) Create(ctx context.Context, player *model.Player) (*model.Player, error) {
	_, err := pr.tbl.InsertOne(ctx, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (pr *playerRepository) Update(ctx context.Context, player *model.Player) (*model.Player, error) {
	_, err := pr.tbl.UpdateByID(ctx, player.ID, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (pr *playerRepository) Delete(ctx context.Context, id string) error {
	_, err := pr.tbl.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (pr *playerRepository) FindByID(ctx context.Context, id string) (*model.Player, error) {
	var player model.Player
	err := pr.tbl.FindOne(ctx, bson.M{"id": id}).Decode(&player)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &player, nil
}

func (pr *playerRepository) FindByIdWithVideos(ctx context.Context, playerId string, limit, offset int, sortKey string, orderBy model.OrderBy) ([]*model.PlayerWithVideos, error) {
	// 検索条件
	var matchStage = bson.D{{
		Key: "$match", Value: bson.D{{
			Key: "id", Value: playerId,
		}},
	}}
	// JOIN句
	lookupStage := bson.D{{
		Key: "$lookup", Value: bson.D{
			// JOIN先のコレクション名
			{Key: "from", Value: "videos"},
			// JOIN元のコレクションのフィールド名
			{Key: "localField", Value: "id"},
			// JOIN先のコレクションのフィールド名
			{Key: "foreignField", Value: "playerId"},
			// JOINする対象の条件と並び順
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$sort", Value: bson.D{{
					Key: sortKey, Value: orderBy.Int()},
				}}},
				bson.D{{Key: "$skip", Value: offset}},
				bson.D{{Key: "$limit", Value: limit}},
			}},
			// JOIN結果のフィールド名
			{Key: "as", Value: "videos"},
		}},
	}
	cur, err := pr.tbl.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage})
	if err != nil {
		return nil, err
	}
	var playerVideos []*model.PlayerWithVideos
	if err = cur.All(ctx, &playerVideos); err != nil {
		return nil, err
	}
	return playerVideos, nil
}

func (pr *playerRepository) FindAll(ctx context.Context) ([]*model.Player, error) {
	var players []*model.Player
	cur, err := pr.tbl.Find(ctx, bson.M{})
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
