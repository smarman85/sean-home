package server

import (
	"log"
	"net/http"
	"os"
	"sean-home/pkg/data"

	"github.com/gin-gonic/gin"
)

func validateTemplate(templateName string) bool {
	valid := false
	_, err := os.Stat("static/html_templates/" + templateName + ".html")
	if err != nil {
		return valid
	} else {
		valid = true
	}
	return valid
}

func Run() {
	router := gin.Default()

	info, err := data.Read()
	if err != nil {
		log.Fatalf("error getting data file: %v", err)
	}

	web := router.Group("/")
	{
		router.LoadHTMLGlob("static/html_templates/*")
		router.Static("/static/assets/", "./static/assets")

		web.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Sean's Site",
			})
		})

		/*
			web.GET("/posts/:postID", func(c *gin.Context) {
				p := c.Param("postID")
				if validateTemplate(p) {
					c.HTML(http.StatusOK, p+".tmpl", gin.H{
						"title": p,
						"post":  p,
					})
				} else {
					c.Redirect(http.StatusMovedPermanently, "/")
				}
			})
		*/

		web.GET("/drinks", func(c *gin.Context) {
			c.HTML(http.StatusOK, "drinks.html", gin.H{
				"title":  "Drinks",
				"drinks": info.Drinks,
			})
		})

		web.GET("/drinks/:drink", func(c *gin.Context) {
			d := c.Param("drink")
			_, ok := info.Drinks[d]
			if ok {
				c.HTML(http.StatusOK, "drink.html", gin.H{
					"drink": info.Drinks[d],
				})
			} else {
				c.Redirect(http.StatusMovedPermanently, "/")
			}
		})

		web.GET("/workouts", func(c *gin.Context) {
			c.HTML(http.StatusOK, "workouts.html", gin.H{
				"title":     "Workouts",
				"workouts":  info.Workouts,
				"fastHands": info.FastHands,
				"metadata":  info.MetaData,
			})
		})

		web.GET("/misic", func(c *gin.Context) {
			c.HTML(http.StatusOK, "misic.html", gin.H{
				"title": "Misic",
				"misic": info.Misic,
			})
		})

	}

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ping": "pong", "status": http.StatusOK,
			})
		})
		v1.GET("/user/:name", func(c *gin.Context) {
			u := c.Param("name")
			if u != "" {
				c.JSON(http.StatusOK, gin.H{
					"hello": c.Param("name"),
				})
			} else {
				c.Redirect(http.StatusMovedPermanently, "/")
			}
		})
	}

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"title": "oops",
		})
	})

	router.Run(":8080")
}
