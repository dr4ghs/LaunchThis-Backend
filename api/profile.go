package api

import (
	"fmt"
	"net/http"

	"github.com/agsystem/launchthis-be/db"
	"github.com/gin-gonic/gin"
)

// ============================================================================
// PROFILE API FUNCTIONS
// ============================================================================

// Registers all profile APIs
func RegisterProfileAPI(router *gin.Engine) {
	router.GET("/profiles", getProfilesHandler)
	router.PUT("/profiles", createProfileHandler)
	router.GET("/profiles/:id", getProfileHandler)
	router.DELETE("/profiles/:id", deleteProfileHandler)
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
// PROFILE API RESPONSES
// ============================================================================

// Profile API response
type ProfileDetailResponse struct {
	db.Profile
	Users []*db.APIUser `json:",omitempty"`
}

// Casts a profile record to a profile detail API response
func ProfileToResponse(profile *db.Profile) *ProfileDetailResponse {
	detail := &ProfileDetailResponse{}
	detail.ID = profile.ID
	detail.CreatedAt = profile.CreatedAt
	detail.UpdatedAt = profile.UpdatedAt
	detail.Name = profile.Name
	detail.Terminal = profile.Terminal

	apiUsers := []*db.APIUser{}
	for _, element := range profile.Users {
		apiUsers = append(apiUsers, element.ToAPI())
	}
	detail.Users = apiUsers

	return detail
}

// ============================================================================
// PROFILE API HANDLERS
// ============================================================================

// Lists all profiles
func getProfilesHandler(context *gin.Context) {
	var profiles db.ProfileSlice
	profiles.GetProfiles()

	response := []*ProfileDetailResponse{}
	for _, element := range profiles {
		response = append(response, ProfileToResponse(&element))
	}

	context.IndentedJSON(http.StatusOK, response)
}

// Retrieves a profile given its ID
func getProfileHandler(context *gin.Context) {
	profile := getProfile(context)
	if profile == nil {
		return
	}

	context.IndentedJSON(http.StatusOK, ProfileToResponse(profile))
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

	context.IndentedJSON(http.StatusOK, ProfileToResponse(&profile))
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
