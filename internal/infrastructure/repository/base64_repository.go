package repository

import "encoding/base64"

type Base64Repository struct {
}

func NewBase64Repository() *Base64Repository {
	return &Base64Repository{}
}

func (br *Base64Repository) Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func (br *Base64Repository) Decode(s string) (string, error) {
	sArray, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(sArray), nil
}
