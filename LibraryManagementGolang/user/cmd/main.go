package main

import (
	"context"
	repository "libmanage/user/repository/mongo"
	"libmanage/user/transport"
	userservice "libmanage/user/userService"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// mongoClient
	mongoCtx, mongoCancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer mongoCancel()
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(mongoCtx)

	// repo object

	repoStructObj := repository.NewRepoConstructor(client, "User")

	// svc object

	userSvcObj := userservice.NewUserServiceConstructor(repoStructObj)

	// GIN
	router := gin.Default()
	v1 := router.Group("") // "/api/v1" for gateway
	// LibraryRouterGroup
	//Router Group implementation reside in Gateway Code if we use GRPC, Here implemented for this Microservice Main method
	userRouterGroup := v1.Group("/user")
	{
		// transport Obj
		routerStruct := transport.NewuserRoutesConstructor(userSvcObj, userRouterGroup)
		routerStruct.RegisteruserRoutes()

	}

	// Run Server
	router.Run(":9001")
}
