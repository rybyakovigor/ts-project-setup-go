package stylelint

import (
	"errors"
	"fmt"
	"linters-configuration/pkg/core"
)

func askStylelintInstall() (string, error) {
	options := []string{"yes", "no"}

	fmt.Println("-----------------------")
	fmt.Println("Do you need install stylelint?:")
	fmt.Println("-----------------------")

	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	var choice int

	fmt.Println("")
	fmt.Println("Enter number of answer:")

	fmt.Scanln(&choice)

	if choice >= 1 && choice <= len(options) {
		return options[choice-1], nil
	} else {
		return "", errors.New("wrong input")
	}
}

func Install(command core.PackageManagerInstallCommand) {
	answer, err := askStylelintInstall()

	if err != nil {
		panic(err)
	}

	if answer == "yes" {
		install(command)
		createConfig()
		core.WriteScript("stylelint", "stylelint src/**/*.css")
		core.WriteScript("stylelint:fix", "stylelint src/**/*.css --fix")
	}
}

func install(command core.PackageManagerInstallCommand) {
	core.InstallLib(command, "stylelint")
	core.InstallLib(command, "stylelint-config-standard")
	core.InstallLib(command, "stylelint-order")
	core.InstallLib(command, "stylelint-config-recess-order")
}

func createConfig() {
	core.CreateConfig(".stylelintrc", config)
}

var config = `{
	"plugins": ["stylelint-order"],
	"extends": ["stylelint-config-standard", "stylelint-config-recess-order"],
	"ignoreFiles": ["node_modules"],
	"rules": {}
}`
