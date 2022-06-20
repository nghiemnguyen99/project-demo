package container

import (
	subjectClient "sum/internal/modules/subjects/client"
	UserClient "sum/internal/modules/users/client"

	"google.golang.org/grpc"
)

type ClientContainer struct {
	UserClient    UserClient.ClientUserRepoInterface
	SubjectClient subjectClient.ClientSubjectRepoInterface
}

func NewClientContainer(cc *grpc.ClientConn) *ClientContainer {
	return &ClientContainer{
		UserClient:    UserClient.NewClientUserRepoInterface(cc),
		SubjectClient: subjectClient.NewClientSubjectRepoInterface(cc),
	}
}
