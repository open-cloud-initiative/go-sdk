package dialer

import (
	"fmt"
	"runtime/debug"
)

func UserAgent() string {
	agent := "oci/go-sdk"

	version := "unknown"
	if build, ok := debug.ReadBuildInfo(); ok {
		if build.Main.Version != "" {
			version = build.Main.Version
		}
	}

	return fmt.Sprintf("%s/%s", agent, version)
}
