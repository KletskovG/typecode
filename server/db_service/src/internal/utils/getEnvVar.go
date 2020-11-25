package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	DB string
}

// GetEnvVar returns var from env or config.txt file
func GetEnvVar(name string) string {
	var ENV string

	if os.Getenv(name)  == "" {
		file, _ := ioutil.ReadFile("config.txt")
		envVars := strings.Split(string(file), ";")

		for i := range envVars {
			eqIndex := strings.Index(envVars[i], "=")
			varName := strings.Trim(envVars[i][:eqIndex], "")

			if varName == name {
				varValueLen := len(envVars[i])
				varValue := strings.Trim(
					envVars[i][eqIndex + 2: varValueLen - 1],
					"",
				)

				return varValue
			}
		}
	} else {
		ENV = os.Getenv(name)
	}

	return ENV
}