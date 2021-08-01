package config

import "fmt"

const FileInfoTableName = "file_info"
const DsnFormat = "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"

type Config struct {
	DebugLevel  int
	MaxFileSize int //Max file size in bytes
	RootPath    string
	Key         []byte
	Server      Server
	Database    Database
}

type Server struct {
	Port int
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (db *Database) GetDsn() string {
	return fmt.Sprintf(DsnFormat, db.User, db.Password, db.Host, db.Port, db.DBName)
}