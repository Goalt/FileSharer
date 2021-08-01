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

func ProvideLoggerGorm(level DebugLevel) logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.LogLevel(level),
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
}

func provideServer(cfg config.Server, controller controller.HTTPController) http.Server {
	return http.NewHTTPServer(cfg.Port, controller)
}

func provideFileInfoRepository(db *gorm.DB) usecase_repository.FileInfoRepository {
	return infrastructure_repository.NewGormFileInfoRepository(db)
}

func provideFileSystemRepository(rootPath RootPath) (usecase_repository.FileSystemRepository, error) {
	return infrastructure_repository.NewFileSystemRepository(string(rootPath))
}

func provideCryptoRepository(key Key) (usecase_repository.CryptoRepository, error) {
	return infrastructure_repository.NewAESCrypto(key)
}

func provideUUIDGenerator() usecase_repository.UUIDGenerator {
	return infrastructure_repository.NewUUIDGenerator()
}

func ProvideGORM(config config.Database) (*gorm.DB, func()) {
	db, err := gorm.Open(mysql.Open(config.GetDsn()), &gorm.Config{})
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
	ProvideLoggerGorm,
	provideFileInfoRepository,
	provideFileSystemRepository,
	provideCryptoRepository,
	provideUUIDGenerator,
	ProvideGORM,
	provideLogger,
	provideServicesCleanup,
)
