package container

import (
	"database/sql"
	subjectsDB "sum/internal/modules/subjects/pkg/db"
	usersDB "sum/internal/modules/users/pkg/db"
)

type DatabaseContainer struct {
	UsersDB    usersDB.DBInterface
	SubjectsDB subjectsDB.DBInterface
}

func NewDBContainer(db *sql.DB) *DatabaseContainer {
	return &DatabaseContainer{
		UsersDB:    usersDB.NewDBInterface(db),
		SubjectsDB: subjectsDB.NewDBInterface(db),
	}
}
