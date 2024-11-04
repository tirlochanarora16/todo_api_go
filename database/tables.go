package database

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name     string    `json:"name"`
	Email    string    `gorm:"uniqueIndex" json:"email"`
	Password string    `json:"password"`
	Tasks    []Task    `json:"tasks"`
}

type Task struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      uuid.UUID `json:"user_id"`
}
