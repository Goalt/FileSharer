package provider

import (
	"github.com/Goalt/FileSharer/internal/interface/controller"
	"github.com/Goalt/FileSharer/internal/usecase/interactor"
	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
	"github.com/google/wire"
)

func provideHTTPController(
	maxFileSize MaxFileSize,
	interactor interactor.FileInteractor,
	base64Repository usecase_repository.Base64Repository,
	log usecase_repository.Logger,
) controller.HTTPController {
	return controller.NewHTTPController(int(maxFileSize), interactor, log, base64Repository)
}

var interfaceSet = wire.NewSet(provideHTTPController)
