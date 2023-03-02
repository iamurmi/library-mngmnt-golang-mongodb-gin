package transport

import (
	customresponse "libmanage/common/customResponse"
	"libmanage/user/domain"
	userservice "libmanage/user/userService"

	"github.com/gin-gonic/gin"
)

type userRoutesStruct struct {
	userSvcContract userservice.UserServiceContract

	routerGroup *gin.RouterGroup
}

func NewuserRoutesConstructor(libSvcCont userservice.UserServiceContract, rG *gin.RouterGroup) *userRoutesStruct {
	return &userRoutesStruct{
		userSvcContract: libSvcCont,
		routerGroup:     rG,
	}
}

func (r *userRoutesStruct) AddUserHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var User domain.User
		err := c.ShouldBindJSON(&User)
		if err != nil {
			resp := customresponse.FailedResponse(nil, err)
			c.JSON(400, resp)
			return
		}
		err = r.userSvcContract.AddUser(c, User)
		if err != nil {
			resp := customresponse.FailedResponse(nil, err)
			c.JSON(400, resp)
			return
		}
		resp := customresponse.SuccessResponse(nil, err)
		c.JSON(200, resp)
	}
}
func (r *userRoutesStruct) ListUserHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var Users []domain.User
		Users, err := r.userSvcContract.ListUser(c)
		if err != nil {
			resp := customresponse.FailedResponse(nil, err)
			c.JSON(400, resp)
			return
		}
		resp := customresponse.SuccessResponse(Users, err)
		c.JSON(200, resp)
	}
}
func (r *userRoutesStruct) GetUserHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		User, err := r.userSvcContract.GetUser(c, id)
		if err != nil {
			resp := customresponse.FailedResponse(nil, err)
			c.JSON(400, resp)
			return
		}
		resp := customresponse.SuccessResponse(User, err)
		c.JSON(200, resp)
	}
}

func (r *userRoutesStruct) RegisteruserRoutes() {
	r.routerGroup.POST("/addUser", r.AddUserHandler())
	r.routerGroup.GET("/listUser", r.ListUserHandler())
	r.routerGroup.GET("/getUser/:id", r.GetUserHandler())
}
