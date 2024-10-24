package utils

import (
	"fmt"
	"sort"
)

func WritePackagesToFile(pkgs []string, filename string) error {
	packages := ReadPackagesFromFile(filename)

	packages = append(packages, pkgs...)
	sort.Strings(packages)

	fmt.Print(packages)
	return nil
}

func WritePackageToFile(pkg string, filename string) error {
	return WritePackagesToFile([]string{pkg}, filename)
}
