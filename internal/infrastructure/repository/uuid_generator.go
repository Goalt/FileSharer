package repository

import (
	"github.com/google/uuid"
)

type UUIDGenerator struct {
}

func NewUUIDGenerator() *UUIDGenerator {
	return &UUIDGenerator{}
}

func (ug *UUIDGenerator) GetUUID() string {
	// TODO обработать панику
	return uuid.NewString()
}
