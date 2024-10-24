package world

import (
	"github.com/Jguer/go-alpm/v2"
)

type PackageManager interface {
	Add()
	Remove()
	Install()
	Uninstall()
	Sync()
	Backup()
}

type World struct {
	Packages alpm.IPackageList
}
