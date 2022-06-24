package db

import (
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

func init() {
	db.AutoMigrate(&Profile{})
}

// ============================================================================
// PROFILE TABLE
// ============================================================================

// Profile table
type Profile struct {
	ID        uint64    `gorm:"primaryKey"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
	Name      string
	Users     []*User  `json:",omitempty" gorm:"many2many:user_profiles"`
	Macros    []*Macro `json:",omitempty" gorm:"foreignKey:ProfileID"`
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
