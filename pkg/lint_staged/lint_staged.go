package lint_staged

import (
	"linters-configuration/pkg/core"
)

func install(command core.PackageManagerInstallCommand) {
	core.InstallLib(command, "lint-staged")
}

func createConfig() {
	core.CreateConfig(".lintstagedrc", backendConfig)
}

func Install(command core.PackageManagerInstallCommand) {
	install(command)
	createConfig()
	core.WriteScript("check:all:pre-commit", "concurrently -r \"npx lint-staged --relative\" \"npm:check:types\"")
}

var backendConfig = `{
  "*.ts": ["prettier --check", "eslint"]
}`
