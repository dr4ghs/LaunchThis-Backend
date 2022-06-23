package db

import (
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// ============================================================================
// PROFILE TYPES
// ============================================================================

// Terminal type definition
type TerminalType string

// Terminal types
const (
	Shell         TerminalType = "shell"
	Bash                       = "bash"
	CommandPrompt              = "cmd"
	PowerShell                 = "powershell"
)

// ============================================================================
// PROFILE FUNCTIONS
// ============================================================================

// Migrates the profile table
func MigrateProfile() {
	db.AutoMigrate(&Profile{})
}

// ============================================================================
// PROFILE API
// ============================================================================

// API representation of a profile record
type APIProfile struct {
	ID       uint64
	Name     string
	Terminal TerminalType
}

// API profile representation slice
type APIProfileSlice []APIProfile

// ============================================================================
// PROFILE TABLE
// ============================================================================

// Profile table
type Profile struct {
	ID        uint64    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Name      string
	Users     []*User `json:",omitempty" gorm:"many2many:user_profiles"`
	Terminal  TerminalType
	Macros    []*Macro `json:",omitempty"`
}

// Profile table slice
type ProfileSlice []Profile

// === FUNCTIONS ==============================================================

// Profile to string
func (profile Profile) String() string {
	json, err := json.Marshal(profile)
	if err != nil {
		panic(err)
	}

	return string(json)
}

// Casts a profile to an API representation
func (profile *Profile) ToAPI() *APIProfile {
	apiProfile := APIProfile{
		ID:       profile.ID,
		Name:     profile.Name,
		Terminal: profile.Terminal,
	}

	return &apiProfile
}

// Casts a profile slice to an API slice representation
func (slice *ProfileSlice) ToAPISlice() *APIProfileSlice {
	apiSlice := APIProfileSlice{}
	for _, element := range *slice {
		apiSlice = append(apiSlice, *element.ToAPI())
	}

	return &apiSlice
}

// Lists all profiles
func (profiles *ProfileSlice) GetProfiles() {
	db.Preload("Users").Find(profiles)
}

// Retrieves a user's profiles
func (profiles *ProfileSlice) GetUserProfiles(user *User) {
	db.Model(user).Association("Profiles").Find(profiles)
}

// Retrieves a profile
func (profile *Profile) GetProfile() error {
	result := db.Preload("Users").First(profile)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return nil
}

// Creates a profile
func (profile *Profile) CreateProfile() {
	db.Preload("Users").Create(profile)
}

// Deletes a profile
func (profile *Profile) DeleteProfile() error {
	if err := db.Delete(profile).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}

// Adds a user to the profile
func (profile *Profile) AssignProfile(user *User) {
	profile.Users = append(profile.Users, user)

	db.Save(profile)
}

// Removes an user from the profile
func (profile *Profile) DismissProfile(user *User) error {
	users := profile.Users
	for index, element := range users {
		if element == user {
			users[index] = users[len(users)-1]
			users[len(users)-1] = nil
			users = users[:len(users)-1]

			db.Save(profile)

			return nil
		}
	}

	return gorm.ErrRecordNotFound
}
