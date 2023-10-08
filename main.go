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
	command := package_managers.Choose()
	editorconfig.Install()
	types.Install()
	concurrently.Install(command)
	prettier.Install(command)
	eslint.Install(command)
	stylelint.Install(command)
	husky.Install(command)
	lint_staged.Install(command)
	core.RemoveCapsInDeps()
}
