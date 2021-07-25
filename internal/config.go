package internal

import (
	"github.com/angusgyoung/waiter/internal/task"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Test  string      `yaml:"test"`
	Tasks []task.Task `yaml:"tasks"`
}

var Conf Config

func LoadConfig() {
	err := viper.Unmarshal(&Conf)

	if err != nil {
		Log.WithFields(logrus.Fields{
			"Error": err,
			"Path":  viper.ConfigFileUsed(),
		}).Fatal("Failed to unmarshall config file")
	}

	Log.WithField("Config", Conf).Debug("Loaded config")
}
