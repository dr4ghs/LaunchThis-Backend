package db

import (
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

func init() {
	db.AutoMigrate(&User{})
}

// ============================================================================
// USER TABLE
// ============================================================================

// User table
type User struct {
	ID        uint64     `gorm:"primaryKey"`
	CreatedAt time.Time  `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"-" gorm:"autoUpdateTime"`
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
