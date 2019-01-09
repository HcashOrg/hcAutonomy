package sharedconfig

import (
	"path/filepath"

	"github.com/HcashOrg/hcd/hcutil"
)

const (
	DefaultConfigFilename = "hcAutonomywww.conf"
	DefaultDataDirname    = "data"
)

var (
	// DefaultHomeDir points to hcAutonomywww's home directory for configuration and data.
	DefaultHomeDir = hcutil.AppDataDir("hcAutonomywww", false)

	// DefaultConfigFile points to hcAutonomywww's default config file.
	DefaultConfigFile = filepath.Join(DefaultHomeDir, DefaultConfigFilename)

	// DefaultDataDir points to hcAutonomywww's default data directory.
	DefaultDataDir = filepath.Join(DefaultHomeDir, DefaultDataDirname)
)
