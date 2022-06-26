package models

type Response struct {
	Status int8        `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}
