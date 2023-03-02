package repository

import (
	"context"
	"libmanage/library/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var libCollectionName = "Book"

type LibraryRepoContract interface {
	AddBook(ctx context.Context, book domain.Book) (err error)
	ListBook(ctx context.Context) (books []domain.Book, err error)
	GetBook(ctx context.Context, id string) (book domain.Book, err error)
	UpdateBookQuantity(ctx context.Context, id string, incrementedBy int) error
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

func (repo *repoStruct) AddBook(ctx context.Context, book domain.Book) (err error) {
	_, err = repo.mongoClient.Database(repo.db).Collection(libCollectionName).InsertOne(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repoStruct) ListBook(ctx context.Context) (books []domain.Book, err error) {
	cursor, err := repo.mongoClient.Database(repo.db).Collection(libCollectionName).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book domain.Book
		cursor.Decode(&book)
		books = append(books, book)
	}
	return books, nil
}
func (repo *repoStruct) GetBook(ctx context.Context, id string) (book domain.Book, err error) {
	err = repo.mongoClient.Database(repo.db).Collection(libCollectionName).FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}
func (repo *repoStruct) UpdateBookQuantity(ctx context.Context, id string, incrementedBy int) (err error) {
	_, err = repo.mongoClient.Database(repo.db).Collection(libCollectionName).UpdateByID(ctx, id, bson.M{"$inc": bson.M{"quantity": incrementedBy}})
	if err != nil {
		return err
	}
	return nil
}
