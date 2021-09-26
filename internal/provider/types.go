package provider

import (
	"github.com/Goalt/FileSharer/internal/config"
	"github.com/google/wire"
)

type MaxFileSize int
type Key []byte
type RootPath string

func provideMaxFileSize(config config.Config) MaxFileSize {
	return MaxFileSize(config.MaxFileSize)
}

func provideCnfLogger(config config.Config) config.Logger {
	return config.Logger
}

func provideServerConfig(config config.Config) config.Server {
	return config.Server
}

func provideRootPath(config config.Config) RootPath {
	return RootPath(config.RootPath)
}

func provideKey(config config.Config) Key {
	return config.Key
}

func provideDatabasesConfig(config config.Config) config.Database {
	return config.Database
}

var typesSet = wire.NewSet(
	provideMaxFileSize,
	provideServerConfig,
	provideRootPath,
	provideKey,
	provideDatabasesConfig,
	provideCnfLogger,
)
