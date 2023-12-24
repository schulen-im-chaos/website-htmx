package main

import (
	"fmt"
	"htmx/api"
	"htmx/db"
	"htmx/web"
	"htmx/web/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := db.Open()

	if err != nil {
		println("Error:", err)
	}

	router := gin.Default()
	router.HTMLRender = &template.TemplRender{}

	// load env
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// adding admin from env
	db.DeleteLogin(os.Getenv("ADMIN_USERNAME"))
	db.AddLogin(&db.Login{
		User:     os.Getenv("ADMIN_USERNAME"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		Permission: db.Permissions{
			AccessAdmin: db.Write,
		},
	})

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
	router.GET("/materials", web.GetMaterials)
	router.GET("/materials/kind/:name", web.GetMaterialsGrades)
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

	// api
	router.POST("/item", api.AuthMiddleware(api.AdminWrite), api.PostItem)
	router.DELETE("/item/:id", api.AuthMiddleware(api.AdminWrite), api.DeleteItem)
	router.POST("/login", api.AuthMiddleware(api.AdminWrite), api.PostLogin)
	router.DELETE("/login/:id", api.AuthMiddleware(api.AdminWrite), api.DeleteLogin)

	router.Run("localhost:8080")
}
