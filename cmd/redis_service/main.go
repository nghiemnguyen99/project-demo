package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"strconv"
	"sum/internal/modules/users/entity"
	"sum/internal/pkg/container"
	"sum/internal/pkg/db/redis"
	"sum/redis/sumpb"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"google.golang.org/grpc"
)

type server struct {
	DB *container.DatabaseContainer
}

var dbContainer server

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
	dbContainer.DB = container.NewDBContainer(database)

	lis, err := net.Listen("tcp", "localhost:50069")
	if err != nil {
		log.Fatalf("error %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterCreateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server %s", err)
	}
}

func (*server) Create(ctx context.Context, req *sumpb.UsersRequest) (*sumpb.UsersResponse, error) {
	firstName := req.GetUser().GetFirstName()
	lastName := req.GetUser().GetLastName()
	mssv := req.GetUser().GetMSSV()
	subjectID := int(req.GetUser().GetSubjectID())

	user := entity.Users{
		FirstName: &firstName,
		LastName:  &lastName,
		MSSV:      &mssv,
		SubjectID: &subjectID,
	}

	res, err := dbContainer.DB.UsersDB.Create(user)
	if err != nil {
		return &sumpb.UsersResponse{}, err
	}

	return &sumpb.UsersResponse{
		ID:        int32(res.ID),
		LastName:  req.GetUser().GetLastName(),
		FirstName: req.GetUser().GetFirstName(),
		MSSV:      req.GetUser().GetMSSV(),
		SubjectID: req.GetUser().GetSubjectID(),
	}, redis.SetToRedis("users_"+strconv.Itoa(int(res.ID)), res)
}

func (*server) GetUserByID(ctx context.Context, req *sumpb.GetUserByIDRequest) (*sumpb.UsersResponse, error) {
	userID := req.GetUserID().GetUserID()
	var user entity.Users

	err := redis.GetFromRedis("users_"+strconv.Itoa(int(userID)), &user)
	if err != nil {
		fmt.Println(err)
	}
	if user.ID > 0 {
		return &sumpb.UsersResponse{
			ID:        int32(user.ID),
			FirstName: *user.FirstName,
			LastName:  *user.LastName,
			MSSV:      *user.MSSV,
			SubjectID: int32(*user.SubjectID),
		}, nil
	}

	res, err := dbContainer.DB.UsersDB.GetByID(int(userID))
	if err != nil {
		return &sumpb.UsersResponse{}, err
	}
	return &sumpb.UsersResponse{
		ID:        int32(res.ID),
		FirstName: *res.FirstName,
		LastName:  *res.LastName,
		MSSV:      *res.MSSV,
		SubjectID: int32(*res.SubjectID),
	}, redis.SetToRedis("users_"+strconv.Itoa(int(userID)), res)
}
