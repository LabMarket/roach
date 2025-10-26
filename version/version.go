package version

// This version-string will be updated by the developer.
var Major = "0"
var Minor = "1"
var Patch = "0"
var Version = Major + "." + Minor + "." + Patch


// The 'BuildNumber' variable will be set by the Go linker during compilation.
// It must be a package-level 'var' of type string.
var BuildNumber = "development-build"
