package gol

import (
	"io"
	"log"
	"log/syslog"
	"os"
)

func Initialize(configurations ...*Configuration) error {
	for _, configuration := range configurations {
		backend, error := backendByName(configuration.Backend)
		if error != nil {
			return error
		}

		mask, error := maskByName(configuration.Mask)
		if error != nil {
			return error
		}

		switch (backend) {
		case BACKEND_CONSOLE:
			addConsoleOutput(mask)
		case BACKEND_SYSLOG:
			addSyslogOutput(mask, configuration)
		}
	}

	return nil
}

func addConsoleOutput(mask mask) {
	addOutputWriter(mask, os.Stdout)
}

func addSyslogOutput(mask mask, configuration *Configuration) {
	writer, error := syslog.New(syslog.LOG_DEBUG | syslog.LOG_MAIL, configuration.Prefix)
	if error != nil {
		Handle(error)
		return
	}

	addOutputWriter(mask, writer)
}

func addOutputWriter(mask mask, writer io.Writer) {
	for _, level := range allLevels {
		if int(mask) & int(level) != 0 {
			channel := channels[level]
			channel.writers = append(channel.writers, writer)
			channel.logger = log.New(io.MultiWriter(channel.writers...), channel.logger.Prefix(), channel.logger.Flags())
		}
	}
}