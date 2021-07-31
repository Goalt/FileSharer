package interactor

import (
	"context"
	"errors"

	"github.com/Goalt/FileSharer/internal/domain"
	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
)

var (
	ErrSaveFile     = errors.New("failed to save file")
	ErrSaveFileInfo = errors.New("failed to save file's info")
	ErrUploadFail   = errors.New("failed to upload file")
	ErrDownloadFail = errors.New("failed to download file")
	ErrFindFile     = errors.New("failed to find file")
	ErrFindFileInfo = errors.New("failed to find file's info")
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
		ci.logger.Error(ctx, "failed during encrypting file", err)
		return domain.Token{}, ErrUploadFail
	}

	token := ci.generatorInteractor.GenerateToken()
	fileName := ci.generatorInteractor.GenerateFileName()
	fileInfo := domain.FileInfo{
		Token:          token,
		FileName:       fileName,
		FileNameOrigin: encryptedFile.FileNameOrigin,
	}

	err = ci.fileSystemRepository.Write(fileName, encryptedFile.Data)
	if err != nil {
		ci.logger.Error(ctx, "failed during saving file's data", err)
		return domain.Token{}, ErrSaveFile
	}
	err = ci.fileInfoRepository.Set(fileInfo)
	if err != nil {
		ci.logger.Error(ctx, "failed during saving file's info", err)
		return domain.Token{}, ErrSaveFileInfo
	}

	ci.logger.Info(ctx, "file uploaded with token "+token.Id)

	return domain.Token(fileInfo.Token), nil
}

func (ci *fileInteractor) Download(ctx context.Context, token domain.Token) (domain.File, error) {
	fileInfo, err := ci.fileInfoRepository.Get(token)
	if err != nil {
		ci.logger.Error(ctx, "failed during searching file's info", err)
		return domain.File{}, ErrFindFileInfo
	}
	encryptedData, err := ci.fileSystemRepository.Read(fileInfo.FileName)
	if err != nil {
		ci.logger.Error(ctx, "failed during reading file's data", err)
		return domain.File{}, ErrFindFile
	}

	encryptedFile := domain.File{
		Data:           encryptedData,
		FileNameOrigin: fileInfo.FileNameOrigin,
	}
	decryptedFile, err := ci.cryptoInteractor.Decrypt(encryptedFile)
	if err != nil {
		ci.logger.Error(ctx, "failed during decrypting file's data", err)
		return domain.File{}, ErrDownloadFail
	}

	ci.logger.Info(ctx, "file downloaded with token "+token.Id)

	return decryptedFile, nil
}
