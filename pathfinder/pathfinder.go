package pathfinder

import (
	"os"
	"path/filepath"
)

// Find finds and returns the path to the currently executing binary
func Find() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}
