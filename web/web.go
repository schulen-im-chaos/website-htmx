package web

import (
	"net/http"
	"website/mem"
	"website/web/template"

	"github.com/gin-gonic/gin"
)

func GetMaterialsGrades(c *gin.Context) {
	name := c.Param("name")

	c.HTML(http.StatusOK, "", template.Page(c, "Lernmaterialien - "+name, template.MaterialGrades(name)))
}

func PostMaterialsGrades(c *gin.Context) {
	name := c.Param("name")
	query := c.PostForm("q")
	resources := mem.ResourcesBySubjectAndGradeMap(name, query)

	c.HTML(http.StatusOK, "", template.MaterialGradesResults(resources))
}
