package noteModel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model	// Adds some metadata fields to the table
	ID uuid.UUID `gorm:"type:uuid"` // Excplicitly specify the type to be uuid
	Title string
	SubTitle string
	Text string
}