package config

import (
	"flag"

	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

type Args struct {
	Address       string
	DatabaseURL   string
	SystemAddress string
}

type Param struct {
	Address       string `env:"RUN_ADDRESS"`
	DatabaseURL   string `env:"DATABASE_URI"`
	SystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func loadConfig() (p Param, err error) {
	var parametrs Param
	err = env.Parse(&parametrs)
	if err != nil {
		logrus.Error("Error parse env: ", err)
	}
	return parametrs, nil
}

func TerminateFlags() (a Args, err error) {
	env, err := loadConfig()
	if err != nil {
		logrus.Error("Error terminate env or flags: ", err)
	}
	var flags Args
	var args Args
	flag.StringVar(&flags.Address, "a", "127.0.0.1:8080", "Address of server")
	flag.StringVar(&flags.DatabaseURL, "d", "postgres://user:user@127.0.0.1:5432/db", "Database URL")
	flag.StringVar(&flags.SystemAddress, "r", "127.0.0.1:8090", "Address of system accrual")
	if env.Address == "" {
		args.Address = flags.Address
	} else {
		args.Address = env.Address
	}
	if env.DatabaseURL == "" {
		args.DatabaseURL = flags.DatabaseURL
	} else {
		args.DatabaseURL = flags.DatabaseURL
	}
	if env.SystemAddress == "" {
		args.SystemAddress = flags.SystemAddress
	} else {
		args.SystemAddress = flags.SystemAddress
	}
	return args, nil
}
