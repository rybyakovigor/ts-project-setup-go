package main

import (
	"linters-configuration/pkg/editorconfig"
	"linters-configuration/pkg/eslint"
	"linters-configuration/pkg/husky"
	"linters-configuration/pkg/package_managers"
	"linters-configuration/pkg/prettier"
	"linters-configuration/pkg/stylelint"
)

func main() {
	command := package_managers.Choose()
	editorconfig.Install()
	prettier.Install(command)
	eslint.Install(command)
	stylelint.Install(command)
	husky.Install(command)
}
