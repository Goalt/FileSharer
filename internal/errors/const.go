package errors

import (
	"errors"
	"fmt"
	"net/http"
)

var MaxFileSize int

var (
	ErrFileFormat = HttpError{
		ResponseCode: http.StatusBadRequest,
		Text:         "file format error",
		ErrorCode:    1,
	}

	ErrMaxFileSize = HttpError{
		ResponseCode: http.StatusRequestEntityTooLarge,
		Text:         fmt.Sprintf("max file size bytes (%v bytes)", MaxFileSize),
		ErrorCode:    2,
	}

	ErrTokenFormat = HttpError{
		ResponseCode: http.StatusBadRequest,
		Text:         "token format error",
		ErrorCode:    3,
	}

	ErrUploadFile = HttpError{
		ResponseCode: http.StatusInternalServerError,
		Text:         "failed to download file",
		ErrorCode:    4,
	}

	ErrFileNotFound = HttpError{
		ResponseCode: http.StatusNotFound,
		Text:         "failed to find file",
		ErrorCode:    5,
	}

	ErrDownloadFile = HttpError{
		ResponseCode: http.StatusInternalServerError,
		Text:         "failed to download file",
		ErrorCode:    6,
	}
)

var (
	ErrIncorrectDataSize = errors.New("data size less aes.BlockSize")
)
