package routes

import(
	"github.com/mounikavari9/Golang-jwt/controllers"
	"github.com/gin-gonic/gin"
	"github.com/mounikavari9/Golang-jwt/middleware"

)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())

	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
