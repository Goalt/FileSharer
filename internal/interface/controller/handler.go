package controller

import "net/http"

type handler struct {
}

func (h *handler) Ok(ctx HTTPContext, body interface{}) error {
	return ctx.JSON(http.StatusAccepted, body)
}

func (h *handler) Fail(ctx HTTPContext, body interface{}, statusCode int) error {
	return ctx.JSON(statusCode, body)
}
