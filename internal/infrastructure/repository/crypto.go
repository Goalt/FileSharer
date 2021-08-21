package repository

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/Goalt/FileSharer/internal/errors"
)

type AESCrypto struct {
	cipher cipher.Block
	prefix []byte
}

func NewAESCrypto(key []byte) (*AESCrypto, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &AESCrypto{cipher, make([]byte, aes.BlockSize)}, nil
}

func (a *AESCrypto) Encrypt(data []byte) ([]byte, error) {
	data = append(data, a.prefix...)
	dst := make([]byte, len(data))
	a.cipher.Encrypt(dst, data)

	return dst, nil
}

func (a *AESCrypto) Decrypt(data []byte) ([]byte, error) {
	if len(data) < aes.BlockSize {
		return nil, errors.ErrIncorrectDataSize
	}

	dst := make([]byte, len(data))
	a.cipher.Decrypt(dst, data)

	return dst[:len(dst)-aes.BlockSize], nil
}

func (a *AESCrypto) EncryptString(data string) (string, error) {
	dst, err := a.Encrypt([]byte(data))
	if err != nil {
		return "", err
	}

	return string(dst), nil
}

func (a *AESCrypto) DecryptString(data string) (string, error) {
	dst, err := a.Decrypt([]byte(data))
	if err != nil {
		return "", err
	}

	return string(dst), nil
}
