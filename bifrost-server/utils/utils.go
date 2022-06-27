package utils

import "github.com/aloks98/bifrost/bifrost-server/models"

func CreateResponse(status int, err interface{}, data interface{}) models.Response {
	return models.Response{
		Status: status,
		Error:  err,
		Data:   data,
	}
}
