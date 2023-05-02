package server

import (
	"fmt"
	"net/http"
  "os"

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
      d := c.Param("drink")

      // check if the template exists before rendering it 
      // else redirect to home
      if _, err := os.Stat("static/html/" + d + ".tmpl"); err != nil {
        c.Redirect(http.StatusMovedPermanently, "/")
        return
      }

			c.HTML(http.StatusOK, fmt.Sprintf("%s.tmpl", d), gin.H{
				"drink":     data.Drinks[c.Param("drink")],
				"drinkName": c.Param("drink"),
			})
		})

		web.GET("/misic", func(c *gin.Context) {
			c.HTML(http.StatusOK, "misic.tmpl", gin.H{
				"category": data.Misic,
			})
		})
	}

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.tmpl", gin.H{
			"title": "oops",
		})
	})

	router.Run(":8080")
}
