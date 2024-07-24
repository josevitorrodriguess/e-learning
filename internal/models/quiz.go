package models

type Quiz struct {
	ID        int    `json:"id"`
	Subject   string `json:"subject"`
	Creator   Tutor  `json:"creator"`
	Questions []Test `json:"questions"`
}

type Test struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Response string `json:"response"`
}
