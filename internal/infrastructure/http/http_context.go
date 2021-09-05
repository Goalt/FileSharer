package http

import (
	"fmt"
	"io"
	"strconv"

	"github.com/Goalt/FileSharer/internal/errors"
	"github.com/labstack/echo"
	"golang.org/x/net/context"
)

const (
	formFileParameter         = "source"
	contentTypeOctetStream    = "application/octet-stream"
	contentDispositionPattern = `attachment; filename="%v"`
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

	response.Header().Add(echo.HeaderContentDisposition, fmt.Sprintf(contentDispositionPattern, fileName))
	response.Header().Add(echo.HeaderContentType, contentTypeOctetStream)
	response.Header().Add(echo.HeaderContentLength, strconv.Itoa(len(data)))

	if _, err := response.Writer.Write(data); err != nil {
		return nil
	}

	return nil
}

func (c *Context) GetReqId() string {
	return c.c.Response().Header().Get(echo.HeaderXRequestID)
}
