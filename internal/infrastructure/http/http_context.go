package http

import (
	"io"

	"github.com/labstack/echo"
	"golang.org/x/net/context"
)

type Context struct {
	c echo.Context
}

func (c *Context) BodyReader() io.Reader {
	return c.c.Request().Body
}

func (c *Context) HeaderGet(key string) string {
	return c.c.Request().Header.Get(key)
}

func (c *Context) QueryGet(key string) string {
	return c.c.Request().URL.Query().Get(key)
}

func (c *Context) JSON(httpCode int, value interface{}) error {
	return c.c.JSON(httpCode, value)
}

func (c *Context) Context() context.Context {
	return c.c.Request().Context()
}
