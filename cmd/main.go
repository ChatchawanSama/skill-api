package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fusic/skill-api/internal/pkg/configurer"
	"github.com/fusic/skill-api/internal/pkg/logger"
	"github.com/fusic/skill-api/internal/pkg/mysql_database"
	"go.uber.org/zap"
)

func main() {
	// create root context
	ctx, _ := context.WithCancel(context.Background())

	// load config
	conf := configurer.LoadConfig()

	// load secret env
	configurer.LoadSecret(&conf.Secret)

	fmt.Println("Config loaded:", conf)

	// database
	db, err := mysql_database.Connect(&mysql_database.MysqlConfig{
		Name:                conf.DB.Name,
		SSLMode:             conf.DB.SSLMode,
		MaxOpenConns:        &conf.DB.MaxOpenConnection,
		MaxIdleConns:        &conf.DB.MaxIdleConnection,
		ConnMaxLifetimeHour: conf.DB.ConnectionMaxLifetime,
		MysqlHost:           conf.Secret.DBHost,
		MysqlPort:           conf.Secret.DBPort,
		MysqlUser:           conf.Secret.DBUser,
		MysqlPassword:       conf.Secret.DBPassword,
		Loc:                 time.Local,
	})

	if err != nil {
		logger.Fatal(ctx, "fail connect database", zap.Error(err))
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			logger.Warn(ctx, "fail get sql.DB", zap.Error(err))
			return
		}
		if err := sqlDB.Close(); err != nil {
			logger.Warn(ctx, "fail close database", zap.Error(err))
		}
	}()

	// TODO : Init repositories, services, handlers, and start the HTTP server.
}
