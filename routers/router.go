package routers

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ridhompra/controllers"
)

func Routers() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to connect .env")
	}
	port := os.Getenv("PORT")

	app := gin.Default()
	v1 := app.Group("api/v1")

	v1.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "status Ok",
			"port":    ":" + port,
		})
	})
	v1.GET("/book", controllers.GetAll)
	v1.POST("/book", controllers.Create)
	v1.GET("/book/:id", controllers.GetById)
	v1.DELETE("/book/:id", controllers.Delete)
	v1.PUT("/book/:id", controllers.Put)

	v1.POST("/signup", controllers.SignUp)

	app.Run(":" + port)
}
