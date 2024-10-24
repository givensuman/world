package utils

import (
	"io/fs"
	"os"
	"strings"
)

func ReadPackagesFromFile(filename string) []string {
	fsys := os.DirFS(GetWorldDir())

	data, err := fs.ReadFile(fsys, filename)
	if err != nil {
		panic(err)
	}

	packages := strings.Split(string(data), "\n")
	for i := 0; i < len(packages); i++ {
		packages[i] = strings.TrimSpace(packages[i])
	}

	return packages
}
