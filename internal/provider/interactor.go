package provider

import (
	"github.com/Goalt/FileSharer/internal/usecase/interactor"
	"github.com/google/wire"
)

func provideCalculatorInteractor() interactor.CalculatorInteractor {
	return interactor.NewCalculatorInteractor()
}

var interactorSet = wire.NewSet(provideCalculatorInteractor)
