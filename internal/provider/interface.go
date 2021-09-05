package provider

import (
	"github.com/Goalt/FileSharer/internal/interface/controller"
	"github.com/Goalt/FileSharer/internal/usecase/interactor"
	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
	"github.com/google/wire"
)

func provideHTTPController(maxFileSize MaxFileSize, interactor interactor.FileInteractor, log usecase_repository.Logger) controller.HTTPController {
	return controller.NewHTTPController(int(maxFileSize), interactor, log)
}

var interfaceSet = wire.NewSet(provideHTTPController)
