package interactor

import (
	"context"

	"github.com/Goalt/FileSharer/internal/domain"
	"github.com/Goalt/FileSharer/internal/errors"
	repository "github.com/Goalt/FileSharer/internal/usecase/repository"
)

type FileInteractor interface {
	Upload(ctx context.Context, file domain.File, log repository.Logger) (domain.Token, error)
	Download(ctx context.Context, token domain.Token, log repository.Logger) (domain.File, error)
}

type fileInteractor struct {
	fileInfoRepository   repository.FileInfoRepository
	fileSystemRepository repository.FileSystemRepository
	cryptoInteractor     CryptoInteractor
	generatorInteractor  GeneratorInteractor
}

func NewFileInteractor(
	fileInfoRepository repository.FileInfoRepository,
	fileSystemRepository repository.FileSystemRepository,
	cryptoInteractor CryptoInteractor,
	generatorInteractor GeneratorInteractor,
) *fileInteractor {
	return &fileInteractor{fileInfoRepository, fileSystemRepository, cryptoInteractor, generatorInteractor}
}

func (ci *fileInteractor) Upload(ctx context.Context, file domain.File, log repository.Logger) (domain.Token, error) {
	encryptedFile, err := ci.cryptoInteractor.Encrypt(file)
	if err != nil {
		log.Error(err)
		return domain.Token{}, errors.ErrUploadFile
	}

	token := ci.generatorInteractor.GenerateToken()
	fileName := ci.generatorInteractor.GenerateFileName()
	fileInfo := domain.FileInfo{
		Token:          token,
		FileName:       fileName,
		FileNameOrigin: encryptedFile.FileNameOrigin,
	}

	if err = ci.fileSystemRepository.Write(fileName, encryptedFile.Data); err != nil {
		log.Error(err)
		return domain.Token{}, errors.ErrUploadFile
	}

	if err = ci.fileInfoRepository.Set(fileInfo); err != nil {
		log.Error(err)

		if err = ci.fileSystemRepository.Delete(fileName); err != nil {
			log.Error(err)
		}

		return domain.Token{}, errors.ErrUploadFile
	}

	log.Infof("file uploaded with token %v", token.Id)

	return domain.Token(fileInfo.Token), nil
}

func (ci *fileInteractor) Download(ctx context.Context, token domain.Token, log repository.Logger) (domain.File, error) {
	fileInfo, err := ci.fileInfoRepository.Get(token)
	if err != nil {
		log.Error(err)
		return domain.File{}, errors.ErrFileNotFound
	}
	encryptedData, err := ci.fileSystemRepository.Read(fileInfo.FileName)
	if err != nil {
		log.Error(err)
		return domain.File{}, errors.ErrDownloadFile
	}

	encryptedFile := domain.File{
		Data:           encryptedData,
		FileNameOrigin: fileInfo.FileNameOrigin,
	}
	decryptedFile, err := ci.cryptoInteractor.Decrypt(encryptedFile)
	if err != nil {
		log.Error(err)
		return domain.File{}, errors.ErrDownloadFile
	}

	log.Infof("file downloaded with token %v", token.Id)

	return decryptedFile, nil
}
