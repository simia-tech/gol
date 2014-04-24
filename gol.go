/*
gol is an extended logger for go, that can log to a console, file and syslog backend via a single interface. It
supports multiple logging channels and the output can be filtered using a mask. On top of this, multiple of
backends can be used at the same time.
*/
package gol

import (
	"fmt"
	"io"
	"log"
	"log/syslog"
	"os"
)

/*
Initialize sets up gol with a number of given configurations. Each configuration specified a single backend. By
default, no backend is specified. In order to see any log message, Initialize has to be called.

For example

		gol.Initialize(&gol.Configuration{Backend: "console", Mask: "all"})

will log all messages to the console.
*/
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

		switch backend {
		case BACKEND_CONSOLE:
			addConsoleOutput(mask, configuration)
		case BACKEND_FILE:
			addFileOutput(mask, configuration)
		case BACKEND_SYSLOG:
			addSyslogOutput(mask, configuration)
		}
	}

	return nil
}

func addConsoleOutput(mask mask, configuration *Configuration) {
	if configuration.Color {
		for _, level := range allLevels {
			channels[level].mapper = buildTerminalColorMapper(colorMapping[level])
		}
	}

	addOutputWriter(mask, os.Stdout)
}

func addFileOutput(mask mask, configuration *Configuration) {
	var mode int
	fmt.Sscanf(configuration.Mode, "%o", &mode)

	logFile, error := os.OpenFile(configuration.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.FileMode(mode))
	if error != nil {
		log.Fatal(error)
	}

	addOutputWriter(mask, logFile)
}

func addSyslogOutput(mask mask, configuration *Configuration) {
	writer, error := syslog.New(syslog.LOG_EMERG|syslog.LOG_LOCAL0, configuration.Prefix)
	if error != nil {
		log.Fatal(error)
	}

	addOutputWriter(mask, writer)
}

func addOutputWriter(mask mask, writer io.Writer) {
	for _, level := range allLevels {
		if int(mask)&int(level) != 0 {
			channel := channels[level]
			channel.writers = append(channel.writers, writer)
			channel.logger = log.New(io.MultiWriter(channel.writers...), channel.logger.Prefix(), channel.logger.Flags())
		}
	}
}
