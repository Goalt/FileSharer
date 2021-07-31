package controller

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/Goalt/FileSharer/internal/domain"
	"github.com/Goalt/FileSharer/internal/usecase/interactor"
	"github.com/Goalt/FileSharer/internal/usecase/repository"
	"github.com/go-playground/validator"
)

var (
	contetTypeHeader = "Content-Type"
	multipartPrefix  = "multipart/"
	boundaryKey      = "boundary"
	fileNameHeader   = "filename"
	tokenQuery       = "token_id"

	ErrBadRequest = errors.New("bad request")
)

type HTTPController interface {
	Upload(httpCtx HTTPContext) error
	Download(httpCtx HTTPContext) error
}

type httpController struct {
	fileInteractor interactor.FileInteractor
	*validator.Validate
	logger repository.Logger
}

func NewHTTPController(fileInteractor interactor.FileInteractor, logger repository.Logger) HTTPController {
	return &httpController{fileInteractor, validator.New(), logger}
}

func (hc *httpController) Upload(httpCtx HTTPContext) error {
	mediaType, params, err := mime.ParseMediaType(httpCtx.HeaderGet(contetTypeHeader))
	if err != nil {
		hc.logger.Error(httpCtx.Context(), "parse media type error", err)
		return mapError(httpCtx, ErrBadRequest)
	}

	if !strings.HasPrefix(mediaType, multipartPrefix) {
		hc.logger.Error(httpCtx.Context(), "media type error", nil)
		return mapError(httpCtx, ErrBadRequest)
	}

	body := httpCtx.BodyReader()
	multipartReader := multipart.NewReader(body, params[boundaryKey])

	part, err := multipartReader.NextPart()
	fmt.Print(part.Header.Get(fileNameHeader))
	switch {
	case err != nil:
		hc.logger.Error(httpCtx.Context(), "multipart read error", err)
		return mapError(httpCtx, ErrBadRequest)
	}

	data, err := io.ReadAll(part)
	if err != nil {
		hc.logger.Error(httpCtx.Context(), "data read error", err)
		return mapError(httpCtx, ErrBadRequest)
	}

	file := &domain.File{
		Data:           data,
		FileNameOrigin: part.FileName(),
	}
	if err := hc.Validate.Struct(file); err != nil {
		hc.logger.Error(httpCtx.Context(), "input data validate error", err)
		return mapError(httpCtx, ErrBadRequest)
	}

	token, err := hc.fileInteractor.Upload(httpCtx.Context(), file)
	if err != nil {
		return mapError(httpCtx, err)
	}

	return httpCtx.JSON(http.StatusOK, token)
}

func (hc *httpController) Download(httpCtx HTTPContext) error {
	token := domain.Token{Id: httpCtx.QueryGet(tokenQuery)}
	if err := hc.Validate.Struct(token); err != nil {
		hc.logger.Error(httpCtx.Context(), "input data validate error", err)
		return mapError(httpCtx, ErrBadRequest)
	}

	file, err := hc.fileInteractor.Download(httpCtx.Context(), token)
	if err != nil {
		return mapError(httpCtx, err)
	}

	return httpCtx.JSON(http.StatusOK, file)
}

type httpError struct {
	Text string `json:"text"`
}

func mapError(httpCtx HTTPContext, err error) error {
	switch {
	case errors.Is(err, ErrBadRequest):
		fallthrough
	default:
		return httpCtx.JSON(http.StatusBadRequest, httpError{Text: err.Error()})
	}
}
