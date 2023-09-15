package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Chart struct {
	APIVersion string `yaml:"apiVersion"`
	Name       string `yaml:"name"`
	Version    string `yaml:"version"`
	AppVersion string `yaml:"appVersion"`
	// Add other fields as needed
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: helmchart-updater <chart-path> <new-version> <new-appversion>")
		os.Exit(1)
	}

	chartPath := os.Args[1]
	newVersion := os.Args[2]
	newAppVersion := os.Args[3]

	// Read the existing Chart.yaml content
	chartYAML, err := ioutil.ReadFile(chartPath)
	if err != nil {
		fmt.Printf("Error reading Chart.yaml: %v\n", err)
		os.Exit(1)
	}

	var chart Chart

	// Unmarshal the YAML content into the struct
	if err := yaml.Unmarshal(chartYAML, &chart); err != nil {
		fmt.Printf("Error unmarshalling Chart.yaml: %v\n", err)
		os.Exit(1)
	}

	// Update the version and appVersion fields
	chart.Version = newVersion
	chart.AppVersion = newAppVersion

	// Marshal the updated struct back to YAML
	updatedYAML, err := yaml.Marshal(&chart)
	if err != nil {
		fmt.Printf("Error marshalling Chart.yaml: %v\n", err)
		os.Exit(1)
	}

	// Write the updated YAML back to the file
	if err := ioutil.WriteFile(chartPath, updatedYAML, 0644); err != nil {
		fmt.Printf("Error writing updated Chart.yaml: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Chart.yaml updated successfully.")
}
