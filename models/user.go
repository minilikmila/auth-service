package models

type User struct {
	ID string `json:"id,omitempty" gorm:"autoIncrement"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
	FullName string `json:"full_name,omitempty"`
	EmailConfirmed bool `json:"email_confirmed,omitempty"`
	Password		string `json:"-"`
}