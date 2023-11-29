package core

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/shafiqaimanx/shellvenom/src/menu"
)

//go:embed shelListz.json
var jsonShellData embed.FS

type Item struct{
	Name 		string `json:"name"`
	Shell 		string `json:"shell"`
	Platform 	interface{} `json:"platform"`
	Description string `json:"description"`
}

type Config struct {
	ShelListz []Item `json:"shelListz"`
}

func ReadFileofShell() Config {
	data, err := jsonShellData.ReadFile("shelListz.json")
	if err != nil {
		fmt.Printf("%s[-]%s %v\n", menu.CRIMSON, menu.RESET, err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("%s[-]%s %v\n", menu.CRIMSON, menu.RESET, err)
	}
	return config
}

// Extract and convert platform data into a slice of strings
func ExtractPlatforms(platform interface{}) []string {
	var platforms []string

	switch platformData := platform.(type) {
	case []interface{}:
		for _, value := range platformData {
			if str, ok := value.(string); ok {
				platforms = append(platforms, str)
			}
		}
	case string:
		platforms = append(platforms, platformData)
	}
	return platforms
}

// Checking if a string exists in a slice of strings
func Contains(platformSlice []string, str string) bool {
	for _, platformStr := range platformSlice {
		if platformStr == str {
			return true
		}
	}
	return false
}