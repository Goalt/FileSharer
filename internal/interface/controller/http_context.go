package controller

import (
	"context"
	"io"
)

type HTTPContext interface {
	HeaderGet(string) string
	BodyReader() io.Reader
	JSON(httpCode int, value interface{}) error
	Context() context.Context
	QueryGet(key string) string
}
