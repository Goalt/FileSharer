package errors

import (
	"net/http"
	"testing"
)

var testError = HttpError{
	ResponseCode: http.StatusInternalServerError,
	ErrorCode:    1,
	Text:         "test text",
}

func TestHttpError_String(t *testing.T) {
	tests := []struct {
		name string
		err  HttpError
		want string
	}{
		{
			name: "simple test",
			err:  testError,
			want: "test text",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.String(); got != tt.want {
				t.Errorf("HttpError.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHttpError_Error(t *testing.T) {
	tests := []struct {
		name   string
		err  HttpError
		want   string
	}{
		{
			name: "simple test",
			err:  testError,
			want: "test text",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("HttpError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}