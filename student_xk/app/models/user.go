package models

type User struct {
	Type     uint   `json:"type"`
	ID       uint   `json:"user_id"`
	Account  string `json:"-"`
	Password string `json:"-"`
	Name     string `json:"-"`
}
