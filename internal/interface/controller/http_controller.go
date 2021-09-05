package controller

import (
	"fmt"

	"github.com/Goalt/FileSharer/internal/domain"
	"github.com/Goalt/FileSharer/internal/errors"
	"github.com/Goalt/FileSharer/internal/usecase/interactor"
	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
	"github.com/go-playground/validator"
)

const (
	contetTypeHeader = "Content-Type"
	multipartPrefix  = "multipart/"
	boundaryKey      = "boundary"
	fileNameHeader   = "filename"
	tokenQuery       = "token_id"
)

type HTTPController interface {
	Upload(httpCtx HTTPContext) error
	Download(httpCtx HTTPContext) error
}

type httpController struct {
	maxFileSize    int // MaxFileSize in bytes
	fileInteractor interactor.FileInteractor
	handler
	*validator.Validate
	logger usecase_repository.Logger
}

func NewHTTPController(maxFileSize int, fileInteractor interactor.FileInteractor, logger usecase_repository.Logger) HTTPController {
	return &httpController{maxFileSize, fileInteractor, handler{}, validator.New(), logger}
}

func (hc *httpController) Upload(httpCtx HTTPContext) error {
	fileData, fileName, _, err := httpCtx.GetFormFile(hc.maxFileSize)
	switch {
	case errors.Is(err, errors.ErrMaxFileSize):
		hc.logger.Info(httpCtx.Context(), err.Error())
		return hc.Fail(httpCtx, errors.ErrMaxFileSize)
	case err != nil:
		hc.logger.Error(httpCtx.Context(), fmt.Sprintf("file read error from form file: %v", err))
	}

	file := domain.File{
		Data:           fileData,
		FileNameOrigin: fileName,
	}

	if err := hc.Validate.Struct(file); err != nil {
		hc.logger.Error(httpCtx.Context(), fmt.Sprintf("input data validate error %v", err))
		return hc.Fail(httpCtx, errors.ErrFileFormat)
	}

	token, err := hc.fileInteractor.Upload(httpCtx.Context(), file)
	if err != nil {
		return hc.Fail(httpCtx, err)
	}

	return hc.Ok(httpCtx, token)
}

func (hc *httpController) Download(httpCtx HTTPContext) error {
	token := domain.Token{Id: httpCtx.GetQuery(tokenQuery)}
	if err := hc.Validate.Struct(token); err != nil {
		hc.logger.Error(httpCtx.Context(), fmt.Sprintf("input data validate error %v", err))
		return hc.Fail(httpCtx, errors.ErrTokenFormat)
	}

	file, err := hc.fileInteractor.Download(httpCtx.Context(), token)
	if err != nil {
		return hc.Fail(httpCtx, err)
	}

	return hc.Ok(httpCtx, file)
}
