package gol

import (
	"io"
	"io/ioutil"
	"log"
)

type channel struct {
	writers []io.Writer
	logger  *log.Logger
}

var channels = map[level]*channel{
	LEVEL_CRITICAL: &channel{writers: make([]io.Writer, 0, 2), logger: log.New(ioutil.Discard, "C ", log.LstdFlags)},
	LEVEL_ERROR:    &channel{writers: make([]io.Writer, 0, 2), logger: log.New(ioutil.Discard, "E ", log.LstdFlags)},
	LEVEL_WARNING:  &channel{writers: make([]io.Writer, 0, 2), logger: log.New(ioutil.Discard, "W ", log.LstdFlags)},
	LEVEL_INFO:     &channel{writers: make([]io.Writer, 0, 2), logger: log.New(ioutil.Discard, "I ", log.LstdFlags)},
	LEVEL_DEBUG:    &channel{writers: make([]io.Writer, 0, 2), logger: log.New(ioutil.Discard, "D ", log.LstdFlags)},
}

func Critical(format string, values ...interface{}) {
	channels[LEVEL_CRITICAL].logger.Printf(format, values...)
}

func Error(format string, values ...interface{}) {
	channels[LEVEL_ERROR].logger.Printf(format, values...)
}

func Warning(format string, values ...interface{}) {
	channels[LEVEL_WARNING].logger.Printf(format, values...)
}

func Info(format string, values ...interface{}) {
	channels[LEVEL_INFO].logger.Printf(format, values...)
}

func Debug(format string, values ...interface{}) {
	channels[LEVEL_DEBUG].logger.Printf(format, values...)
}

func Handle(error error, levels ...level) {
	if len(levels) == 0 {
		Handle(error, LEVEL_ERROR)
	} else {
		for _, level := range levels {
			channels[level].logger.Print(error)
		}
	}
}
