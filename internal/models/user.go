package models

type User struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Login          string   `json:"login"`
	Email          string   `json:"email"`
	CPF            string   `json:"cpf"`
	Role           string   `json:"role"`
	Password       string   `json:"password"`
	Wallet         float32  `json:"wallet"`
	OnGoingCourses []Course `json:"onGoingCourses"`
}