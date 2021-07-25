package task

import (
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Task struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
}

func (t Task) Execute() ([]byte, error) {
	log.WithFields(log.Fields{
		"Name":    t.Name,
		"Command": t.Command,
	}).Debug("Running task")

	var components []string = strings.Fields(t.Command)

	// Assumes the first component of the command is the binary to execute, followed
	// by any arguments or flags
	name := components[0]
	args := components[1:]

	command := exec.Command(name, args...)
	return command.Output()
}
