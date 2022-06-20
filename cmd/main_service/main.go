package main

import (
	"log"
	"sum/internal/pkg/container"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50069", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect : %v", err)
	}
	defer conn.Close()

	router := gin.Default()
	clientContainer := container.NewClientContainer(conn)
	repository := container.NewRepositoryContainer(clientContainer)
	handler := container.NewHandlerContainer(repository)
	r := router.Group("/api")
	{
		userHandler := handler.UsersHandler
		userRouter := r.Group("/users")
		{
			userRouter.POST("/create", userHandler.CreateUser)
			userRouter.GET("/get/:user_id", userHandler.GetUserByID)
		}

		subjectHandler := handler.SubjectHandler
		subjectRouter := r.Group("/subjects")
		{
			subjectRouter.POST("/create", subjectHandler.CreateSubject)
		}
	}
	router.Run(":8080")
}
