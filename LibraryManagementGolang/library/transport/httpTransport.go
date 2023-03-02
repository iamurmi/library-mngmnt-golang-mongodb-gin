package transport

import (
	customresponse "libmanage/common/customResponse"

	"libmanage/library/domain"
	libraryservice "libmanage/library/libraryService"

	"github.com/gin-gonic/gin"
)

type libraryRoutesStruct struct {
	librarySvcContract libraryservice.LibraryServiceContract

	routerGroup *gin.RouterGroup
}

func NewLibraryRoutesConstructor(libSvcCont libraryservice.LibraryServiceContract, rG *gin.RouterGroup) *libraryRoutesStruct {
	return &libraryRoutesStruct{
		librarySvcContract: libSvcCont,
		routerGroup:        rG,
	}
}

func (r *libraryRoutesStruct) AddBookHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var book domain.Book
		err := c.ShouldBindJSON(&book)
		if err != nil {
			resp := customresponse.FailedResponse(nil, err)
			c.JSON(400, resp)
			return
		}
		err = r.librarySvcContract.AddBook(c, book)
		if err != nil {
			resp := customresponse.FailedResponse(nil, err)
			c.JSON(400, resp)
			return
		}
		resp := customresponse.SuccessResponse(nil, err)
		c.JSON(200, resp)
	}
}
func (r *libraryRoutesStruct) ListBookHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var books []domain.Book
		books, err := r.librarySvcContract.ListBook(c)
		if err != nil {
			resp := customresponse.FailedResponse(nil, err)
			c.JSON(400, resp)
			return
		}
		resp := customresponse.SuccessResponse(books, err)
		c.JSON(200, resp)
	}
}
func (r *libraryRoutesStruct) GetBookHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		book, err := r.librarySvcContract.GetBook(c, id)
		if err != nil {
			resp := customresponse.FailedResponse(nil, err)
			c.JSON(400, resp)
			return
		}
		resp := customresponse.SuccessResponse(book, err)
		c.JSON(200, resp)
	}
}
func (r *libraryRoutesStruct) UserIssueHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		bookId := c.Param("book_id")
		err := r.librarySvcContract.UserIssue(c, userId, bookId)
		if err != nil {
			resp := customresponse.FailedResponse(nil, err)
			c.JSON(400, resp)
			return
		}
		resp := customresponse.SuccessResponse(nil, err)
		c.JSON(200, resp)
	}
}
func (r *libraryRoutesStruct) UserBookRecevHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		bookId := c.Param("book_id")
		err := r.librarySvcContract.UserBookRecev(c, userId, bookId)
		if err != nil {
			resp := customresponse.FailedResponse(nil, err)
			c.JSON(400, resp)
			return
		}
		resp := customresponse.SuccessResponse(nil, err)
		c.JSON(200, resp)
	}
}
func (r *libraryRoutesStruct) RegisterLibraryRoutes() {
	r.routerGroup.POST("/addbook", r.AddBookHandler()) // auth middleware can be added here
	r.routerGroup.GET("/listbook", r.ListBookHandler())
	r.routerGroup.GET("/getbook/:id", r.GetBookHandler())
	r.routerGroup.GET("/issuebook/:user_id/:book_id", r.UserIssueHandler())
	r.routerGroup.GET("/recevbook/:user_id/:book_id", r.UserBookRecevHandler())
}
