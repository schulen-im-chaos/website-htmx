package web

import (
	"htmx/db"
	"htmx/web/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMaterials(c *gin.Context) {
	subjects, err := db.SubjectList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "", template.Page(c, "Lernmaterialien", template.Materials(subjects)))
}

func GetMaterialsGrades(c *gin.Context) {
	name := c.Param("name")

	subject, err := db.FetchSubjectByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "", template.Page(c, "Lernmaterialien", template.MaterialGrades(subject.Name)))
}

func PostMaterialsGrades(c *gin.Context) {
	name := c.Param("name")
	query := c.PostForm("q")

	items, err := db.ItemsBySubjectAndGradeMap(name, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "", template.MaterialGradesSearch(items))
}
