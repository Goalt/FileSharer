package errors

import "errors"

type HttpError struct {
	ResponseCode int    `json:"-"`
	ErrorCode    int    `json:"error_code"`
	Text         string `json:"text"`
}

func (e HttpError) String() string {
	return e.Text
}

func (e HttpError) Error() string {
	return e.Text
}

func New(s string) error {
	return errors.New(s)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}
