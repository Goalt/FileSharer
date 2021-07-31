package repository

import (
	"time"

	"github.com/Goalt/FileSharer/internal/config"
	"github.com/Goalt/FileSharer/internal/domain"
	"gorm.io/gorm"
)

type GormfileInfoRepository struct {
	db *gorm.DB
}

func NewGormFileInfoRepository(db *gorm.DB) *GormfileInfoRepository {
	return &GormfileInfoRepository{db}
}

func (gr *GormfileInfoRepository) Get(token domain.Token) (domain.FileInfo, error) {
	dbModel := fileInfoDBModel{
		TokenId: token.Id,
	}

	err := gr.db.First(&dbModel).Error
	if err != nil {
		return domain.FileInfo{}, err
	}

	return domain.FileInfo{
		Token: domain.Token{
			Id: dbModel.TokenId,
		},
		FileName:       dbModel.FileName,
		FileNameOrigin: dbModel.FileNameOrigin,
		CreatedAt:      dbModel.CreatedAt,
	}, nil

}

func (gr *GormfileInfoRepository) Set(fileInfo domain.FileInfo) error {
	dbModel := fileInfoDBModel{
		TokenId:        fileInfo.Token.Id,
		FileName:       fileInfo.FileName,
		FileNameOrigin: fileInfo.FileNameOrigin,
		CreatedAt:      fileInfo.CreatedAt,
	}

	err := gr.db.Save(&dbModel).Error
	if err != nil {
		return err
	}

	return nil
}

type fileInfoDBModel struct {
	TokenId        string
	FileName       string
	FileNameOrigin string
	CreatedAt      time.Time
}

func (fi *fileInfoDBModel) Tabler() string {
	return config.FileInfoTableName
}
