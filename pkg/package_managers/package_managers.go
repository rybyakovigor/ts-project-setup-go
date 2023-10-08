package package_managers

import (
	"errors"
	"fmt"
	"linters-configuration/pkg/core"
)

var installCommandMap = map[string]core.PackageManagerInstallCommand{
	"yarn": {"yarn", "add", "-D"},
	"npm":  {"npm", "install", "--save-dev"},
}

var installAllCommandMap = map[string]core.PackageManagerInstallCommand{
	"yarn": {"yarn", "install"},
	"npm":  {"npm", "install"},
}

func selectPackageManager() (string, error) {
	options := []string{"yarn", "npm"}

	fmt.Println("-----------------------")
	fmt.Println("Select package manager:")
	fmt.Println("-----------------------")

	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	var choice int

	fmt.Println("")
	fmt.Println("Enter number of package manager:")

	fmt.Scanln(&choice)

	if choice >= 1 && choice <= len(options) {
		return options[choice-1], nil
	} else {
		return "", errors.New("wrong input")
	}
}

func Choose() (core.PackageManagerInstallCommand, core.PackageManagerInstallCommand) {
	packageManager, err := selectPackageManager()

	if err != nil {
		panic(err)
	}

	return installCommandMap[packageManager], installAllCommandMap[packageManager]
}
