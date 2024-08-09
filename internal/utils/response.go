package utils

import (
	"github.com/labstack/echo"
)

type Meta struct {
	StatusCode  int         `json:"status_code"`
	MessageCode string      `json:"message_code"`
	Message     string      `json:"message"`
	Pagination  *Pagination `json:"pagination"`
}

type Pagination struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"Limit"`
}

type ResponseJson struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

func ResponseHandler(c echo.Context, data interface{}, meta Meta) error {
	responData := ResponseJson{
		Meta: meta,
		Data: data,
	}
	return c.JSON(meta.StatusCode, responData)
}

func InstancePagination(total int, offset int, limit int) *Pagination {
	return &Pagination{
		Total:  total,
		Offset: offset,
		Limit:  limit,
	}
}

func InstanceMeta(statusCode int, messageCode string, message string, pagination *Pagination) Meta {
	return Meta{
		StatusCode:  statusCode,
		MessageCode: messageCode,
		Message:     message,
		Pagination:  pagination,
	}
}
