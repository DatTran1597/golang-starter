package model

import "time"

type User struct {
	ID        int        `json:"id,omitempty" gorm:"column:id" `
	Name      string     `json:"name,omitempty" gorm:"column:name" binding:"required,lte=21"`
	CreatedAt time.Time  `json:"_,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"_,omitempty" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"_,omitempty" gorm:"column:deleted_at"`
}

func (user *User) IsValid() error {
	return nil
}
