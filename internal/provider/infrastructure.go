package provider

import (
	"log"
	"os"
	"time"

	"github.com/Goalt/FileSharer/internal/config"
	"github.com/Goalt/FileSharer/internal/infrastructure/http"
	infrastructure_repository "github.com/Goalt/FileSharer/internal/infrastructure/repository"
	"github.com/Goalt/FileSharer/internal/interface/controller"
	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServicesCleanup func()

func provideLoggerGorm(config config.Config) logger.Interface {
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
	return http.NewHTTPServer(config.Server.Port, controller)
}

func provideFileInfoRepository(db *gorm.DB) usecase_repository.FileInfoRepository {
	return infrastructure_repository.NewGormFileInfoRepository(db)
}

func provideFileSystemRepository(config config.Config) (usecase_repository.FileSystemRepository, error) {
	return infrastructure_repository.NewFileSystemRepository(config.RootPath)
}

func provideCryptoRepository(config config.Config) (usecase_repository.CryptoRepository, error) {
	return infrastructure_repository.NewAESCrypto(config.Key)
}

func provideUUIDGenerator() usecase_repository.UUIDGenerator {
	return infrastructure_repository.NewUUIDGenerator()
}

func provideGORM(config config.Config) (*gorm.DB, func()) {
	db, err := gorm.Open(mysql.Open(config.Database.GetDsn()), &gorm.Config{})
	if err != nil {
		return nil, nil
	}

	cleanup := func() {
		sqlDB, err := db.DB()
		if err != nil {
			return
		}

		_ = sqlDB.Close()
	}

	return db, cleanup
}

func provideLogger(log logger.Interface) usecase_repository.Logger {
	return log
}

func provideServicesCleanup(cleanup func()) ServicesCleanup {
	return cleanup
}

var infrastructureSet = wire.NewSet(
	provideServer,
	provideLoggerGorm,
	provideFileInfoRepository,
	provideFileSystemRepository,
	provideCryptoRepository,
	provideUUIDGenerator,
	provideGORM,
	provideLogger,
	provideServicesCleanup,
)
