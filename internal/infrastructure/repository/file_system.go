package repository

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	errWrongRootDirFileMode = errors.New("wrong permission on root directory, 2147484141")
	errWrongFileMode        = errors.New("wrong permission on file, 2147484141")
	errRootDirNotExist      = errors.New("root directory is not exists")
	rootDirFileMode         = os.FileMode(2147484141)
	fileMode                = os.FileMode(493)
)

type FileSystemRepository struct {
	rootPath string
}

func NewFileSystemRepository(rootPath string) (*FileSystemRepository, error) {
	if fileInfo, err := os.Stat(rootPath); os.IsNotExist(err) {
		return nil, errRootDirNotExist
	} else {
		mode := fileInfo.Mode()
		if mode != rootDirFileMode {
			fmt.Println(fileInfo.Mode())
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

	mode := fileInfo.Mode()
	if mode != fileMode {
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

	return ioutil.WriteFile(fr.rootPath+fileName, data, fileMode)
}

func (fr *FileSystemRepository) Delete(fileName string) error {
	fileInfo, err := os.Stat(fr.rootPath + fileName)
	if err != nil {
		return err
	}

	if fileInfo.Mode() != fileMode {
		return errWrongFileMode
	}

	if err = os.Remove(fr.rootPath + fileName); err != nil {
		return err
	}

	return nil
}
