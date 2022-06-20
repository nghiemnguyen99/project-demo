package client

import (
	"sum/internal/modules/subjects/entity"

	"google.golang.org/grpc"
)

type ClientSubjectServer struct {
	CC *grpc.ClientConn
}

type ClientSubjectRepoInterface interface {
	Create(user entity.Subjects) (entity.Subjects, error)
}

func NewClientSubjectRepoInterface(cc *grpc.ClientConn) ClientSubjectRepoInterface {
	return &ClientSubjectServer{CC: cc}
}

func (r *ClientSubjectServer) Create(user entity.Subjects) (entity.Subjects, error) {
	return entity.Subjects{}, nil
}
