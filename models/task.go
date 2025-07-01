package models
import (
	"gorm.io/gorm"
)
type Task struct {
	gorm.Model
	Message      string     `gorm:"not null"`
	Delay        int 
	Executed  bool   `gorm:"default:false"`

}