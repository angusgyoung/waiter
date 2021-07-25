package internal

import (
	"fmt"
	"strings"

	"github.com/angusgyoung/waiter/internal/task"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Tasks []task.Task `yaml:"tasks"`
}

var Conf Config

func LoadConfig() {
	err := viper.Unmarshal(&Conf)

	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
			"Path":  viper.ConfigFileUsed(),
		}).Fatal("Failed to unmarshall config file")
	}

	// Ensure task names are lower kebab-cased
	for i, task := range Conf.Tasks {
		Conf.Tasks[i].Name = strings.ToLower(strings.ReplaceAll(task.Name, " ", "-"))
	}

	log.WithField("Config", Conf).Debug("Loaded config")
}

func GetTaskDefinition(taskName string) (*task.Task, error) {
	for i := range Conf.Tasks {
		if Conf.Tasks[i].Name == taskName {
			return &Conf.Tasks[i], nil
		}
	}

	return nil, fmt.Errorf("couldn't find task with name \"%v\"", taskName)
}
