package api

import (
	"fmt"
	"net/http"

	"github.com/agsystem/launchthis-be/db"
	"github.com/gin-gonic/gin"
)

func init() {
	router.GET("/users", getUsersHandler)
	router.PUT("/users", createUserHandler)
	router.GET("/users/:id", getUserHandler)
	router.DELETE("/users/:id", deleteUserHandler)
	router.GET("/users/:id/profiles", getUserProfilesHandler)
	router.PUT("/users/:id/profiles", addNewProfileHandler)
	router.PATCH("/users/:id/profiles", addProfileHandler)
	router.DELETE("/users/:id/profiles", removeProfileHandler)
}

// ============================================================================
// USER API REQUESTS
// ============================================================================

// Request for the update profile API
type UpdateProfileRequest struct {
	ProfileID uint64
}

// ============================================================================
// USER API HANDLERS
// ============================================================================

// Retrieves all users
func getUsersHandler(context *gin.Context) {
	var users db.UserSlice
	users.GetUsers()

	context.IndentedJSON(http.StatusOK, users)
}

// Retrieves a user given his ID
func getUserHandler(context *gin.Context) {
	user := getUser(context)
	if user == nil {
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

// Creates a new user
func createUserHandler(context *gin.Context) {
	var user db.User
	if err := context.BindJSON(&user); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	(&user).CreateUser()

	context.IndentedJSON(http.StatusOK, user)
}

// Deletes a user given his ID
func deleteUserHandler(context *gin.Context) {
	user := getUser(context)
	if user == nil {
		return
	}

	user.DeleteUser()

	context.Status(http.StatusOK)
}

// Retrieves all user's profiles given the user's ID
func getUserProfilesHandler(context *gin.Context) {
	user := getUser(context)
	if user == nil {
		return
	}

	var profiles db.ProfileSlice
	profiles.GetUserProfiles(user)

	context.IndentedJSON(http.StatusOK, profiles)
}

// Creates a new profile and adds it to a user given the user's ID
func addNewProfileHandler(context *gin.Context) {
	user := getUser(context)
	if user == nil {
		return
	}

	var profile db.Profile
	if err := context.ShouldBindJSON(&profile); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint(err),
		})

		return
	}

	(&profile).AssignProfile(user)
	(&profile).CreateProfile()

	context.IndentedJSON(http.StatusOK, profile)
}

// Adds an existing profile to an user given the user and profile IDs
func addProfileHandler(context *gin.Context) {
	user := getUser(context)
	if user == nil {
		return
	}

	var request UpdateProfileRequest
	if err := context.BindJSON(&request); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	profile := &db.Profile{
		ID: request.ProfileID,
	}
	profile.GetProfile()

	profile.AssignProfile(user)

	context.Status(http.StatusOK)
}

// Removes a profile from a user given the user and profile IDs
func removeProfileHandler(context *gin.Context) {
	user := getUser(context)
	if user == nil {
		return
	}

	var request UpdateProfileRequest
	if err := context.BindJSON(&request); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	profile := &db.Profile{
		ID: request.ProfileID,
	}
	profile.GetProfile()

	profile.DismissProfile(user)

	context.Status(http.StatusOK)
}
