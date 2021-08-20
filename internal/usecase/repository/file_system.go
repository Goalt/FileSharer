package usecase_repository

type FileSystemRepository interface {
	Read(fileName string) ([]byte, error)
	Write(fileName string, data []byte) error
	Delete(fileName string) error
}
