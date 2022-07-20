package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type EditUser struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
