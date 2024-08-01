package response

type User struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Login  string  `json:"login"`
	Email  string  `json:"email"`
	CPF    string  `json:"cpf"`
	Role   string  `json:"role"`
	Wallet float32 `json:"wallet"`
}


