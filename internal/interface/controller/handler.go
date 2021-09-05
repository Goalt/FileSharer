package controller

import (
	"net/http"

	"github.com/Goalt/FileSharer/internal/errors"
)

type handler struct {
}

func (h *handler) Ok(ctx HTTPContext, body interface{}) error {
	return ctx.JSON(http.StatusAccepted, body)
}

func (h *handler) Fail(ctx HTTPContext, body interface{}) error {
	if httpError, ok := body.(errors.HttpError); ok {
		return ctx.JSON(httpError.ResponseCode, httpError)
	}

	return ctx.JSON(http.StatusBadRequest, body)
}

func (h *handler) File(ctx HTTPContext, data []byte, fileName string) error {
	return ctx.File(http.StatusOK, data, fileName)
}
