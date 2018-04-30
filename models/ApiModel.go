package models


type Pagination struct {
	Limit int `json:"limit"`
	Skip int `json:"skip"`
}


type GetAllUser struct {
	Pagination
	Type string `json:"type"`
}