package repository

import (
	"time"

	"github.com/Goalt/FileSharer/internal/config"
	"github.com/Goalt/FileSharer/internal/domain"
	"gorm.io/gorm"
)

type GormFileInfoRepository struct {
	db *gorm.DB
}

func NewGormFileInfoRepository(db *gorm.DB) *GormFileInfoRepository {
	return &GormFileInfoRepository{db}
}

func (gr *GormFileInfoRepository) Get(token domain.Token) (domain.FileInfo, error) {
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

func (gr *GormFileInfoRepository) Set(fileInfo domain.FileInfo) error {
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
	TokenId        string    `gorm:"primaryKey"`
	FileName       string    `gorm:"column:file_name"`
	FileNameOrigin string    `gorm:"column:file_name_origin"`
	CreatedAt      time.Time `gorm:"column:created_at"`
}

func (fi *fileInfoDBModel) TableName() string {
	return config.FileInfoTableName
}
