package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// CurrentConfig exposes the set configuration
var CurrentConfig IConfiguration

// IConfiguration is a interface with all available functions of config.go
type IConfiguration interface {
	SetConfig()
	GetShellToUse() string
	GetLogFilePath() string
	GetScannedFilesPath() string
	GetModStructure() ModStructure
	GetProjectStructure() ProjectStructure
}

// Configuration for Installation and Initialization
type Configuration struct {
	IConfiguration
	ShellToUse       string
	LogFilePath      string
	ScannedFilesPath string
	ModStructure     ModStructure
	ProjectStructure ProjectStructure
}

// ModStructure defines the default paths for the mod
type ModStructure struct {
	RootPath       string
	MapsPath       string
	ModelsPath     string
	TexturesPath   string
	ShadersPath    string
	MusicPath      string
	SfxPath        string
	NamesPath      string
	ResourcesPath  string
	ScriptsPath    string
	TextPath       string
	DefaultXMLPath string
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

// GetModStructure returns the default structure of the mod
func (config Configuration) GetModStructure() ModStructure {
	return config.ModStructure
}

// GetShellToUse returns the shell which is used for commands
func (config Configuration) GetShellToUse() string {
	return config.ShellToUse
}

// GetLogFilePath returns the path to the log file
func (config Configuration) GetLogFilePath() string {
	return config.LogFilePath
}

// GetScannedFilesPath return the path of log file for the scanned files
func (config Configuration) GetScannedFilesPath() string {
	return config.ScannedFilesPath
}

// SetConfig sets or resets the configuration to what is set in the config.json
func SetConfig() {
	file, readError := ioutil.ReadFile("configs/config.json")
	CurrentConfig = &Configuration{}
	readError = json.Unmarshal([]byte(file), CurrentConfig)

	if readError != nil {
		log.Println(readError)
	}
}
