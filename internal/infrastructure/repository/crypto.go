package repository

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type AESCrypto struct {
	cipher cipher.Block
}

func NewAESCrypto(key []byte) (*AESCrypto, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &AESCrypto{cipher}, nil
}

func (a *AESCrypto) Encrypt(data []byte) ([]byte, error) {
	aesGCM, err := cipher.NewGCM(a.cipher)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, data, nil)

	return ciphertext, nil
}

func (a *AESCrypto) Decrypt(data []byte) ([]byte, error) {
	aesGCM, err := cipher.NewGCM(a.cipher)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return nil, errors.New("invalid data's size")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	decrypted, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

func (a *AESCrypto) EncryptString(data string) (string, error) {
	dst, err := a.Encrypt([]byte(data))
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(dst), nil
}

func (a *AESCrypto) DecryptString(data string) (string, error) {
	dataRaw, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	dst, err := a.Decrypt(dataRaw)
	if err != nil {
		return "", err
	}

	return string(dst), nil
}
