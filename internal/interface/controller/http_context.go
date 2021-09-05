package controller

import (
	"context"
)

type HTTPContext interface {
	Context() context.Context

	GetFormFile(size int) ([]byte, string, int, error)
	GetQuery(key string) string

	JSON(httpCode int, value interface{}) error
	File(httpCode int, data []byte, fileName string) error
}
