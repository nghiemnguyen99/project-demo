package server

import (
	"context"
	"fmt"
	"strconv"
	userEntity "sum/internal/modules/users/entity"
	"sum/internal/pkg/container"
	"sum/internal/pkg/db/redis"
	"sum/redis/pkg"
	"sum/redis/sumpb"
)

type Server struct {
	DB *container.DatabaseContainer
}

var cr, u = make(chan interface{}), make(chan interface{})

func (s *Server) Create(ctx context.Context, req *sumpb.UsersRequest) (*sumpb.UsersResponse, error) {
	firstName := req.GetUser().GetFirstName()
	lastName := req.GetUser().GetLastName()
	mssv := req.GetUser().GetMSSV()
	subjectID := int(req.GetUser().GetSubjectID())

	user := userEntity.Users{
		FirstName: &firstName,
		LastName:  &lastName,
		MSSV:      &mssv,
		SubjectID: &subjectID,
	}

	res, err := s.DB.UsersDB.Create(user)
	if err != nil {
		return &sumpb.UsersResponse{}, err
	}

	go pkg.HanlderRedis(cr, u, "users_"+strconv.Itoa(int(res.ID)))
	cr <- res
	return &sumpb.UsersResponse{
		ID:        int32(res.ID),
		LastName:  req.GetUser().GetLastName(),
		FirstName: req.GetUser().GetFirstName(),
		MSSV:      req.GetUser().GetMSSV(),
		SubjectID: req.GetUser().GetSubjectID(),
	}, nil
}

func (s *Server) GetUserByID(ctx context.Context, req *sumpb.GetUserByIDRequest) (*sumpb.UsersResponse, error) {
	userID := req.GetUserID().GetUserID()
	var user userEntity.Users

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

	res, err := s.DB.UsersDB.GetByID(int(userID))
	if err != nil {
		return &sumpb.UsersResponse{}, err
	}

	go pkg.HanlderRedis(cr, u, "users_"+strconv.Itoa(int(res.ID)))
	cr <- res

	return &sumpb.UsersResponse{
		ID:        int32(res.ID),
		FirstName: *res.FirstName,
		LastName:  *res.LastName,
		MSSV:      *res.MSSV,
		SubjectID: int32(*res.SubjectID),
	}, nil
}
