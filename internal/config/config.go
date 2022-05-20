package config

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

type Param struct {
	Address string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
}

func LoadConfig() (p Param, err error) {
	var parametrs Param
	err = env.Parse(&parametrs)
	if err != nil {
		logrus.Error("Error parse env: ", err)
	}
	return parametrs, nil
}
