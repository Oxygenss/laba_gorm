package handler

import (
	"net/http"
	"strconv"

	"github.com/Oxygenss/laba_gorm/internal/models"
	"github.com/Oxygenss/laba_gorm/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func GetAll(c *gin.Context) {
	students, err := repository.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errors.Wrap(err, "get all handler").Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func CreateStudent(c *gin.Context) {
	var student models.Students

	err := c.ShouldBind(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parse form"})
		return
	}

	err = repository.CreateStudent(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create student"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func DeleteStudent(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "parse ID"})
		return
	}

	if err := repository.DeleteStudent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete student"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}
