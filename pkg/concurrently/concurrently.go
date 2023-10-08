package concurrently

import (
	"linters-configuration/pkg/core"
)

func Install(command core.PackageManagerInstallCommand) {
	install(command)
}

func install(command core.PackageManagerInstallCommand) {
	core.InstallLib(command, "concurrently")
	core.WriteScript("check:all", "concurrently \"npm:check:*(!all)\"")
	core.WriteScript("fix:all", "concurrently \"npm:check:*(!all)\"")
}
