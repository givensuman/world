package utils

import (
	"fmt"
	"os"
)

type Env interface {
	GetWorldDir()
}

// Get the directory for packages
func GetWorldDir() string {
	world_dir, exists := os.LookupEnv("WORLD_DIR")
	if !exists {
		home_dir := os.Getenv("HOME")
		world_dir = fmt.Sprintf("%s/.world", home_dir)
	}

	return world_dir
}
