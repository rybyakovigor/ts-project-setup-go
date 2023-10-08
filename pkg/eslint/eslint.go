package eslint

import (
	"linters-configuration/pkg/core"
)

func install(command core.PackageManagerInstallCommand) {
	core.InstallLib(command, "eslint")
	core.InstallLib(command, "@typescript-eslint/eslint-plugin")
	core.InstallLib(command, "@typescript-eslint/parser")
	core.InstallLib(command, "eslint-plugin-sonarjs")
	core.InstallLib(command, "eslint-plugin-prettier")
	core.InstallLib(command, "eslint-config-prettier")
}

func createConfig() {
	core.CreateConfig(".eslintrc", config)
}

func Install(command core.PackageManagerInstallCommand) {
	install(command)
	createConfig()
	core.WriteScript("check:lint", "eslint src")
	core.WriteScript("fix:lint", "eslint src --fix")
}

var config = `{
	"root": true,
	"extends": [
	  "eslint:recommended",
	  "plugin:@typescript-eslint/recommended",
	  "plugin:sonarjs/recommended",
	  "plugin:prettier/recommended"
	],
	"parser": "@typescript-eslint/parser",
	"parserOptions": {
	  "project": [
		"./tsconfig.json"
	  ],
	  "sourceType": "module",
	  "ecmaVersion": "latest"
	},
	"plugins": [
	  "@typescript-eslint",
	  "sonarjs"
	],
	"rules": {
	  "@typescript-eslint/explicit-member-accessibility": "error",
	  "@typescript-eslint/no-unused-vars": "error",
	  "@typescript-eslint/unbound-method": "off",
	  "@typescript-eslint/explicit-function-return-type": [
		"error",
		{
		  "allowExpressions": true
		}
	  ],
	  "no-else-return": [
		"error",
		{
		  "allowElseIf": false
		}
	  ],
	  "no-console": [
		"error",
		{
		  "allow": [
			"info",
			"warn",
			"error"
		  ]
		}
	  ],
	  "no-nested-ternary": "error",
	  "@typescript-eslint/no-non-null-assertion": "warn",
	  "sonarjs/no-duplicate-string": [
		"error",
		{
		  "threshold": 2
		}
	  ],
	  "prettier/prettier": "off"
	},
	"ignorePatterns": [
	  "src/**/*.test.ts"
  ]
}`
