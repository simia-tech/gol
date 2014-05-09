package gol

import "fmt"

type mask level

const (
	MASK_ALL     mask = mask(LEVEL_CRITICAL) | mask(LEVEL_ERROR) | mask(LEVEL_WARNING) | mask(LEVEL_INFO) | mask(LEVEL_DEBUG) | mask(LEVEL_PROTOCOL)
	MASK_DEBUG   mask = mask(LEVEL_CRITICAL) | mask(LEVEL_ERROR) | mask(LEVEL_WARNING) | mask(LEVEL_INFO) | mask(LEVEL_DEBUG)
	MASK_DEFAULT mask = mask(LEVEL_CRITICAL) | mask(LEVEL_ERROR) | mask(LEVEL_WARNING) | mask(LEVEL_INFO)
	MASK_WARNING mask = mask(LEVEL_CRITICAL) | mask(LEVEL_ERROR) | mask(LEVEL_WARNING)
)

func maskByName(name string) (mask mask, error error) {
	switch name {
	case "all":
		mask = MASK_ALL
		return
	case "debug":
		mask = MASK_DEBUG
		return
	case "default", "":
		mask = MASK_DEFAULT
		return
	case "warning":
		mask = MASK_WARNING
		return
	}
	error = fmt.Errorf("unknown mask '%s', try 'all', 'debug', 'default' or 'warning'", name)
	return
}
