package husky

import (
	"bytes"
	"linters-configuration/pkg/core"
	"text/template"
)

func Install(command core.PackageManagerInstallCommand) {
	install(command)
	createCommitlintConfig()
	addCommitMsgHook()
	addPreCommitHook(command)
}

func install(command core.PackageManagerInstallCommand) {
	var c = core.PackageManagerInstallCommand{"npx", "husky-init", "&&"}

	core.InstallLib(c, command[0])
	core.InstallLib(command, "@commitlint/config-conventional")
	core.InstallLib(command, "@commitlint/cli")
}

func createCommitlintConfig() {
	commitlintConfig := "module.exports = {extends: ['@commitlint/config-conventional']}"
	core.CreateConfig("commitlint.config.js", commitlintConfig)
}

func addCommitMsgHook() {
	var c = core.PackageManagerInstallCommand{"npx", "husky", "add"}
	core.InstallLib(c, ".husky/commit-msg")

	hook := `#!/usr/bin/env sh
. "$(dirname -- "$0")/_/husky.sh"

npx --no-install commitlint --edit "$1"`

	core.CreateConfig(".husky/commit-msg", hook)
}

func addPreCommitHook(command core.PackageManagerInstallCommand) {
	hook := `#!/usr/bin/env sh
. "$(dirname -- "$0")/_/husky.sh"

{{ . }} run lint
{{ . }} run prettier
{{ . }} run stylelint
`
	t, err := template.New("hook").Parse(hook)

	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, command[0])
	if err != nil {
		panic(err)
	}

	core.CreateConfig(".husky/pre-commit", buf.String())
}
