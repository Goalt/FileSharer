package provider

import (
	"github.com/Goalt/FileSharer/internal/config"
	"github.com/Goalt/FileSharer/internal/interface/controller"
	"github.com/Goalt/FileSharer/internal/usecase/interactor"
	"github.com/google/wire"
	"gorm.io/gorm/logger"
)

func provideHTTPController(config config.Config, interactor interactor.FileInteractor, log logger.Interface) controller.HTTPController {
	return controller.NewHTTPController(config.MaxFileSize, interactor, log)
}

var interfaceSet = wire.NewSet(provideHTTPController)
