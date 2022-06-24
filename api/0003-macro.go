package api

import (
	"fmt"
	"net/http"

	"github.com/agsystem/launchthis-be/db"
	"github.com/gin-gonic/gin"
)

func init() {
	router.GET("/macros", getMacrosHandler)
	router.POST("/macros", createOrUpdateMacroHandler)
	router.GET("/macros/:id", getMacroHandler)
	router.DELETE("/macros/:id", deleteMacroHandler)
}

func getMacrosHandler(context *gin.Context) {
	var macros db.MacroSlice
	macros.GetMacros()

	context.IndentedJSON(http.StatusOK, macros)
}

func getMacroHandler(context *gin.Context) {
	macro := getMacro(context)
	if macro == nil {
		return
	}

	macro.GetMacro()

	context.IndentedJSON(http.StatusOK, macro)
}

func createOrUpdateMacroHandler(context *gin.Context) {
	var macro db.Macro
	if err := context.ShouldBindJSON(&macro); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint(err),
		})
	}

	(&macro).CreateOrUpdateMacro()

	context.IndentedJSON(http.StatusOK, macro)
}

func deleteMacroHandler(context *gin.Context) {
	macro := getMacro(context)
	if macro == nil {
		return
	}

	macro.DeleteMacro()

	context.Status(http.StatusOK)
}
