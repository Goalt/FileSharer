package controller

import (
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"strings"

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

	minFileNameLength = 16
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
		hc.logger.Error(httpCtx.Context(), fmt.Sprintf("parse media type error %v", err))
		return hc.Fail(httpCtx, errors.ErrFileFormat)
	}

	if !strings.HasPrefix(mediaType, multipartPrefix) {
		hc.logger.Error(httpCtx.Context(), "media type error")
		return hc.Fail(httpCtx, errors.ErrFileFormat)
	}

	body := httpCtx.BodyReader()
	multipartReader := multipart.NewReader(body, params[boundaryKey])

	part, err := multipartReader.NextPart()
	fmt.Print(part.Header.Get(fileNameHeader))
	switch {
	case err != nil:
		hc.logger.Error(httpCtx.Context(), fmt.Sprintf("multipart read error %v", err))
		return hc.Fail(httpCtx, errors.ErrFileFormat)
	}

	data := make([]byte, hc.maxFileSize+1)
	fileSize, err := part.Read(data)
	if fileSize == hc.maxFileSize+1 {
		hc.logger.Error(httpCtx.Context(), fmt.Sprintf("max file size %v", err))
		return hc.Fail(httpCtx, errors.ErrMaxFileSize)
	}
	if (err != nil) && !errors.Is(err, io.EOF) {
		hc.logger.Error(httpCtx.Context(), fmt.Sprintf("data read error %v", err))
		return hc.Fail(httpCtx, errors.ErrFileFormat)
	}

	file := domain.File{
		Data:           data,
		FileNameOrigin: part.FileName(),
	}
	if len(file.FileNameOrigin) < minFileNameLength {
		appendix := make([]byte, minFileNameLength-len(file.FileNameOrigin))
		file.FileNameOrigin += string(appendix)
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
	token := domain.Token{Id: httpCtx.QueryGet(tokenQuery)}
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
