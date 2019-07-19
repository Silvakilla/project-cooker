package configuration

import (
	"encoding/json"
	"io/ioutil"

	"project-cooker/constants"
	fileWriter "project-cooker/file"
)

// CurrentConfig exposes the set configuration
var CurrentConfig IConfiguration

// IConfiguration is a interface with all available functions of config.go
type IConfiguration interface {
	SetConfig()
	GetShellToUse() string
	GetProjectStructure() ProjectStructure
	TransformProjectStructureToJSON() string
}

// Configuration for Installation and Initialization
type Configuration struct {
	IConfiguration
	ShellToUse       string
	ProjectStructure ProjectStructure
}

// ProjectStructure defines the development paths for developing the mod
type ProjectStructure struct {
	RootPath                string
	UnitModelRootPath       string
	UnitTextureRootPath     string
	UnitAudioRootPath       string
	BuildingModelRootPath   string
	BuildingTextureRootPath string
	BuildingAudioRootPath   string
	AiScriptsRootPath       string
	NamesRootPath           string
	TextRootPath            string
	ResourcesRootPath       string
	DefaultXMLPath          string
}

func init() {
	SetConfig()
}

// GetProjectStructure returns the structure of the project
func (config Configuration) GetProjectStructure() ProjectStructure {
	return config.ProjectStructure
}

// GetShellToUse returns the shell which is used for commands
func (config Configuration) GetShellToUse() string {
	return config.ShellToUse
}

// TransformProjectStructureToJSON transforms the project structure to a printable JSON format
func (config Configuration) TransformProjectStructureToJSON() string {
	json, marshalError := json.Marshal(config.ProjectStructure)

	if marshalError != nil {
		fileWriter.WriteStringToFile(marshalError.Error(), constants.ERROR)
	}

	return string(json)
}

// SetConfig sets or resets the configuration to what is set in the config.json
func SetConfig() {
	file, readError := ioutil.ReadFile(constants.CONFIGFILEPATH)
	CurrentConfig = &Configuration{}
	readError = json.Unmarshal([]byte(file), CurrentConfig)

	if readError != nil {
		fileWriter.WriteStringToFile(readError.Error(), constants.ERROR)
	}
}
