package internal

import (
	"strings"

	"github.com/angusgyoung/waiter/internal/task"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
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

	// Ensure task names are lower kebab-cased
	for i, task := range Conf.Tasks {
		Conf.Tasks[i].Name = strings.ToLower(strings.ReplaceAll(task.Name, " ", "-"))
	}

	Log.WithField("Config", Conf).Debug("Loaded config")
}
