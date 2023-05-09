package eslint

import (
	"linters-configuration/pkg/core"
)

func install(command core.PackageManagerInstallCommand) {
	core.InstallLib(command, "eslint")
	core.InstallLib(command, "@typescript-eslint/eslint-plugin")
	core.InstallLib(command, "@typescript-eslint/parser")
	core.InstallLib(command, "eslint-plugin-prettier")
	core.InstallLib(command, "eslint-config-prettier")
}

func createConfig() {
	core.CreateConfig(".eslintrc", config)
}

func Install(command core.PackageManagerInstallCommand) {
	install(command)
	createConfig()
	core.WriteScript("lint", "eslint src")
	core.WriteScript("lint:fix", "eslint src --fix")
}

var config = `{
	"root": true,
	"extends": [
	  "eslint:recommended",
	  "plugin:@typescript-eslint/recommended",
	  "plugin:prettier/recommended"
	],
	"parser": "@typescript-eslint/parser",
	"parserOptions": {
	  "project": [
		"./tsconfig.json"
	  ]
	},
	"plugins": [
	  "@typescript-eslint"
	],
	"rules": {
	  "no-console": [
		"error",
		{
		  "allow": [
			"warn",
			"error"
		  ]
		}
	  ],
	  "prettier/prettier": "off"
	},
	"ignorePatterns": [
	  "src/**/*.test.ts"
	]
  }`
