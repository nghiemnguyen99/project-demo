package handler

import (
	"net/http"
	"strconv"
	subjectRepository "sum/internal/modules/subjects/repository"
	userEntity "sum/internal/modules/users/entity"
	userRepository "sum/internal/modules/users/repository"

	"github.com/gin-gonic/gin"
)

type HTTPUsersHandlerInterface interface {
	CreateUser(c *gin.Context)
	GetUserByID(c *gin.Context)
}

func NewHTTPHandlerInterface(
	userRepo userRepository.UsersRepositoryInterface,
	subjectRepo subjectRepository.SubjectsRepositoryInterface,
) HTTPUsersHandlerInterface {
	return &HTTPUsersHandler{
		UsersRepo:   userRepo,
		SubjectRepo: subjectRepo,
	}
}

type HTTPUsersHandler struct {
	UsersRepo   userRepository.UsersRepositoryInterface
	SubjectRepo subjectRepository.SubjectsRepositoryInterface
}

func (h *HTTPUsersHandler) GetUserByID(c *gin.Context) {
	var user userEntity.Users
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err = h.UsersRepo.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
	return
}

func (h *HTTPUsersHandler) CreateUser(c *gin.Context) {
	firstName := c.Query("first_name")
	lastName := c.Query("last_name")
	mssv := c.Query("mssv")
	subjectID, err := strconv.Atoi(c.Query("subject_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := userEntity.Users{
		FirstName: &firstName,
		LastName:  &lastName,
		MSSV:      &mssv,
		SubjectID: &subjectID,
	}

	newUser, err := h.UsersRepo.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newUser,
	})
	return
}
