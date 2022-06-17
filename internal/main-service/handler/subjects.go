package handler

import (
	"net/http"
	"sum/internal/modules/subjects/entity"
	"sum/internal/modules/subjects/repository"

	"github.com/gin-gonic/gin"
)

type HTTPSubjectsHandler struct {
	SubjectRepository repository.SubjectsRepositoryInterface
}

func NewHTTPSubjectHandlerInterface(subjectRepo repository.SubjectsRepositoryInterface) HTTPSubjectHandlerInterface {
	return &HTTPSubjectsHandler{SubjectRepository: subjectRepo}
}

type HTTPSubjectHandlerInterface interface {
	CreateSubject(c *gin.Context)
}

func (h *HTTPSubjectsHandler) CreateSubject(c *gin.Context) {
	name := c.Query("name")
	subject := entity.Subjects{Name: &name}
	data, err := h.SubjectRepository.Create(subject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
	return
}
