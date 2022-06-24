package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/agsystem/launchthis-be/db"
	"github.com/gin-gonic/gin"
)

func getUser(context *gin.Context) *db.User {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprint(err),
		})
		return nil
	}

	user := &db.User{
		ID: id,
	}
	if err := user.GetUser(); err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprint(err),
		})
		return nil
	}

	return user
}

func getProfile(context *gin.Context) *db.Profile {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprint(err),
		})
		return nil
	}

	profile := &db.Profile{
		ID: id,
	}
	if err := profile.GetProfile(); err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprint(err),
		})
		return nil
	}

	return profile
}

func getMacro(context *gin.Context) *db.Macro {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprint(err),
		})
		return nil
	}

	macro := &db.Macro{
		ID: id,
	}
	if err := macro.GetMacro(); err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprint(err),
		})
		return nil
	}

	return macro
}
