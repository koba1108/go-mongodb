package database

import (
	"context"

	"github.com/koba1108/go-mongodb/internals/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db  *mongo.Database
	tbl *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *userRepository {
	tbl := db.Collection("users")
	return &userRepository{
		db:  db,
		tbl: tbl,
	}
}

func (ur *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	_, err := ur.tbl.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	_, err := ur.tbl.UpdateByID(ctx, user.ID, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) Delete(ctx context.Context, id string) error {
	_, err := ur.tbl.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user *model.User
	if err := ur.tbl.FindOne(ctx, bson.M{"id": id}).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user *model.User
	if err := ur.tbl.FindOne(ctx, bson.M{"email": email}).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	cursor, err := ur.tbl.Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
