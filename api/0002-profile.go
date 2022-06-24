package api

import (
	"fmt"
	"net/http"

	"github.com/agsystem/launchthis-be/db"
	"github.com/gin-gonic/gin"
)

func init() {
	router.GET("/profiles", getProfilesHandler)
	router.PUT("/profiles", createProfileHandler)
	router.GET("/profiles/:id", getProfileHandler)
	router.DELETE("/profiles/:id", deleteProfileHandler)
	router.GET("/profiles/:id/macros", getProfileMacrosHandler)
	router.PATCH("/profiles/:id/assign", assignProfileHandler)
}

// ============================================================================
// PROFILE API REQUESTS
// ============================================================================

// Request for assign a profile to a user
type AssignProfileRequest struct {
	UserID uint64
	Remove bool
}

// ============================================================================
// PROFILE API HANDLERS
// ============================================================================

// Lists all profiles
func getProfilesHandler(context *gin.Context) {
	var profiles db.ProfileSlice
	profiles.GetProfiles()

	context.IndentedJSON(http.StatusOK, profiles)
}

// Retrieves a profile given its ID
func getProfileHandler(context *gin.Context) {
	profile := getProfile(context)
	if profile == nil {
		return
	}

	context.IndentedJSON(http.StatusOK, profile)
}

// Creates a new profile
func createProfileHandler(context *gin.Context) {
	var profile db.Profile
	if err := context.ShouldBindJSON(&profile); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	(&profile).CreateProfile()

	context.IndentedJSON(http.StatusOK, profile)
}

// Deletes a profile given its ID
func deleteProfileHandler(context *gin.Context) {
	profile := getProfile(context)
	if profile == nil {
		return
	}

	profile.DeleteProfile()

	context.Status(http.StatusOK)
}

func getProfileMacrosHandler(context *gin.Context) {
	profile := getProfile(context)
	if profile == nil {
		return
	}

	var macros db.MacroSlice
	macros.GetProfileMacros(profile)

	context.IndentedJSON(http.StatusOK, macros)
}

// Assign a profile to a user given the profile and user IDs
func assignProfileHandler(context *gin.Context) {
	profile := getProfile(context)
	if profile == nil {
		return
	}

	var request *AssignProfileRequest
	if err := context.BindJSON(request); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint(err),
		})

		return
	}

	user := &db.User{
		ID: request.UserID,
	}
	user.GetUser()

	if request.Remove {
		profile.DismissProfile(user)
	} else {
		profile.AssignProfile(user)
	}

	context.Status(http.StatusOK)
}
