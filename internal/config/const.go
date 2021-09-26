package config

import (
	gormLog "gorm.io/gorm/logger"
)

const (
	FileInfoTableName = "file_info"

	DsnFormat = "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"

	FilePath = "/api/file"

	InfoLevel  = 1
	WarnLevel  = 2
	ErrorLevel = 3

	GormLogLevel = gormLog.Error
)
