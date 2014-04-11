package gol

type Configuration struct {
	// common
	Backend string `json:"backend"`
	Mask    string `json:"mask"`

	// file
	Path string `json:"path"`
	Mode string `json:"mode"`

	// syslog
	Prefix string `json:"prefix"`
}
