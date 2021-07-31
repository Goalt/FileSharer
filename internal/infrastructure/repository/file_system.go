package repository

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	errWrongRootDirFileMode = errors.New("wrong permission on root directory, should be 700")
	errWrongFileMode        = errors.New("wrong permission on file, should be 700")
	errRootDirNotExist      = errors.New("root directory is not exists")
	rootDirFileMode         = os.FileMode(0700)
	fileMode                = os.FileMode(0700)
)

type FileSystemRepository struct {
	rootPath string
}

func NewFileSystemRepository(rootPath string) (*FileSystemRepository, error) {
	if fileInfo, err := os.Stat(rootPath); os.IsNotExist(err) {
		return nil, errRootDirNotExist
	} else {
		if fileInfo.Mode() != rootDirFileMode {
			return nil, errWrongRootDirFileMode
		}
	}

	return &FileSystemRepository{
		rootPath: rootPath,
	}, nil
}

func (fr *FileSystemRepository) Read(fileName string) ([]byte, error) {
	fileInfo, err := os.Stat(fr.rootPath + fileName)
	if err != nil {
		return nil, err
	}

	if fileInfo.Mode() != fileMode {
		return nil, errWrongFileMode
	}

	data, err := ioutil.ReadFile(fr.rootPath + fileName)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (fr *FileSystemRepository) Write(fileName string, data []byte) error {
	_, err := os.Stat(fr.rootPath + fileName)
	if err == nil {
		return fmt.Errorf("file with name %v already exists", fileName)
	}
	if !os.IsNotExist(err) {
		return err
	}

	return ioutil.WriteFile(fileName, data, fileMode)
}
