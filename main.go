package main

import (
	"fmt"
	"log"
	"project-cooker/configuration"
	fileWriter "project-cooker/file"
)

func main() {

	fmt.Println("Welcome to Project Cooker")

	log.Println(configuration.CurrentConfig.GetShellToUse())
	log.Println(configuration.CurrentConfig.GetLogFilePath())
	log.Println(configuration.CurrentConfig.GetScannedFilesPath())
	log.Println(configuration.CurrentConfig.GetProjectStructure())
	log.Println(configuration.CurrentConfig.GetModStructure())

	fileWriter.WriteStringToFile("Configuration:")
	//logger.WriteStringToFile(config.GetConfig().LogFilePath)
	//logger.WriteStringToFile(configuration.IConfiguration.GetConfig().ScannedFilesPath)
	//logger.WriteStringToFile(configuration.IConfiguration.GetConfig().ShellToUse)
	//logger.WriteStringToFile(configuration.IConfiguration.ProjectStructureToJSON())

}
