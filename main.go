package main

import (
	"fmt"
	"go-rest-prisma/configs"
	"go-rest-prisma/controllers"
	"go-rest-prisma/helpers"
	repositories "go-rest-prisma/repositories/post"
	"go-rest-prisma/routers"
	service "go-rest-prisma/services/post"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Could not load variable", err)
	}
	fmt.Printf("Start Server" + os.Getenv("GO_PORT"))

	// server := &http.Server{
	// 	Addr:           ":" + os.Getenv("GO_PORT"),
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// server_err := server.ListenAndServe()
	// if server_err != nil {
	// 	panic(server_err)
	// }

	//handle DB
	db, err := configs.ConnectDB()
	helpers.ErrorPanic(err)

	defer db.Prisma.Disconnect()

	// repositories
	postRepository := repositories.NewPostRepository(db)

	// service
	postService := service.NewPostService(postRepository)

	// controller
	postController := controllers.NewPostController(postService)

	// router
	routes := routers.NewRouter(postController)

	server := &http.Server{
		Addr:           ":" + os.Getenv("GO_PORT"),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	if server_err != nil {
		panic(server_err)
	}
}
