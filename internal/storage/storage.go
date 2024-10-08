package storage

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/npavlov/storage-go/internal/file"
)

type Storage struct {
	Files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	return newFile, nil
}

func (s *Storage) GetById(fileId uuid.UUID) (*file.File, error) {
	foundFile, ok := s.Files[fileId]

	if !ok {
		return nil, fmt.Errorf("file %v not found", fileId)
	}
	return foundFile, nil
}
