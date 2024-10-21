package world

// World object for use in CLI
type World struct {
	Packages []string
}

// Arguments passed to various world methods
//
// $ world add curl --to fetch_tools
type Args struct {
	Package string
	To      string
}

// Initializes world object
func (world *World) init() error {
	return nil
}

// Adds package(s) to world file and installs it
func (world *World) add(pkg string) error {
	return nil
}
