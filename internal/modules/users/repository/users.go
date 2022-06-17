package repository

import (
	"database/sql"
	"errors"
	"sum/internal/modules/users/entity"
	"sum/internal/modules/users/pkg/db"
)

type UsersRepository struct {
	DB db.DBInterface
}

type UsersRepositoryInterface interface {
	CreateUser(user entity.Users) (entity.Users, error)
	GetUserByID(userID int) (entity.Users, error)
	UpdateInfo(user entity.Users) (entity.Users, error)
}

func NewUsersRepository(db db.DBInterface) UsersRepositoryInterface {
	return &UsersRepository{DB: db}
}

func (u *UsersRepository) CreateUser(user entity.Users) (entity.Users, error) {
	olderUser, err := u.DB.GetByMSSV(*user.MSSV)
	if err != nil && err != sql.ErrNoRows {
		return entity.Users{}, err
	}
	if olderUser.ID > 0 {
		return entity.Users{}, errors.New("user already exists")
	}
	newUser := entity.Users{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		MSSV:      user.MSSV,
		SubjectID: user.SubjectID,
	}
	return u.DB.Create(newUser)
}

func (u *UsersRepository) GetUserByID(userID int) (entity.Users, error) {
	return u.DB.GetByID(userID)
}

func (u *UsersRepository) UpdateInfo(user entity.Users) (entity.Users, error) {
	return entity.Users{}, nil
}
