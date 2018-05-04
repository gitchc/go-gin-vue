package main

import (
	"github.com/gin-gonic/gin"
	"goginvue/handlers"
	"net/http"
	"math/rand"
)

func Init() {
	r := gin.Default()  // Grouping routes
	// group： v1
	v1 := r.Group("/v1")

	v1.GET("/hello", handlers.HelloPage)
	v1.GET("/hello/:name",handlers.Getname)
	v1.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	v1.GET("/line", func(c *gin.Context) {
		legendData := []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
		xAxisData := []int{120, 240, rand.Intn(500), rand.Intn(500), 150, 230, 180}
		c.JSON(http.StatusOK, gin.H{
			"legend_data": legendData,
			"xAxis_data": xAxisData,
		})
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8000

}

