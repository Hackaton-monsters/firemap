package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBUser          string `envconfig:"db_user" default:"docker"`
	DBPassword      string `envconfig:"db_password" default:"secret"`
	DBHost          string `envconfig:"db_host" default:"host.docker.internal"`
	DBPort          string `envconfig:"db_port" default:"3306"`
	DBName          string `envconfig:"db_name" default:"firemap"`
	MaxOpenConn     int    `envconfig:"db_pool_max_conn" default:"10"`
	MaxIdleConn     int    `envconfig:"db_pool_max_idle" default:"5"`
	ConnMaxLifetime int    `envconfig:"db_pool_conn_lifetime" default:"10"`
	HttpServerPort  int    `envconfig:"HTTP_SERVER_PORT" default:"8081"`
}

func LoadFromEnvironment() *Config {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		panic(err)
	}

	return &config
}
