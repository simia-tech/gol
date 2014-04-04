package gol

import "fmt"

type mask level

const (
	MASK_ALL     mask = mask(LEVEL_CRITICAL) | mask(LEVEL_ERROR) | mask(LEVEL_WARNING) | mask(LEVEL_INFO) | mask(LEVEL_DEBUG)
	MASK_DEFAULT mask = mask(LEVEL_CRITICAL) | mask(LEVEL_ERROR) | mask(LEVEL_WARNING) | mask(LEVEL_INFO)
)

func maskByName(name string) (mask mask, error error) {
	switch (name) {
	case "all":
		mask = MASK_ALL
		return
	case "default":
		mask = MASK_DEFAULT
		return
	}
	error = fmt.Errorf("unknown mask '%s', try 'all' or 'default'", name)
	return
}
