package prettier

import (
	"linters-configuration/pkg/core"
)

func install(command core.PackageManagerInstallCommand) {
	core.InstallLib(command, "prettier")
}

func createConfig() {
	core.CreateConfig(".prettierrc", config)
}

func Install(command core.PackageManagerInstallCommand) {
	install(command)
	createConfig()
	core.WriteScript("check:prettier", "prettier --check src")
	core.WriteScript("fix:prettier", "prettier --write src")
}

var config = `{
  "arrowParens": "always",
  "bracketSpacing": true,
  "embeddedLanguageFormatting": "auto",
  "htmlWhitespaceSensitivity": "css",
  "insertPragma": false,
  "jsxSingleQuote": false,
  "printWidth": 120,
  "proseWrap": "preserve",
  "quoteProps": "as-needed",
  "requirePragma": false,
  "semi": true,
  "singleQuote": true,
  "tabWidth": 2,
  "trailingComma": "es5",
  "useTabs": false,
  "vueIndentScriptAndStyle": false
}`
