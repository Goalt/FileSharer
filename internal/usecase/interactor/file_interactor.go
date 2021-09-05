package interactor

import (
	"context"
	"fmt"

	"github.com/Goalt/FileSharer/internal/domain"
	"github.com/Goalt/FileSharer/internal/errors"
	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
)

type FileInteractor interface {
	Upload(ctx context.Context, file domain.File) (domain.Token, error)
	Download(ctx context.Context, token domain.Token) (domain.File, error)
}

type fileInteractor struct {
	fileInfoRepository   usecase_repository.FileInfoRepository
	fileSystemRepository usecase_repository.FileSystemRepository
	cryptoInteractor     CryptoInteractor
	generatorInteractor  GeneratorInteractor
	logger               usecase_repository.Logger
}

func NewFileInteractor(
	fileInfoRepository usecase_repository.FileInfoRepository,
	fileSystemRepository usecase_repository.FileSystemRepository,
	cryptoInteractor CryptoInteractor,
	generatorInteractor GeneratorInteractor,
	logger usecase_repository.Logger,
) FileInteractor {
	return &fileInteractor{fileInfoRepository, fileSystemRepository, cryptoInteractor, generatorInteractor, logger}
}

func (ci *fileInteractor) Upload(ctx context.Context, file domain.File) (domain.Token, error) {
	encryptedFile, err := ci.cryptoInteractor.Encrypt(file)
	if err != nil {
		ci.logger.Error(ctx, fmt.Sprintf("failed during encrypting file %v", err))
		return domain.Token{}, errors.ErrUploadFile
	}

	fder, err := ci.cryptoInteractor.Decrypt(encryptedFile)
	if err != nil {
		ci.logger.Error(ctx, fmt.Sprintf("failed during encrypting file %v", err))
		return domain.Token{}, errors.ErrUploadFile
	}
	_ = fder

	token := ci.generatorInteractor.GenerateToken()
	fileName := ci.generatorInteractor.GenerateFileName()
	fileInfo := domain.FileInfo{
		Token:          token,
		FileName:       fileName,
		FileNameOrigin: encryptedFile.FileNameOrigin,
	}

	if err = ci.fileSystemRepository.Write(fileName, encryptedFile.Data); err != nil {
		ci.logger.Error(ctx, fmt.Sprintf("failed during saving file's data %v", err))
		return domain.Token{}, errors.ErrUploadFile
	}

	if err = ci.fileInfoRepository.Set(fileInfo); err != nil {
		ci.logger.Error(ctx, fmt.Sprintf("failed during saving file's info %v", err))

		if err = ci.fileSystemRepository.Delete(fileName); err != nil {
			ci.logger.Error(ctx, fmt.Sprintf("failed during deleting file %v", err))
		}

		return domain.Token{}, errors.ErrUploadFile
	}

	ci.logger.Info(ctx, fmt.Sprintf("file uploaded with token %v", token.Id))

	return domain.Token(fileInfo.Token), nil
}

func (ci *fileInteractor) Download(ctx context.Context, token domain.Token) (domain.File, error) {
	fileInfo, err := ci.fileInfoRepository.Get(token)
	if err != nil {
		ci.logger.Error(ctx, fmt.Sprintf("failed during searching file's info %v", err))
		return domain.File{}, errors.ErrFileNotFound
	}
	encryptedData, err := ci.fileSystemRepository.Read(fileInfo.FileName)
	if err != nil {
		ci.logger.Error(ctx, fmt.Sprintf("failed during reading file's data %v", err))
		return domain.File{}, errors.ErrDownloadFile
	}

	encryptedFile := domain.File{
		Data:           encryptedData,
		FileNameOrigin: fileInfo.FileNameOrigin,
	}
	decryptedFile, err := ci.cryptoInteractor.Decrypt(encryptedFile)
	if err != nil {
		ci.logger.Error(ctx, fmt.Sprintf("failed during decrypting file's data %v", err))
		return domain.File{}, errors.ErrDownloadFile
	}

	ci.logger.Info(ctx, fmt.Sprintf("file downloaded with token %v", token.Id))

	return decryptedFile, nil
}
