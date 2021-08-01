package provider

import (
	"github.com/Goalt/FileSharer/internal/config"
	"github.com/google/wire"
)

type MaxFileSize int

func provideMaxFileSize(config config.Config) MaxFileSize {
	return MaxFileSize(config.MaxFileSize)
}

type DebugLevel int

func provideDebugLevel(config config.Config) DebugLevel {
	return DebugLevel(config.DebugLevel)
}

func provideServerConfig(config config.Config) config.Server {
	return config.Server
}

type RootPath string

func provideRootPath(config config.Config) RootPath {
	return RootPath(config.RootPath)
}

type Key []byte

func provideKey(config config.Config) Key {
	return config.Key
}

func provideDatabasesConfig(config config.Config) config.Database {
	return config.Database
}

var typesSet = wire.NewSet(
	provideMaxFileSize,
	provideDebugLevel,
	provideServerConfig,
	provideRootPath,
	provideKey,
	provideDatabasesConfig,
)
