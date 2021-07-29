package provider

import (
	"github.com/Goalt/FileSharer/internal/interface/controller"
	"github.com/Goalt/FileSharer/internal/usecase/interactor"
	"github.com/google/wire"
)

func provideHTTPController(interactor interactor.CalculatorInteractor) controller.HTTPController {
	return controller.NewHTTPController(interactor)
}

var interfaceSet = wire.NewSet(provideHTTPController)
