package internal

import log "github.com/sirupsen/logrus"

func RunTask(taskName string) {
	task, err := GetTaskDefinition(taskName)

	if err != nil {
		log.WithError(err).Error("Failed to get task definition")
		return
	}

	err = task.Execute()

	if err != nil {
		log.WithError(err).Error("Failed to execute task")
	}

	log.WithField("Name", task.Name).Trace("Command executed successfully")
}
