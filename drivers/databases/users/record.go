package users

import (
	"finalProject/business/users"
	"time"
)

type User struct {
	Id        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `gorm:"index" json:"deletedAt"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
}

func (rec *User) toDomain() users.Domain {
	return users.Domain{
		Id:        rec.Id,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		Name:      rec.Name,
		Email:     rec.Email,
		Address:   rec.Address,
		Password:  rec.Password,
		Role:      rec.Role,
	}
}

func fromDomain(user users.Domain) *User {
	return &User{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Password:  user.Password,
		Role:      user.Role,
	}
}
