package api

import "github.com/gin-gonic/gin"

func Run() {
	router := gin.Default()
	RegisterUserAPI(router)
	RegisterProfileAPI(router)
	router.Run()
}
