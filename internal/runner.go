package internal

import log "github.com/sirupsen/logrus"

func RunTask(taskName string) {
	task, err := GetTaskDefinition(taskName)

	if err != nil {
		log.WithError(err).Error("Failed to get task definition")
		return
	}

	out, err := task.Execute()

	if err != nil {
		log.WithError(err).Error("Failed to execute task")
	}

	log.WithFields(log.Fields{
		"Name":   task.Name,
		"Output": string(out),
	}).Trace("Command executed successfully")
}
