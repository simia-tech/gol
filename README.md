# gol

gol is an extended logger for go, that can log to a console, file and syslog backend via a single interface. It
supports multiple logging channels and the output can be filtered using a mask. On top of this, multiple of
backends can be used at the same time.

See [godoc](http://godoc.org/github.com/simia-tech/gol) for more details.

## Usage

	package main

	import (
		"fmt"

		"github.com/simia-tech/gol"
	)

	func main() {
		gol.Initialize(
			&gol.Configuration{Backend: "console", Mask: "all", Color: true},
			&gol.Configuration{Backend: "syslog", Prefix: "gol-test"})

		gol.Info("number %d", 1)

		error := fmt.Errorf("nasty!")
		gol.Handle(error)
	}
