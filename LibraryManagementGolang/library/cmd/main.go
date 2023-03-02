package main

import (
	"context"
	libraryservice "libmanage/library/libraryService"
	repository "libmanage/library/repository/mongo"
	"libmanage/library/transport"
	userRepository "libmanage/user/repository/mongo"
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
	repoStructObj := repository.NewRepoConstructor(client, "Library")

	// Depenedency Repos
	userRepoObj := userRepository.NewRepoConstructor(client, "User")

	// Depenedency services
	userSvcObj := userservice.NewUserServiceConstructor(userRepoObj)

	// svc object
	librarySvcObj := libraryservice.NewLibraryServiceConstructor(repoStructObj, userSvcObj)

	// GIN
	router := gin.Default()
	v1 := router.Group("") // "/api/v1" for gateway

	// LibraryRouterGroup
	//Router Group implementation reside in Gateway Code, Here implemented for this Microservice Main method
	libraryRouterGroup := v1.Group("/library")
	{
		// transport Obj
		routerStruct := transport.NewLibraryRoutesConstructor(librarySvcObj, libraryRouterGroup)
		routerStruct.RegisterLibraryRoutes()

	}

	// Run Server
	router.Run(":9000")
}
