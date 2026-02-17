package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mounikavari9/Golang-jwt/controllers"

)

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}
