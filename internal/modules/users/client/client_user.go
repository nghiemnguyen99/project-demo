package client

import (
	"context"
	"fmt"
	"sum/internal/modules/users/entity"
	"sum/redis/sumpb"

	"google.golang.org/grpc"
)

type ClientUserServer struct {
	CC *grpc.ClientConn
}

type ClientUserRepoInterface interface {
	Create(user entity.Users) (entity.Users, error)
	GetUserByID(userID int) (entity.Users, error)
}

func NewClientUserRepoInterface(cc *grpc.ClientConn) ClientUserRepoInterface {
	return &ClientUserServer{CC: cc}
}

func (cus *ClientUserServer) Create(user entity.Users) (entity.Users, error) {
	c := sumpb.NewCreateServiceClient(cus.CC)

	newUser := sumpb.Users{
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		MSSV:      *user.MSSV,
		SubjectID: int32(*user.SubjectID),
	}

	res, err := c.Create(context.Background(), &sumpb.UsersRequest{User: &newUser})
	if err != nil {
		return entity.Users{}, err
	}
	user.ID = int(res.ID)
	return user, nil
}

func (cus *ClientUserServer) GetUserByID(userID int) (entity.Users, error) {
	c := sumpb.NewCreateServiceClient(cus.CC)
	var user entity.Users
	userIDRequest := sumpb.GetUserById{UserID: int32(userID)}

	res, err := c.GetUserByID(context.Background(), &sumpb.GetUserByIDRequest{UserID: &userIDRequest})
	if err != nil {
		fmt.Println("nghiem=================")
		return entity.Users{}, err
	}
	fmt.Println("-----------------------------")
	subjectID := int(res.SubjectID)
	user = entity.Users{
		ID:        int(res.ID),
		FirstName: &res.FirstName,
		LastName:  &res.LastName,
		MSSV:      &res.MSSV,
		SubjectID: &subjectID,
	}
	return user, nil
}
