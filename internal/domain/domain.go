package domain

import "time"

type Token struct {
	Id string `json:"token_id" validate:"required"`
}

type File struct {
	Data           []byte `json:"data" validate:"required"`
	FileNameOrigin string `json:"file_name" validate:"required"`
}

type FileInfo struct {
	Token          Token
	FileName       string
	FileNameOrigin string
	CreatedAt      time.Time
}
