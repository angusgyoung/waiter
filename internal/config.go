package internal

import (
	"fmt"

	"github.com/angusgyoung/waiter/internal/task"
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
		fmt.Println("Failed to unmarshal config", err)
	}
}
