package users

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
  email := c.Param("email")
  password := c.Param("password")
  if (len(email) == 0) {
	   c.JSON(http.StatusOK, gin.H{"message": "missing email"})
     return
  }
  if (len(password) == 0) {
     c.JSON(http.StatusOK, gin.H{"message": "missing password"})
     return
  }

  c.JSON(http.StatusOK, gin.H{"message": "okie"})
}