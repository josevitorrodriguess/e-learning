package models

type Course struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Tutor       User   `json:"tutor"`
	Quizes      []Quiz `json:"quiz"`
}
