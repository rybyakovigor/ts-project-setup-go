package main

import (
	"linters-configuration/pkg/concurrently"
	"linters-configuration/pkg/core"
	"linters-configuration/pkg/editorconfig"
	"linters-configuration/pkg/eslint"
	"linters-configuration/pkg/husky"
	"linters-configuration/pkg/lint_staged"
	"linters-configuration/pkg/package_managers"
	"linters-configuration/pkg/prettier"
	"linters-configuration/pkg/stylelint"
	"linters-configuration/pkg/types"
)

func main() {
	installOneCommand, installAllCommand := package_managers.Choose()
	editorconfig.Install()
	types.Install()
	concurrently.Install(installOneCommand)
	prettier.Install(installOneCommand)
	eslint.Install(installOneCommand)
	stylelint.Install(installOneCommand)
	husky.Install(installOneCommand)
	lint_staged.Install(installOneCommand)
	core.RemoveCapsInDeps()
	core.ReInstallAll((installAllCommand))
}
