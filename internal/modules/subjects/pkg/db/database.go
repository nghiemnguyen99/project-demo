package db

import (
	"database/sql"
	"sum/internal/modules/subjects/entity"
)

type DB struct {
	DB *sql.DB
}

type DBInterface interface {
	Create(subject entity.Subjects) (entity.Subjects, error)
	GetByID(subjectID int) (entity.Subjects, error)
}

func NewDBInterface(db *sql.DB) DBInterface {
	return &DB{DB: db}
}

func (db *DB) Create(subject entity.Subjects) (entity.Subjects, error) {
	res, err := db.DB.Exec(`INSERT INTO demo.subjects (name) VALUES (?)`, subject.Name)
	if err != nil {
		return entity.Subjects{}, err
	}
	id, err := res.LastInsertId()
	return entity.Subjects{ID: int(id), Name: subject.Name}, err
}

func (db *DB) GetByID(subjectID int) (entity.Subjects, error) {
	res := db.DB.QueryRow(`SELECT * FROM demo.subjects where id = ?`, subjectID)
	var subject entity.Subjects
	err := res.Scan(&subject.ID, &subject.Name)
	if err != nil {
		return entity.Subjects{}, err
	}
	return subject, nil
}
