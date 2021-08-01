package file_sharer_migrations

import (
	"github.com/Boostport/migration"
	"github.com/Boostport/migration/driver/mysql"
	"github.com/Goalt/FileSharer/cmd/subcomands"
	"github.com/Goalt/FileSharer/internal/config"
	"github.com/Goalt/FileSharer/internal/migrations"
	"github.com/Goalt/FileSharer/internal/provider"
	"github.com/urfave/cli/v2"
)

var (
	MysqlDatabaseName = "MYSQL_DATABASE"
	MysqlUser         = "MYSQL_USER"
	MysqlPassword     = "MYSQL_PASSWORD"
	MysqlHost         = "MYSQL_HOST"
	MysqlPort         = "MYSQL_PORT"
)

func init() {
	subcommand := &cli.Command{
		Name:  "filesharer_migrations",
		Usage: "filesharer_migrations",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    MysqlDatabaseName,
				EnvVars: []string{MysqlDatabaseName},
			},
			&cli.StringFlag{
				Name:    MysqlUser,
				EnvVars: []string{MysqlUser},
			},
			&cli.StringFlag{
				Name:    MysqlPassword,
				EnvVars: []string{MysqlPassword},
			},
			&cli.StringFlag{
				Name:    MysqlHost,
				EnvVars: []string{MysqlHost},
			},
			&cli.StringFlag{
				Name:    MysqlPort,
				EnvVars: []string{MysqlPort},
			},
		},
		Action: func(ctx *cli.Context) error {
			logger := provider.ProvideLoggerGorm(4)

			configDB := config.Database{
				Host:     ctx.String(MysqlHost),
				Port:     ctx.String(MysqlPort),
				User:     ctx.String(MysqlUser),
				Password: ctx.String(MysqlPassword),
				DBName:   ctx.String(MysqlDatabaseName),
			}

			embedSource := &migration.EmbedMigrationSource{
				EmbedFS: migrations.SQL,
				Dir:     "sql",
			}

			driver, err := mysql.New(configDB.GetDsn())

			// Run all up migrations
			applied, err := migration.Migrate(driver, embedSource, migration.Up, 0)
			if err != nil {
				logger.Error(ctx.Context, "migrations failed", err)
			} else {
				logger.Info(ctx.Context, "applied version", applied)
			}

			return nil
		},
	}

	subcomands.Add(subcommand)
}
