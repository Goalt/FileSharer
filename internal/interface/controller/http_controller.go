package controller

import (
	"errors"
	"fmt"
	"mime"
	"mime/multipart"
	"strings"

	"github.com/Goalt/FileSharer/internal/domain"
	"github.com/Goalt/FileSharer/internal/usecase/interactor"
	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
	"github.com/go-playground/validator"
)

var (
	contetTypeHeader = "Content-Type"
	multipartPrefix  = "multipart/"
	boundaryKey      = "boundary"
	fileNameHeader   = "filename"
	tokenQuery       = "token_id"

	errBadRequest    = errors.New("bad request")
	statusBadRequest = 500
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
	mediaType, params, err := mime.ParseMediaType(httpCtx.HeaderGet(contetTypeHeader))
	if err != nil {
		hc.logger.Error(httpCtx.Context(), "parse media type error", err)
		return hc.Fail(httpCtx, errBadRequest, statusBadRequest)
	}

	if !strings.HasPrefix(mediaType, multipartPrefix) {
		hc.logger.Error(httpCtx.Context(), "media type error", nil)
		return hc.Fail(httpCtx, errBadRequest, statusBadRequest)
	}

	body := httpCtx.BodyReader()
	multipartReader := multipart.NewReader(body, params[boundaryKey])

	part, err := multipartReader.NextPart()
	fmt.Print(part.Header.Get(fileNameHeader))
	switch {
	case err != nil:
		hc.logger.Error(httpCtx.Context(), "multipart read error", err)
		return hc.Fail(httpCtx, errBadRequest, statusBadRequest)
	}

	data := make([]byte, hc.maxFileSize+1)
	fileSize, err := part.Read(data)
	if fileSize == hc.maxFileSize+1 {
		hc.logger.Error(httpCtx.Context(), "max file size", err)
		return hc.Fail(httpCtx, errBadRequest, statusBadRequest)
	}
	if err != nil {
		hc.logger.Error(httpCtx.Context(), "data read error", err)
		return hc.Fail(httpCtx, errBadRequest, statusBadRequest)
	}

	file := domain.File{
		Data:           data,
		FileNameOrigin: part.FileName(),
	}
	if err := hc.Validate.Struct(file); err != nil {
		hc.logger.Error(httpCtx.Context(), "input data validate error", err)
		return hc.Fail(httpCtx, errBadRequest, statusBadRequest)
	}

	token, err := hc.fileInteractor.Upload(httpCtx.Context(), file)
	if err != nil {
		return hc.Fail(httpCtx, errBadRequest, statusBadRequest)
	}

	return hc.Ok(httpCtx, token)
}

func (hc *httpController) Download(httpCtx HTTPContext) error {
	token := domain.Token{Id: httpCtx.QueryGet(tokenQuery)}
	if err := hc.Validate.Struct(token); err != nil {
		hc.logger.Error(httpCtx.Context(), "input data validate error", err)
		return hc.Fail(httpCtx, errBadRequest, statusBadRequest)
	}

	file, err := hc.fileInteractor.Download(httpCtx.Context(), token)
	if err != nil {
		return hc.Fail(httpCtx, errBadRequest, statusBadRequest)
	}

	return hc.Ok(httpCtx, file)
}
