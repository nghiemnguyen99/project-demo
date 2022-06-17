package container

import (
	subjectRepo "sum/internal/modules/subjects/repository"
	usersRepo "sum/internal/modules/users/repository"
)

type RepositoryContainer struct {
	UsersRepo   usersRepo.UsersRepositoryInterface
	SubjectRepo subjectRepo.SubjectsRepositoryInterface
}

func NewRepositoryContainer(db *DatabaseContainer) *RepositoryContainer {
	return &RepositoryContainer{
		UsersRepo:   usersRepo.NewUsersRepository(db.UsersDB),
		SubjectRepo: subjectRepo.NewSubjectRepository(db.SubjectsDB),
	}
}
