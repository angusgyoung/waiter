package task

import (
	"os"
	"os/exec"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type Task struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Output  string `yaml:"output"`
	Count   int    `yaml:"count"`
	Delay   int    `yaml:"delay"`
}

func (t Task) Execute() error {
	log.WithFields(log.Fields{
		"Name":    t.Name,
		"Command": t.Command,
	}).Debug("Running task")

	var components []string = strings.Fields(t.Command)

	// Assumes the first component of the command is the binary to execute, followed
	// by any arguments or flags
	name := components[0]
	args := components[1:]

	var err error

	for i := 0; i < t.Count; i++ {
		command := exec.Command(name, args...)

		out, err := command.Output()

		// Terminate execution loop on error.
		// In future we might allow continue-on-error
		if err != nil {
			break
		}

		// We must be writing to a file
		if t.Output != "" {
			log.WithField("Filename", t.Output).Trace("Writing task output to file")

			file, err := os.OpenFile(t.Output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			checkFileWriteErr(t.Output, err)
			defer file.Close()

			_, err = file.Write(out)
			checkFileWriteErr(t.Output, err)
		}

		// Don't delay if we are on our final execution
		if t.Delay != 0 && i != t.Count-1 {
			time.Sleep(time.Duration(t.Delay) * time.Millisecond)
		}

	}

	return err
}

func checkFileWriteErr(fileName string, err error) {
	if err != nil {
		log.WithField("Filename", fileName).WithError(err).Fatal("Couldn't write task output to file")
	}
}
