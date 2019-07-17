package configuration

import (
	logger "project-cooker/interfaces"
)

var sLog logger.File

// IConfiguration is a interface with all available functions of config.go
type IConfiguration interface {
	GetShellToUse() string
	GetLogFilePath() string
	GetScannedFilesPath() string
	GetModStructure() modStructure
	GetProjectStructure() projectStructure
}

// Configuration for Installation and Initialization
type Configuration struct {
	IConfiguration
	ShellToUse       string
	LogFilePath      string
	ScannedFilesPath string
	ModStructure     modStructure
	ProjectStructure projectStructure
}

// ModStructure defines the default paths for the mod
type modStructure struct {
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
type projectStructure struct {
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

func (config Configuration) GetProjectStructure() projectStructure {
	return config.ProjectStructure
}

func (config Configuration) GetModStructure() modStructure {
	return config.ModStructure
}

func (config Configuration) GetShellToUse() string {
	return config.ShellToUse
}

func (config Configuration) GetLogFilePath() string {
	return config.LogFilePath
}

func (config Configuration) GetScannedFilesPath() string {
	return config.ScannedFilesPath
}
