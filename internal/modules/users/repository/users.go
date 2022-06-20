package repository

import (
	"sum/internal/modules/users/client"
	"sum/internal/modules/users/entity"
)

type UsersRepository struct {
	CC client.ClientUserRepoInterface
}

type UsersRepositoryInterface interface {
	CreateUser(user entity.Users) (entity.Users, error)
	GetUserByID(userID int) (entity.Users, error)
}

func NewUsersRepository(cc client.ClientUserRepoInterface) UsersRepositoryInterface {
	return &UsersRepository{CC: cc}
}

func (u *UsersRepository) CreateUser(user entity.Users) (entity.Users, error) {
	newUser := entity.Users{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		MSSV:      user.MSSV,
		SubjectID: user.SubjectID,
	}
	return u.CC.Create(newUser)
}

func (u *UsersRepository) GetUserByID(userID int) (entity.Users, error) {
	return u.CC.GetUserByID(userID)
}
