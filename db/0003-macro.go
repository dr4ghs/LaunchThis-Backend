package db

import (
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

func init() {
	db.AutoMigrate(&Macro{})
}

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
// MACRO TABLE
// ============================================================================

// Macro table
type Macro struct {
	ID        uint64       `gorm:"primaryKey"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime"`
	Name      string       `gorm:"unique,not null"`
	ProfileID uint64       `json:"-"`
	Terminal  TerminalType `gorm:"not null"`
	Command   string       `gorm:"not null"`
}

// Macro table slice
type MacroSlice []Macro

// === FUNCTIONS ==============================================================

// Macro to string
func (macro Macro) String() string {
	json, err := json.Marshal(macro)
	if err != nil {
		panic(err)
	}

	return string(json)
}

func (macros *MacroSlice) GetMacros() {
	db.Preload("Profiles").Find(macros)
}

func (macros *MacroSlice) GetProfileMacros(profile *Profile) {
	db.Preload("Profiles").Where("ID = ?", profile.ID).Find(macros)
}

func (macro *Macro) GetMacro() error {
	result := db.Preload("Profiles").First(macro)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return nil
}

func (macro *Macro) CreateOrUpdateMacro() error {
	preload := db.Preload("Profiles")
	if preload.First(macro); macro != nil {
		if err := db.Save(macro).Error; err != nil {
			return err
		}
	} else if err := preload.FirstOrCreate(macro).Error; err != nil {
		return err
	}

	return nil
}

func (macro *Macro) DeleteMacro() error {
	if err := db.Delete(macro).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}
