package usecase_repository

import "github.com/Goalt/FileSharer/internal/domain"

type FileInfoRepository interface {
	Get(token domain.Token) (domain.FileInfo, error)
	Set(fileInfo domain.FileInfo) error
}
