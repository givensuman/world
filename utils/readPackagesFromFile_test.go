package utils

import (
	"fmt"
	"testing"
)

func TestReadPackagesFromFile(t *testing.T) {
	strs := ReadPackagesFromFile("packages")
	fmt.Print(strs)
}
