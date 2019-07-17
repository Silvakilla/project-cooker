package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"project-cooker/configuration"
	fileWriter "project-cooker/file"
)

var config configuration.IConfiguration

func setConfig() {
	file, readError := ioutil.ReadFile("configs/config.json")
	config = &configuration.Configuration{}
	readError = json.Unmarshal([]byte(file), config)

	if readError != nil {
		log.Println(readError)
	}
}

func main() {
	setConfig()

	fmt.Println("Welcome to Project Cooker")

	log.Println(config.GetShellToUse())
	log.Println(config.GetLogFilePath())
	log.Println(config.GetScannedFilesPath())
	log.Println(config.GetProjectStructure())
	log.Println(config.GetModStructure())

	fileWriter.WriteStringToFile("Configuration:")
	//logger.WriteStringToFile(config.GetConfig().LogFilePath)
	//logger.WriteStringToFile(configuration.IConfiguration.GetConfig().ScannedFilesPath)
	//logger.WriteStringToFile(configuration.IConfiguration.GetConfig().ShellToUse)
	//logger.WriteStringToFile(configuration.IConfiguration.ProjectStructureToJSON())

}
