package main

import (
	"htmx/mem"
	"htmx/web"
	"htmx/web/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	err := mem.ImportFiles()
	if err != nil {
		println("Error:", err)
	}

	gin.SetMode(gin.ReleaseMode)
	gin.Default().SetTrustedProxies([]string{})

	router := gin.Default()
	router.HTMLRender = &template.TemplRender{}

	// server files
	router.Static("/static", "./static")
	router.StaticFile("/favicon.png", "./static/favicon.png")

	// web
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", template.Page(c, "Schulen im Chaos", template.Index()))
	})
	router.GET("/help", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", template.Page(c, "Wie funktioniert's?", template.Help()))
	})
	router.GET("/materials", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", template.Page(c, "Lernmaterialien", template.Materials(mem.Subjects)))

	})
	router.GET("/materials/kind/:name", web.GetMaterialsGrades)
	router.POST("/materials/kind/:name", web.PostMaterialsGrades)
	router.GET("/materials/contribute", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", template.Page(c, "Lernmaterialien Hinzufügen", template.MaterialsContribute()))
	})
	router.GET("/aim", func(c *gin.Context) { c.HTML(http.StatusOK, "", template.Page(c, "Unser Ziel", template.Aim())) })
	router.GET("/team", func(c *gin.Context) { c.HTML(http.StatusOK, "", template.Page(c, "Unser Team", template.Team())) })
	router.GET("/contact", func(c *gin.Context) { c.HTML(http.StatusOK, "", template.Page(c, "Kontakt", template.Contact())) })
	router.GET("/legal/imprint", func(c *gin.Context) { c.HTML(http.StatusOK, "", template.Page(c, "Impressum", template.Imprint())) })
	router.GET("/legal/data-protection", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", template.Page(c, "Datenschutz", template.DataProtection()))
	})

	// 404
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "", template.Page(c, "404 Page Not Found", template.NotFound()))
	})

	router.Run("localhost:8080")
}
