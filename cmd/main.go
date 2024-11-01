package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunatictiol/portfolio/cmd/render"
	"github.com/lunatictiol/portfolio/models"
	"github.com/lunatictiol/portfolio/src"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	r := gin.Default()
	r.HTMLRender = &render.TemplRender{}
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "okay",
		})
	})

	projects := []models.Project{
		{ID: 3, Name: "That pet place", Description: "\"That pet place\" is a mobile application that leverages technology to create a convenient and efficient platform that connects pet owners with local pet stores and services.The system comprises a mobile application developed in Kotlin, and a backend system built with Go. Users can log in, create accounts, manage pet profiles, search for and book appointments at nearby shops and clinics. The server side allows shop owners to manage available services and appointments.", URL: "https://github.com/lunatictiol/that-pet-place-backend-go"},
		{ID: 1, Name: "Rest api using go", Description: "This Go REST API provides endpoints for managing events. Users can register, log in, create events, and view event details. The API utilizes the Gin framework for efficient routing and JWT authentication for secure user sessions.", URL: "https://github.com/lunatictiol/rest-api-with-go"},
		{ID: 2, Name: "worldle-app-clone", Description: "Wordle app clone using react native, expo and clerk.", URL: "https://github.com/lunatictiol/worldle-app"},
	}

	exp := []models.Experience{
		{ID: 1, CompanyName: "MyClnq", Description: "Worked as a React Native intern, where I integrated chat between two apps using WebSockets, helped resolve bugs, and utilized Jira and Agile methodologies.", Duration: "March 2024 - June 2024"},
		{ID: 2, CompanyName: "Tabbly", Description: "Served as a solo developer, where I developed the iOS app, implemented policies, and successfully launched it on the App Store.", Duration: "Part-time"},
	}

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", src.Index())
	})
	r.GET("/projects", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", src.Projects(projects))
	})
	r.GET("/experience", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", src.Experience(exp))
	})
	r.Run()
}
