package users

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	DOB      string `json:"dob"`
	Roles    string `json:"roles"`
}
