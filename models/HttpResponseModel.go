package models

// Http response model
type ResponseModel struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
}


func EmptyArray() []interface{}{
	return make([]interface{},0)
}

