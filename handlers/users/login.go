package users

import(
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}