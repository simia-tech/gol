package gol

import "fmt"

type color int

const (
	colorBlack color = (iota + 30)
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite
)

var colorMapping = map[level]color{
	LEVEL_CRITICAL: colorMagenta,
	LEVEL_ERROR:    colorRed,
	LEVEL_WARNING:  colorYellow,
	LEVEL_INFO:     colorGreen,
	LEVEL_DEBUG:    colorCyan,
}

func surroundColor(text string, color color) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", int(color), text)
}
