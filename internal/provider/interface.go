package provider

import (
	"github.com/Goalt/FileSharer/internal/interface/controller"
	"github.com/Goalt/FileSharer/internal/usecase/interactor"
	"github.com/google/wire"
	"gorm.io/gorm/logger"
)

func provideHTTPController(maxFileSize MaxFileSize, interactor interactor.FileInteractor, log logger.Interface) controller.HTTPController {
	return controller.NewHTTPController(int(maxFileSize), interactor, log)
}

var interfaceSet = wire.NewSet(provideHTTPController)
