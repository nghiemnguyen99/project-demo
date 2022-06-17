package handler

import (
	"log"
	"net/http"
	"strconv"
	subjectEntity "sum/internal/modules/subjects/entity"
	subjectRepository "sum/internal/modules/subjects/repository"
	userEntity "sum/internal/modules/users/entity"
	userRepository "sum/internal/modules/users/repository"
	"sum/internal/pkg/db/redis"

	"github.com/gin-gonic/gin"
)

type HTTPUsersHandlerInterface interface {
	CreateUser(c *gin.Context)
	GetUserAndSubject(c *gin.Context)
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

func (h *HTTPUsersHandler) GetUserAndSubject(c *gin.Context) {
	var user userEntity.Users
	var subject subjectEntity.Subjects
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	subjectID, err := strconv.Atoi(c.Param("subject_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = redis.GetFromRedis("users_"+strconv.Itoa(userID), &user)
	if err != nil {
		log.Fatalf("get user from redis failed")
	}
	if user.MSSV == nil {
		user, err = h.UsersRepo.GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	err = redis.GetFromRedis("subject_"+strconv.Itoa(subjectID), &subject)
	if err != nil {
		log.Fatalf("get subject from redis failed")
	}
	if subject.Name == nil {
		subject, err = h.SubjectRepo.GetSubjectByID(subjectID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user":    user,
		"subject": subject,
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
