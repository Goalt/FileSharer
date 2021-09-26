package provider

import (
	"context"
	"net"
	"time"

	"github.com/Goalt/FileSharer/internal/config"
	gorm_repository "github.com/Goalt/FileSharer/internal/infrastructure/gorm"
	"github.com/Goalt/FileSharer/internal/infrastructure/http"
	"github.com/Goalt/FileSharer/internal/infrastructure/logger"
	infrastructure_repository "github.com/Goalt/FileSharer/internal/infrastructure/repository"
	"github.com/Goalt/FileSharer/internal/interface/controller"
	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
)

type ServicesCleanup func()

func provideServer(cfg config.Server, controller controller.HTTPController) http.Server {
	return http.NewHTTPServer(cfg.Port, controller)
}

func provideFileInfoRepository(db *gorm.DB) usecase_repository.FileInfoRepository {
	return gorm_repository.NewGormFileInfoRepository(db)
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

func provideBase64Repository() usecase_repository.Base64Repository {
	return infrastructure_repository.NewBase64Repository()
}

func ProvideGORM(cnf config.Database, ctx context.Context, log usecase_repository.Logger) (*gorm.DB, func()) {
	gormLog := gormLog.New(
		log,
		gormLog.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  config.GormLogLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	var err error
	var db *gorm.DB
	for {
		db, err = gorm.Open(mysql.Open(cnf.GetDsn()), &gorm.Config{Logger: gormLog})
		if _, ok := err.(*net.OpError); ok {
			log.Info("db unavailable, sleep for 5 seconds")
			time.Sleep(time.Second * 5)
			continue
		}

		break
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

func ProvideLogger(cnf config.Logger) usecase_repository.Logger {
	return logger.NewLogger(cnf)
}

func provideServicesCleanup(cleanup func()) ServicesCleanup {
	return cleanup
}

var infrastructureSet = wire.NewSet(
	provideServer,
	provideFileInfoRepository,
	provideFileSystemRepository,
	provideCryptoRepository,
	provideUUIDGenerator,
	ProvideGORM,
	ProvideLogger,
	provideServicesCleanup,
	provideBase64Repository,
)
