package userservice

import (
	"context"
	"libmanage/user/domain"
	repository "libmanage/user/repository/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserServiceContract interface {
	AddUser(ctx context.Context, User domain.User) (err error)
	ListUser(ctx context.Context) (Users []domain.User, err error)
	GetUser(ctx context.Context, id string) (User domain.User, err error)
	UpdateUser(ctx context.Context, id string, book string) (err error)      // not exposed as API, only for internal service call
	UpdateUserRecev(ctx context.Context, id string, book string) (err error) // not exposed as API, only for internal service call
}
type userService struct {
	repoContract repository.UserRepoContract
}

func NewUserServiceConstructor(repoC repository.UserRepoContract) *userService {
	return &userService{
		repoContract: repoC,
	}
}

func (svc *userService) AddUser(ctx context.Context, User domain.User) (err error) {
	User.Id = string(primitive.NewObjectID().Hex())

	err = svc.repoContract.AddUser(ctx, User)
	if err != nil {
		return err
	}
	return nil
}

func (svc *userService) ListUser(ctx context.Context) (Users []domain.User, err error) {
	Users, err = svc.repoContract.ListUser(ctx)
	if err != nil {
		return nil, err
	}
	return Users, nil
}
func (svc *userService) GetUser(ctx context.Context, id string) (User domain.User, err error) {
	User, err = svc.repoContract.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return User, nil
}

func (svc *userService) UpdateUser(ctx context.Context, id string, book string) (err error) {
	user, err := svc.repoContract.GetUser(ctx, id)
	if err != nil {
		return err
	}
	books := user.Books
	books = append(books, book)
	err = svc.repoContract.UpdateUser(ctx, id, books)
	if err != nil {
		return err
	}
	return nil
}
func (svc *userService) UpdateUserRecev(ctx context.Context, id string, book string) (err error) {
	user, err := svc.repoContract.GetUser(ctx, id)
	if err != nil {
		return err
	}
	books := []string{}
	for i := range user.Books {
		if user.Books[i] == book {
			continue
		} else {
			books = append(books, user.Books[i])
		}
	}
	err = svc.repoContract.UpdateUser(ctx, id, books)
	if err != nil {
		return err
	}
	return nil
}
