package main

import (
	"database/sql"
	"sum/internal/main-service/container"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	database, err := sql.Open("mysql", "root:abc123@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err.Error())
	}

	_, err = database.Exec("DROP DATABASE IF EXISTS demo")
	if err != nil {
		panic(err)
	}
	_, err = database.Exec("CREATE DATABASE demo")
	if err != nil {
		panic(err)
	}
	_, err = database.Exec("USE demo")
	if err != nil {
		panic(err)
	}
	defer database.Close()

	//create table by migration
	driver, err := mysql.WithInstance(database, &mysql.Config{})
	if err != nil {
		panic(err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///grpc-tutorial/internal/main-service/pkg/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		panic(err.Error())
	}
	if err = m.Up(); err != nil {
		panic(err.Error())
	}

	router := gin.Default()
	db := container.NewDBContainer(database)
	repository := container.NewRepositoryContainer(db)
	handler := container.NewHandlerContainer(repository)
	r := router.Group("/api")
	{
		userHandler := handler.UsersHandler
		userRouter := r.Group("/users")
		{
			userRouter.POST("/create", userHandler.CreateUser)
			userRouter.GET("/get/:user_id/:subject_id", userHandler.GetUserAndSubject)
		}

		subjectHandler := handler.SubjectHandler
		subjectRouter := r.Group("/subjects")
		{
			subjectRouter.POST("/create", subjectHandler.CreateSubject)
		}
	}
	router.Run(":8080")
}
