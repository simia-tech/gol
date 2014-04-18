package gol

// Configuration holds all information that specify a single gol backend. While the common properties are needed
// for each backend, it holds also backend-specific properties.
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
