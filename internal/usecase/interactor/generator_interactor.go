package interactor

import (
	"github.com/Goalt/FileSharer/internal/domain"
	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
)

type GeneratorInteractor interface {
	GenerateToken() domain.Token
	GenerateFileName() string
}

type generatorInteractor struct {
	uuidGenerator usecase_repository.UUIDGenerator
}

func NewGeneratorInteractor(uuidGenerator usecase_repository.UUIDGenerator) *generatorInteractor {
	return &generatorInteractor{uuidGenerator}
}

func (gi *generatorInteractor) GenerateToken() domain.Token {
	return domain.Token{
		Id: gi.uuidGenerator.GetUUID(),
	}
}

func (gi *generatorInteractor) GenerateFileName() string {
	return gi.uuidGenerator.GetUUID()
}
