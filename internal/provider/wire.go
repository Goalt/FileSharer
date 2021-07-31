//+build wireinject

package provider

import (
	"context"
	"github.com/Goalt/FileSharer/internal/config"
	"github.com/Goalt/FileSharer/internal/infrastructure/http"
	"github.com/google/wire"
	"gorm.io/gorm/logger"
)

type Application struct {
	ctx context.Context
	log logger.Interface

	server  http.Server
	cleanup ServicesCleanup
}

func (a *Application) Run() error {
	// Server start
	go func() {
		err := a.server.Run()
		if err != nil {
			a.log.Error(a.ctx, err.Error())
		}
	}()

	<-a.ctx.Done()

	//Server stop
	err := a.server.Stop()
	if err != nil {
		a.log.Error(a.ctx, err.Error())
	}

	return nil
}

func provideApp(server http.Server, cfg config.Config, ctx context.Context) Application {
	return Application{
		server: server,
		ctx:    ctx,
	}
}

func InitializeApp(cfg config.Config, context context.Context) (Application, func(), error) {
	panic(wire.Build(provideApp, infrastructureSet, interfaceSet, interactorSet))
}
