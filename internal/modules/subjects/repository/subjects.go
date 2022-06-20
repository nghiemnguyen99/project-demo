package repository

import (
	"sum/internal/modules/subjects/client"
	"sum/internal/modules/subjects/entity"
)

type SubjectRepository struct {
	CC client.ClientSubjectRepoInterface
}

type SubjectsRepositoryInterface interface {
	Create(subject entity.Subjects) (entity.Subjects, error)
}

func NewSubjectRepository(cc client.ClientSubjectRepoInterface) SubjectsRepositoryInterface {
	return &SubjectRepository{CC: cc}
}

func (s *SubjectRepository) Create(subject entity.Subjects) (entity.Subjects, error) {
	return s.CC.Create(subject)
}
