package repository

import (
	"context"
	"libmanage/user/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollectionName = "User"

type UserRepoContract interface {
	AddUser(ctx context.Context, User domain.User) (err error)
	ListUser(ctx context.Context) (Users []domain.User, err error)
	GetUser(ctx context.Context, id string) (User domain.User, err error)
	UpdateUser(ctx context.Context, id string, book []string) (err error)
}
type repoStruct struct {
	mongoClient *mongo.Client
	db          string
}

func NewRepoConstructor(client *mongo.Client, dbName string) *repoStruct {
	return &repoStruct{
		mongoClient: client,
		db:          dbName,
	}
}

func (repo *repoStruct) AddUser(ctx context.Context, User domain.User) (err error) {
	_, err = repo.mongoClient.Database(repo.db).Collection(userCollectionName).InsertOne(ctx, User)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repoStruct) ListUser(ctx context.Context) (Users []domain.User, err error) {
	cursor, err := repo.mongoClient.Database(repo.db).Collection(userCollectionName).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var User domain.User
		cursor.Decode(&User)
		Users = append(Users, User)
	}
	return Users, nil
}
func (repo *repoStruct) GetUser(ctx context.Context, id string) (User domain.User, err error) {
	err = repo.mongoClient.Database(repo.db).Collection(userCollectionName).FindOne(ctx, bson.M{"_id": id}).Decode(&User)
	if err != nil {
		return domain.User{}, err
	}
	return User, nil
}
func (repo *repoStruct) UpdateUser(ctx context.Context, id string, books []string) (err error) {

	_, err = repo.mongoClient.Database(repo.db).Collection(userCollectionName).UpdateByID(ctx, id, bson.M{"$set": bson.M{"books": books}})
	if err != nil {
		return err
	}
	return nil

}
