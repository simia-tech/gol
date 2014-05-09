package gol

import (
	"fmt"
	"github.com/juju/errgo"
	"io"
	"io/ioutil"
	"log"
)

type channel struct {
	writers []io.Writer
	mapper  mapper
	logger  *log.Logger
}

var channels = map[level]*channel{
	LEVEL_CRITICAL: &channel{writers: make([]io.Writer, 0, 3), logger: log.New(ioutil.Discard, "C ", log.LstdFlags)},
	LEVEL_ERROR:    &channel{writers: make([]io.Writer, 0, 3), logger: log.New(ioutil.Discard, "E ", log.LstdFlags)},
	LEVEL_WARNING:  &channel{writers: make([]io.Writer, 0, 3), logger: log.New(ioutil.Discard, "W ", log.LstdFlags)},
	LEVEL_INFO:     &channel{writers: make([]io.Writer, 0, 3), logger: log.New(ioutil.Discard, "I ", log.LstdFlags)},
	LEVEL_DEBUG:    &channel{writers: make([]io.Writer, 0, 3), logger: log.New(ioutil.Discard, "D ", log.LstdFlags)},
	LEVEL_PROTOCOL: &channel{writers: make([]io.Writer, 0, 3), logger: log.New(ioutil.Discard, "P ", log.LstdFlags)},
}

func (c *channel) Printf(format string, values ...interface{}) {
	if c.mapper == nil {
		c.logger.Printf(format, values...)
	} else {
		c.logger.Printf(c.mapper(fmt.Sprintf(format, values...)))
	}
}

// Critical uses format and values to generate a log message on the critical channel.
func Critical(format string, values ...interface{}) {
	channels[LEVEL_CRITICAL].Printf(format, values...)
}

// Error uses format and values to generate a log message on the error channel.
func Error(format string, values ...interface{}) {
	channels[LEVEL_ERROR].Printf(format, values...)
}

// Warning uses format and values to generate a log message on the warning channel.
func Warning(format string, values ...interface{}) {
	channels[LEVEL_WARNING].Printf(format, values...)
}

// Info uses format and values to generate a log message on the info channel.
func Info(format string, values ...interface{}) {
	channels[LEVEL_INFO].Printf(format, values...)
}

// Debug uses format and values to generate a log message on the debug channel.
func Debug(format string, values ...interface{}) {
	channels[LEVEL_DEBUG].Printf(format, values...)
}

// Debug uses format and values to generate a log message on the debug channel.
func Protocol(format string, values ...interface{}) {
	channels[LEVEL_PROTOCOL].Printf(format, values...)
}

// Handle takes the given error and writes it to all channels that correspond to the given levels. If no level is
// given, the error will be published on the error channel.
func Handle(err error, levels ...level) {
	if len(levels) == 0 {
		Handle(err, LEVEL_ERROR)
	} else {
		for _, level := range levels {
			channels[level].Printf(errgo.Details(err))
		}
	}
}
