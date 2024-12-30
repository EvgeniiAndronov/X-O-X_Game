package server

import "github.com/gin-gonic/gin"

func SetUpRouting(router *gin.Engine) {
	router.POST("/get_cookie", CreateCookie)
	//email - email

	router.GET("/create_user", CreateUser)
	//userLogin - userLogin
	//userName - userName
	//userPassword - userPassword

	router.GET("/Move", CheckCookie(), func(context *gin.Context) { context.JSON(200, gin.H{"message": "cookie is good foe you"}) })
	//id_user
	//coockie - my_cookie

}
