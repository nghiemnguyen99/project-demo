package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"sum/internal/pkg/container"
	"sum/internal/pkg/db/redis"
	redisEntity "sum/redis/entity"
	"sum/redis/server"
	"sum/redis/sumpb"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"google.golang.org/grpc"
)

func main() {
	//connect mysql
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
		"file:///project-demo/internal/pkg/db/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		panic(err.Error())
	}
	if err = m.Up(); err != nil {
		panic(err.Error())
	}

	//connect redis
	ctx := context.Background()
	redis.ConnectRedis(ctx)

	var dbContainer redisEntity.DBContainer
	dbContainer.DB = container.NewDBContainer(database)

	lis, err := net.Listen("tcp", "localhost:50069")
	if err != nil {
		log.Fatalf("error %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterCreateServiceServer(s, &server.Server{DB: container.NewDBContainer(database)})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server %s", err)
	}
}
