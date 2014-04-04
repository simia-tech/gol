package gol

type level int

const (
	LEVEL_CRITICAL level = 1 << iota
	LEVEL_ERROR
	LEVEL_WARNING
	LEVEL_INFO
	LEVEL_DEBUG
)

var allLevels = []level{LEVEL_CRITICAL, LEVEL_ERROR, LEVEL_WARNING, LEVEL_INFO, LEVEL_DEBUG}
