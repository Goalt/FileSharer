package provider

import (
	"github.com/Goalt/FileSharer/internal/config"
	"github.com/Goalt/FileSharer/internal/infrastructure/http"
	"github.com/Goalt/FileSharer/internal/interface/controller"
	"github.com/google/wire"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func provideLogger(config config.Config) logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.LogLevel(config.DebugLevel),
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
}

func provideServer(config config.Config, controller controller.HTTPController) http.Server {
	return http.NewHTTPServer(controller)
}

var infrastructureSet = wire.NewSet(provideServer, provideLogger)
