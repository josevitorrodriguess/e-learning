package models


type User struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Login          string   `json:"login"`
	CPF            string   `json:"cpf"`
	Password       string   `json:"password"`
	Wallet         float32  `json:"wallet"`
	OnGoingCourses []Course `json:"onGoingCourses"`
}

type Tutor struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Login          string   `json:"login"`
	CPF            string   `json:"cpf"`
	Password       string   `json:"password"`
	Wallet         float32  `json:"wallet"`
	OfferedCoruses []Course `json:"offeredCourses"`
}
