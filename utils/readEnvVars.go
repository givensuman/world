package utils

import "os"

type EnvVars struct {
	WORLD_DIR string
}

func ReadEnvVars() EnvVars {
	world_dir, exists := os.LookupEnv("WORLD_DIR")
	if !exists {
		world_dir = "~/.world"
	}

	return EnvVars{
		WORLD_DIR: world_dir,
	}
}
