package db

import (
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// ============================================================================
// USER FUNCTIONS
// ============================================================================

// Migrates the user table
func MigrateUser() {
	db.AutoMigrate(&User{})
}

// ============================================================================
// USER TABLE
// ============================================================================

// User table
type User struct {
	ID        uint64     `gorm:"primaryKey"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	Username  string     `gorm:"unique"`
	Profiles  []*Profile `json:",omitempty" gorm:"many2many:user_profiles"`
}

// User table slice
type UserSlice []User

// === FUNCTIONS ==============================================================

// User to string
func (user User) String() string {
	json, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	return string(json)
}

// Casts a user to an API representation
func (user *User) ToAPI() *APIUser {
	apiUser := &APIUser{
		ID:       user.ID,
		Username: user.Username,
	}

	return apiUser
}

// Casts a user slice to an API slice representation
func (slice *UserSlice) ToAPISlice() *APIUserSlice {
	apiSlice := APIUserSlice{}
	for _, element := range *slice {
		apiSlice = append(apiSlice, *element.ToAPI())
	}

	return &apiSlice
}

// Lists all users
func (users *UserSlice) GetUsers() {
	db.Preload("Profiles").Find(users)
}

// Retrieves a user by his ID
func (user *User) GetUser() error {
	result := db.Preload("Profiles").First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return nil
}

// Creates a user
func (user *User) CreateUser() {
	db.Create(user)
}

// Deletes a user
func (user *User) DeleteUser() error {
	if err := db.Delete(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}

// ============================================================================
// API USER
// ============================================================================

// API representation of a user record
type APIUser struct {
	ID       uint64
	Username string
}

// API user representation slice
type APIUserSlice []APIUser
