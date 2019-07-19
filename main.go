package main

import (
	"project-cooker/configuration"
	"project-cooker/constants"
	"project-cooker/file"
)

func main() {
	file.CreateProjectRegistry("examples/" + configuration.CurrentConfig.GetProjectStructure().RootPath)

	file.WriteStringToFile("Loading Configuration...", constants.INFO)
	file.WriteStringToFile(configuration.CurrentConfig.GetShellToUse(), constants.INFO)
	file.WriteStringToFile(configuration.CurrentConfig.TransformProjectStructureToJSON(), constants.INFO)

}
