package interactor

import (
	"github.com/Goalt/FileSharer/internal/domain"
	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
)

type CryptoInteractor interface {
	Encrypt(file domain.File) (domain.File, error)
	Decrypt(file domain.File) (domain.File, error)
}

type cryptoInteractor struct {
	crypto usecase_repository.CryptoRepository
}

func NewCryptoInteractor(crypto usecase_repository.CryptoRepository) CryptoInteractor {
	return &cryptoInteractor{crypto}
}

func (ci *cryptoInteractor) Encrypt(file domain.File) (domain.File, error) {
	encryptedData, err := ci.crypto.Encrypt(file.Data)
	if err != nil {
		return domain.File{}, err
	}

	encryptedFileNameOrigin, err := ci.crypto.EncryptString(file.FileNameOrigin)
	if err != nil {
		return domain.File{}, err
	}

	return domain.File{
		Data:           encryptedData,
		FileNameOrigin: encryptedFileNameOrigin,
	}, nil
}

func (ci *cryptoInteractor) Decrypt(file domain.File) (domain.File, error) {
	decryptedData, err := ci.crypto.Decrypt(file.Data)
	if err != nil {
		return domain.File{}, err
	}

	decryptedFileNameOrigin, err := ci.crypto.DecryptString(file.FileNameOrigin)
	if err != nil {
		return domain.File{}, err
	}

	return domain.File{
		Data:           decryptedData,
		FileNameOrigin: decryptedFileNameOrigin,
	}, nil
}
