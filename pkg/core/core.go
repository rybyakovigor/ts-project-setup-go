package core

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type PackageManagerInstallCommand [3]string

var PACKAGE_JSON = "package.json"

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
	packageJSON := openJSONFile(PACKAGE_JSON)

	scripts := packageJSON["scripts"].(map[string]interface{})
	scripts[scriptName] = scriptValue

	writeJSONFile(PACKAGE_JSON, packageJSON)
}

func RemoveCapsInDeps() {
	packageJSON := openJSONFile(PACKAGE_JSON)

	deps := packageJSON["dependencies"].(map[string]interface{})
	devDeps := packageJSON["devDependencies"].(map[string]interface{})

	removeSymbolInMap(deps, "^")
	removeSymbolInMap(devDeps, "^")

	writeJSONFile(PACKAGE_JSON, packageJSON)
}

func openJSONFile(fileName string) map[string]interface{} {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		panic(err)
	}

	return jsonData
}

func writeJSONFile(fileName string, jsonData map[string]interface{}) {
	newData, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fileName, newData, 0644)
	if err != nil {
		panic(err)
	}
}

func removeSymbolInMap(data map[string]interface{}, symbol string) {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	for _, key := range keys {
		if data[key] != nil {
			data[key] = strings.Replace(data[key].(string), symbol, "", 1)
		}
	}
}
