package db

import "time"

// ============================================================================
// MACRO FUNCTIONS
// ============================================================================

func MigrateMacro() {
	db.AutoMigrate(&Macro{})
}

type Macro struct {
	ID        uint64    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Name      string    `gorm:"unique"`
	ProfileID uint64
	Command   string
}
