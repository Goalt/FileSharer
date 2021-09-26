package controller

import (
	"github.com/Goalt/FileSharer/internal/domain"
	"github.com/Goalt/FileSharer/internal/errors"
	"github.com/Goalt/FileSharer/internal/usecase/interactor"
	repository "github.com/Goalt/FileSharer/internal/usecase/repository"
	"github.com/go-playground/validator"
)

const tokenQuery = "token_id"

type HTTPController interface {
	Upload(httpCtx HTTPContext) error
	Download(httpCtx HTTPContext) error
}

type httpController struct {
	maxFileSize      int // MaxFileSize in bytes
	fileInteractor   interactor.FileInteractor
	base64Repository repository.Base64Repository
	handler
	*validator.Validate
	logger repository.Logger
}

func NewHTTPController(maxFileSize int, fileInteractor interactor.FileInteractor, logger repository.Logger, base64Repository repository.Base64Repository) *httpController {
	return &httpController{maxFileSize, fileInteractor, base64Repository, handler{}, validator.New(), logger}
}

func (hc *httpController) Upload(httpCtx HTTPContext) error {
	log := hc.logger.WithField("req_id", httpCtx.GetReqId())

	fileData, fileName, _, err := httpCtx.GetFormFile(hc.maxFileSize)
	switch {
	case errors.Is(err, errors.ErrMaxFileSize):
		log.Info(err)
		return hc.Fail(httpCtx, errors.ErrMaxFileSize)
	case err != nil:
		log.Error(err)
	}

	file := domain.File{
		Data:           fileData,
		FileNameOrigin: fileName,
	}

	if err := hc.Validate.Struct(file); err != nil {
		log.Error(err)
		return hc.Fail(httpCtx, errors.ErrFileFormat)
	}

	token, err := hc.fileInteractor.Upload(httpCtx.Context(), file, log)
	if err != nil {
		return hc.Fail(httpCtx, err)
	}

	token.Id = hc.base64Repository.Encode(token.Id)

	return hc.Ok(httpCtx, token)
}

func (hc *httpController) Download(httpCtx HTTPContext) error {
	log := hc.logger.WithField("req_id", httpCtx.GetReqId())

	var err error
	token := domain.Token{Id: httpCtx.GetQuery(tokenQuery)}
	token.Id, err = hc.base64Repository.Decode(token.Id)
	if err != nil {
		log.Error(err)
		return hc.Fail(httpCtx, errors.ErrTokenFormat)
	}

	if err := hc.Validate.Struct(token); err != nil {
		log.Error(err)
		return hc.Fail(httpCtx, errors.ErrTokenFormat)
	}

	file, err := hc.fileInteractor.Download(httpCtx.Context(), token, log)
	if err != nil {
		return hc.Fail(httpCtx, err)
	}

	return hc.File(httpCtx, file.Data, file.FileNameOrigin)
}
