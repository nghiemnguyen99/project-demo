package container

import (
	subjectsRepo "sum/internal/modules/subjects/repository"
	usersRepo "sum/internal/modules/users/repository"
)

type RepositoryContainer struct {
	UsersRepo   usersRepo.UsersRepositoryInterface
	SubjectRepo subjectsRepo.SubjectsRepositoryInterface
}

func NewRepositoryContainer(clientContainer *ClientContainer) *RepositoryContainer {
	return &RepositoryContainer{
		UsersRepo:   usersRepo.NewUsersRepository(clientContainer.UserClient),
		SubjectRepo: subjectsRepo.NewSubjectRepository(clientContainer.SubjectClient),
	}
}
