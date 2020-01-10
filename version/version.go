package version

import (
	"fmt"
	"runtime"
	"strings"
)

var (
	// Build is a time label of the moment when the binary was built
	Build = "unset"
	// Commit is a last commit hash at the moment when the binary was built
	Commit = "unset"
	// Release is a semantic version of current build
	Release = "unset"
)

// String return formatted version string
func String() string {
	return strings.Join([]string{
		fmt.Sprintf(strings.Repeat("-", 80)),
		fmt.Sprintf("    Release:     %s@%s", Release, Commit),
		fmt.Sprintf("    Build:       %s", Build),
		fmt.Sprintf("    GoVersion:   %s", runtime.Version()),
		fmt.Sprintf(strings.Repeat("-", 80)),
	}, "\n")
}
