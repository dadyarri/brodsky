package info

import "runtime/debug"

var version string

func GetVersion() string {
	if version != "" {
		return version // Use the version if already set (e.g., via build flags)
	}

	buildInfo, ok := debug.ReadBuildInfo()
	if ok && buildInfo.Main.Version != "(devel)" {
		return buildInfo.Main.Version
	}

	return "unknown" // Default if version information isn't available
}

func GetAppName() string {
	return "Brodsky"
}
