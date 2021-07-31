package provider

import (
	"github.com/Goalt/FileSharer/internal/usecase/interactor"
	"github.com/google/wire"
)

func provideCalculatorInteractor() interactor.FileInteractor {
	return interactor.NewFileInteractor()
}

var interactorSet = wire.NewSet(provideCalculatorInteractor)
