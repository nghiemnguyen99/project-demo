package repository

import (
	"sum/internal/modules/subjects/entity"
	"sum/internal/modules/subjects/pkg/db"
)

type SubjectRepository struct {
	DB db.DBInterface
}

type SubjectsRepositoryInterface interface {
	Create(subject entity.Subjects) (entity.Subjects, error)
	GetSubjectByID(subjectID int) (entity.Subjects, error)
}

func NewSubjectRepository(db db.DBInterface) SubjectsRepositoryInterface {
	return &SubjectRepository{DB: db}
}

func (s *SubjectRepository) Create(subject entity.Subjects) (entity.Subjects, error) {
	return s.DB.Create(subject)
}

func (s *SubjectRepository) GetSubjectByID(subjectID int) (entity.Subjects, error) {
	return s.DB.GetByID(subjectID)
}
