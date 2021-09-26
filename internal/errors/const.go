package errors

import (
	"net/http"
)

var (
	ErrFileFormat = HttpError{
		ResponseCode: http.StatusBadRequest,
		Text:         "file format error",
		ErrorCode:    1,
	}

	ErrMaxFileSize = HttpError{
		ResponseCode: http.StatusRequestEntityTooLarge,
		Text:         "max file size bytes",
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
