package libraryservice

import (
	"context"
	"libmanage/library/domain"
	repository "libmanage/library/repository/mongo"
	userservice "libmanage/user/userService"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LibraryServiceContract interface {
	AddBook(ctx context.Context, book domain.Book) (err error)
	ListBook(ctx context.Context) (books []domain.Book, err error)
	GetBook(ctx context.Context, id string) (book domain.Book, err error)
	UserIssue(ctx context.Context, UserId string, bookId string) (err error)
	UserBookRecev(ctx context.Context, UserId string, bookId string) (err error)
}
type libraryService struct {
	repoContract repository.LibraryRepoContract

	userSvc userservice.UserServiceContract
}

func NewLibraryServiceConstructor(repoC repository.LibraryRepoContract, userSvcCont userservice.UserServiceContract) *libraryService {
	return &libraryService{
		repoContract: repoC,
		userSvc:      userSvcCont,
	}
}

func (svc *libraryService) AddBook(ctx context.Context, book domain.Book) (err error) {
	book.Id = string(primitive.NewObjectID().Hex())

	err = svc.repoContract.AddBook(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (svc *libraryService) ListBook(ctx context.Context) (books []domain.Book, err error) {
	books, err = svc.repoContract.ListBook(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (svc *libraryService) GetBook(ctx context.Context, id string) (book domain.Book, err error) {
	book, err = svc.repoContract.GetBook(ctx, id)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (svc *libraryService) UserIssue(ctx context.Context, UserId string, bookId string) (err error) {

	// when User issued a book Book collection of Library db Quantity of that book decrement

	// update book quantity
	err = svc.repoContract.UpdateBookQuantity(ctx, bookId, -1)
	if err != nil {
		return
	}
	// Get book
	// Update user's book section
	svc.userSvc.UpdateUser(ctx, UserId, bookId)

	return nil
}

func (svc *libraryService) UserBookRecev(ctx context.Context, UserId string, bookId string) (err error) {

	// when Library recev a book from User Book collection of Library db Quantity of that book Incremented

	// update book quantity
	err = svc.repoContract.UpdateBookQuantity(ctx, bookId, 1)
	if err != nil {
		return
	}
	// Get book
	// Update user's book section
	svc.userSvc.UpdateUserRecev(ctx, UserId, bookId)

	return nil
}
