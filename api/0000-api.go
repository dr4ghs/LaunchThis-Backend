package api

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

var wg sync.WaitGroup

func init() {
	wg.Add(1)

	router = gin.Default()
}

func WaitInit() {
	wg.Wait()
}

func Run() {
	router.Run()
}
