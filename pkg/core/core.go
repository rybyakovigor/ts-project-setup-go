package core

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type PackageManagerInstallCommand [3]string

func CreateConfig(fileName string, config string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, fileWriteErr := fmt.Fprintln(file, config)

	if fileWriteErr != nil {
		panic(fileWriteErr)
	}
}

func InstallLib(command PackageManagerInstallCommand, libName string) {
	cmd := exec.Command(command[0], command[1], command[2], libName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func WriteScript(scriptName string, scriptValue string) {
	data, err := os.ReadFile("package.json")
	if err != nil {
		panic(err)
	}

	var packageJSON map[string]interface{}
	err = json.Unmarshal(data, &packageJSON)
	if err != nil {
		panic(err)
	}

	scripts := packageJSON["scripts"].(map[string]interface{})
	scripts[scriptName] = scriptValue

	newData, err := json.MarshalIndent(packageJSON, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("package.json", newData, 0644)
	if err != nil {
		panic(err)
	}
}
