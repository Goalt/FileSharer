package http

import (
	"io"
	"strconv"

	"github.com/Goalt/FileSharer/internal/errors"
	"github.com/labstack/echo"
	"golang.org/x/net/context"
)

const (
	formFileParameter = "source"
)

type Context struct {
	c echo.Context
}

func (c *Context) GetQuery(key string) string {
	return c.c.Request().URL.Query().Get(key)
}

func (c *Context) JSON(httpCode int, value interface{}) error {
	return c.c.JSON(httpCode, value)
}

func (c *Context) GetFormFile(size int) ([]byte, string, int, error) {
	request := c.c.Request()
	if err := request.ParseMultipartForm(int64(size)); err != nil {
		return nil, "", 0, nil
	}

	data := make([]byte, size+1)
	file, fileHeader, err := request.FormFile(formFileParameter)
	if err != nil {
		return nil, "", 0, err
	}

	fileSize, err := file.Read(data)
	switch {
	case (err != nil) && !errors.Is(err, io.EOF):
		return nil, "", 0, err
	case fileSize > size:
		return nil, "", 0, errors.ErrMaxFileSize
	}

	return data[:fileSize], fileHeader.Filename, fileSize, nil

}

func (c *Context) Context() context.Context {
	return c.c.Request().Context()
}

func (c *Context) File(httpCode int, data []byte, fileName string) error {
	response := c.c.Response()

	response.Header().Add("Content-Disposition", `attachment; filename="`+fileName+`"`)
	response.Header().Add("Content-Type", "application/octet-stream")
	response.Header().Add("Content-Length", strconv.Itoa(len(data)))

	if _, err := response.Writer.Write(data); err != nil {
		return nil
	}

	return nil
}
