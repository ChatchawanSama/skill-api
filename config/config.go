package config

import (
	"fmt"
	"os"
)

type Config struct {
	App     App     `mapstructure:"app" validate:"required"`
	Log     Log     `mapstructure:"log" validate:"required"`
	Server  Server  `mapstructure:"server" validate:"required"`
	Secrets Secrets `validate:"required"`
}

type App struct {
	Name      string `mapstructure:"name"`
	ProjectId string `mapstructure:"project_id"`
	Env       string `mapstructure:"env"`
}

type Log struct {
	Env   string `mapstructure:"env"`
	Level string `mapstructure:"level"`
}

type Server struct {
	Address  string `mapstructure:"address" validate:"required"`
	TimeZone string `mapstructure:"time_zone" validate:"required"`
}

type Secrets struct {
	MySQL struct {
		Host     string
		Port     string
		User     string
		Password string
	}
	Kafka struct {
		Address    string
		Username   string
		Password   string
		TLSCACert  string
	}
}

func New() (*Config, error) {
	required := []string{
		"SECRET_MYSQL_HOST",
		"SECRET_MYSQL_PORT",
		"SECRET_MYSQL_USER",
	}
	for _, k := range required {
		if os.Getenv(k) == "" {
			return nil, fmt.Errorf("env var %s is required", k)
		}
	}

	cfg := &Config{
		App:    App{Name: "loan-api", ProjectId: "loan-api", Env: "local"},
		Log:    Log{Env: "dev", Level: "info"},
		Server: Server{Address: ":8080", TimeZone: "Asia/Bangkok"},
	}
	cfg.Secrets.MySQL.Host = os.Getenv("SECRET_MYSQL_HOST")
	cfg.Secrets.MySQL.Port = os.Getenv("SECRET_MYSQL_PORT")
	cfg.Secrets.MySQL.User = os.Getenv("SECRET_MYSQL_USER")
	cfg.Secrets.MySQL.Password = os.Getenv("SECRET_MYSQL_PASSWORD")
	cfg.Secrets.Kafka.Address = os.Getenv("SECRET_KAFKA_ADDRESS")
	cfg.Secrets.Kafka.Username = os.Getenv("SECRET_KAFKA_USERNAME")
	cfg.Secrets.Kafka.Password = os.Getenv("SECRET_KAFKA_PASSWORD")
	cfg.Secrets.Kafka.TLSCACert = os.Getenv("SECRET_KAFKA_TLS_FILE_CA_CERT")
	return cfg, nil
}

func (s Secrets) MySQLDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/loan_db?parseTime=true&loc=Asia%%2FBangkok",
		s.MySQL.User, s.MySQL.Password, s.MySQL.Host, s.MySQL.Port,
	)
}
