package request

import "finalProject/business/users"

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}

func (req *LoginUser) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}
