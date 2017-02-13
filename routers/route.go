package routers

import(
	"gopkg.in/gin-gonic/gin.v1"
	"LandBorrowBook/handlers/users"
)

func New() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
	    c.JSON(200, gin.H{
	        "message": "pong",
	    })
	})


	usersHandler := r.Group("/users")
    {
        usersHandler.POST("/login", users.Login)
        // v1.POST("/submit", submitEndpoint)
    }

	return r
}
