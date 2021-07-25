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

	for i, task := range Conf.Tasks {
		// Ensure task names are lower kebab-cased
		Conf.Tasks[i].Name = strings.ToLower(strings.ReplaceAll(task.Name, " ", "-"))
		if task.Count == 0 {
			// If no count is set default to 1
			Conf.Tasks[i].Count = 1
		}
	}

	log.WithField("Config", Conf).Debug("Loaded config")
}

func GetTaskDefinition(taskName string) (*task.Task, error) {
	for i := range Conf.Tasks {
		if Conf.Tasks[i].Name == taskName {
			return &Conf.Tasks[i], nil
		}
	}

	return nil, fmt.Errorf("couldn't find task with name '%v'", taskName)
}
