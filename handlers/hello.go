package handlers
import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func HelloPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome",
	})
}
func Getname(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}