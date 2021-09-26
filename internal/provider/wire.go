//+build wireinject

package provider

import (
	"context"

	"github.com/Goalt/FileSharer/internal/config"
	"github.com/Goalt/FileSharer/internal/infrastructure/http"
	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
	"github.com/google/wire"
)

type Application struct {
	ctx context.Context
	log usecase_repository.Logger

	server http.Server
	config config.Config
}

func (a *Application) Run() error {
	// Server start
	go func() {
		err := a.server.Run()
		if err != nil {
			a.log.Error(err)
		}
	}()

	<-a.ctx.Done()

	//Server stop
	err := a.server.Stop()
	if err != nil {
		a.log.Error(err)
	}

	return nil
}

func provideApp(server http.Server, cfg config.Config, ctx context.Context, log usecase_repository.Logger) Application {
	return Application{
		server: server,
		ctx:    ctx,
		config: cfg,
		log:    log,
	}
}

func InitializeApp(cfg config.Config, context context.Context) (Application, func(), error) {
	panic(wire.Build(provideApp, infrastructureSet, interfaceSet, interactorSet, typesSet))
}
