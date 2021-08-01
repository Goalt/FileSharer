package provider

import (
	"github.com/Goalt/FileSharer/internal/usecase/interactor"
	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
	"github.com/google/wire"
)

func provideCalculatorInteractor(
	fileInfoRepository usecase_repository.FileInfoRepository,
	fileSystemRepository usecase_repository.FileSystemRepository,
	cryptoInteractor interactor.CryptoInteractor,
	generatorInteractor interactor.GeneratorInteractor,
	logger usecase_repository.Logger,
) interactor.FileInteractor {
	return interactor.NewFileInteractor(
		fileInfoRepository,
		fileSystemRepository,
		cryptoInteractor,
		generatorInteractor,
		logger,
	)
}

func cryptoInteractor(cryptoRepository usecase_repository.CryptoRepository) interactor.CryptoInteractor {
	return interactor.NewCryptoInteractor(cryptoRepository)
}

func provideGeneratorInteractor(uuidGenerator usecase_repository.UUIDGenerator) interactor.GeneratorInteractor {
	return interactor.NewGeneratorInteractor(uuidGenerator)
}

var interactorSet = wire.NewSet(provideCalculatorInteractor, cryptoInteractor, provideGeneratorInteractor)
