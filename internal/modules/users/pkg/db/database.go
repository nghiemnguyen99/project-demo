package db

import (
	"database/sql"
	"sum/internal/modules/users/entity"
)

type DB struct {
	DB *sql.DB
}

type DBInterface interface {
	GetByMSSV(mssv string) (entity.Users, error)
	GetByID(subjectID int) (entity.Users, error)
	Create(user entity.Users) (entity.Users, error)
}

func NewDBInterface(db *sql.DB) DBInterface {
	return &DB{DB: db}
}

func (db *DB) GetByMSSV(mssv string) (entity.Users, error) {
	res := db.DB.QueryRow(`SELECT * FROM demo.users where mssv like ?`, mssv)
	var user entity.Users
	err := res.Scan(&user.ID, &user.FirstName, &user.LastName, &user.MSSV, &user.SubjectID)
	if err != nil {
		return entity.Users{}, err
	}
	return user, nil
}

func (db *DB) GetByID(userID int) (entity.Users, error) {
	res := db.DB.QueryRow(`SELECT * FROM demo.users where id = ?`, userID)
	var user entity.Users
	err := res.Scan(&user.ID, &user.FirstName, &user.LastName, &user.MSSV, &user.SubjectID)
	if err != nil {
		return entity.Users{}, err
	}
	return user, nil
}

func (db *DB) Create(user entity.Users) (entity.Users, error) {
	res, err := db.DB.Exec(`INSERT INTO demo.users (first_name, last_name, mssv, subject_id) VALUES (?,?,?,?)`,
		user.FirstName, user.LastName, user.MSSV, user.SubjectID)
	if err != nil {
		return entity.Users{}, err
	}
	id, err := res.LastInsertId()
	user.ID = int(id)
	return user, err
}
