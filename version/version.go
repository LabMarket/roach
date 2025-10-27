package version

// This version-string will be updated by the developer.
var (
	Major   = "0"
	Minor   = "1"
	Release = "0"
	Version = Major + "." + Minor + "." + Release
)

// The 'BuildNumber' variable will be set by the Go linker during compilation.
// It must be a package-level 'var' of type string.
var BuildNumber = "development-build"
