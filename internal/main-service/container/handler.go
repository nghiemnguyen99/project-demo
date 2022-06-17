package container

import (
	"sum/internal/main-service/handler"
)

type HandlerContainer struct {
	UsersHandler   handler.HTTPUsersHandlerInterface
	SubjectHandler handler.HTTPSubjectHandlerInterface
}

func NewHandlerContainer(repoContainer *RepositoryContainer) *HandlerContainer {
	return &HandlerContainer{
		UsersHandler:   handler.NewHTTPHandlerInterface(repoContainer.UsersRepo, repoContainer.SubjectRepo),
		SubjectHandler: handler.NewHTTPSubjectHandlerInterface(repoContainer.SubjectRepo),
	}
}
