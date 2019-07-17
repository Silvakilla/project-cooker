package main

import (
	"fmt"
	"log"

	"github.com/tkanos/gonfig"
)

// ProjectStructure defines the Paths to all files
type ProjectStructure struct {
	AiPath         string
	AudioPath      string
	DefaultXMLPath string
	MapsPath       string
	TexturesPath   string
	UnitsPath      string
}

// Configuration for Installation and Initialization
type Configuration struct {
	ShellToUse         string
	LogOutput          string
	ScannedFilesOutput string
	ProjectStructure   ProjectStructure
}

func main() {
	configuration := Configuration{}
	err := gonfig.GetConf("config/config.json", &configuration)

	if err != nil {
		log.Printf("%s", err)
	}

	fmt.Printf("Welcome to Project Cooker")
}
