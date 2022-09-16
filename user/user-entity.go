package user

import "time"

type User struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"`
	Occupation   string    `gorm:"type:varchar(100)" json:"occupation"`
	Email        string    `gorm:"uniqueIndex;type:varchar(100);not null" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"password"`
	Role         string    `gorm:"type:varchar(5);default:user" json:"role"`
	Token        string    `gorm:"-" json:"token,omitempty"`
	ImgProfile   string    `gorm:"type:varchar(100)" json:"imgProfile"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
