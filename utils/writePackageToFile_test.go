package utils

import "testing"

func TestWritePackageToFile(t *testing.T) {
	writePackagesToFile([]string{"zzz"}, "packages")
}
