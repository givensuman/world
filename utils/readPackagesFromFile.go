package utils

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func ReadPackagesFromFile() {
	filesystem := os.DirFS("/home/given/")

	bytes, _ := fs.ReadFile(filesystem, ".world/packages")

	fmt.Println(strings.Split(string(bytes), "\n"))
}
