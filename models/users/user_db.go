package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Address   string         `json:"alamat"`
	DOB       string         `json:"dob"`
	Role      string         `json:"role"`
	Password  string         `json:"-"`
}
