package file_sharer_migrations

import (
	"time"

	"github.com/Boostport/migration"
	"github.com/Boostport/migration/driver/mysql"
	"github.com/Goalt/FileSharer/cmd/subcomands"
	"github.com/Goalt/FileSharer/cmd/variables"
	"github.com/Goalt/FileSharer/internal/config"
	"github.com/Goalt/FileSharer/internal/migrations"
	"github.com/Goalt/FileSharer/internal/provider"
	"github.com/urfave/cli/v2"
)

func init() {
	subcommand := &cli.Command{
		Name:  "file_sharer_migrations",
		Usage: "file_sharer_migrations",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    variables.MysqlDatabaseName,
				EnvVars: []string{variables.MysqlDatabaseName},
			},
			&cli.StringFlag{
				Name:    variables.MysqlUser,
				EnvVars: []string{variables.MysqlUser},
			},
			&cli.StringFlag{
				Name:    variables.MysqlPassword,
				EnvVars: []string{variables.MysqlPassword},
			},
			&cli.StringFlag{
				Name:    variables.MysqlHost,
				EnvVars: []string{variables.MysqlHost},
			},
			&cli.StringFlag{
				Name:    variables.MysqlPort,
				EnvVars: []string{variables.MysqlPort},
			},
		},
		Action: func(ctx *cli.Context) error {
			logger := provider.ProvideLoggerGorm(4)

			configDB := config.Database{
				Host:     ctx.String(variables.MysqlHost),
				Port:     ctx.String(variables.MysqlPort),
				User:     ctx.String(variables.MysqlUser),
				Password: ctx.String(variables.MysqlPassword),
				DBName:   ctx.String(variables.MysqlDatabaseName),
			}

			embedSource := &migration.EmbedMigrationSource{
				EmbedFS: migrations.SQL,
				Dir:     "sql",
			}

			driver, err := mysql.New(configDB.GetDsn())
			if err != nil {
				logger.Error(ctx.Context, "migrations failed", err)
				return err
			}

			// Run all up migrations
			for {
				applied, err := migration.Migrate(driver, embedSource, migration.Up, 0)
				if err != nil {
					logger.Error(ctx.Context, "migrations failed", err)
					time.Sleep(time.Second * 5)
				} else {
					logger.Info(ctx.Context, "applied version", applied)
					break
				}
			}

			return nil
		},
	}

	subcomands.Add(subcommand)
}
