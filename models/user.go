package models

type UserRole uint8

const (
	Guest      UserRole = 0
	NormalUser UserRole = 1
	Admin      UserRole = 2
)

type User struct {
	Model
	Username string   `json:"username" gorm:"index,unique"`
	Email    string   `json:"email" gorm:"index,unique"`
	Password string   `json:"-"`
	Role     UserRole `json:"role" gorm:"default:0"`
	Profile  *Profile `json:"profile"`
}
