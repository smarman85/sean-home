package server

import (
	"fmt"
	"net/http"

	"seanHome/pkg/data"

	"github.com/gin-gonic/gin"
)

func Run(data data.Data) {
	fmt.Println("hello from server")
	router := gin.Default()

	web := router.Group("/")
	{
		router.LoadHTMLGlob("static/html/*")
		router.Static("/static/assets/", "./static/assets")

		web.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "Welcome",
			})
		})

		web.GET("/workouts", func(c *gin.Context) {
			c.HTML(http.StatusOK, "workouts.tmpl", gin.H{
				"fastHands": data.FastHands,
				"metadata":  data.MetaData,
				"weights":   data.Weights,
				"cardio":    data.Cardio,
			})
		})

		web.GET("/drinks/:drink", func(c *gin.Context) {
			c.HTML(http.StatusOK, "drink.tmpl", gin.H{
				"drink":     data.Drinks[c.Param("drink")],
				"drinkName": c.Param("drink"),
			})
		})

		web.GET("/misic/:category", func(c *gin.Context) {
			c.HTML(http.StatusOK, "misic.tmpl", gin.H{
				"category": data.Misic[c.Param("category")],
			})
		})
	}

	router.Run(":5050")
}
