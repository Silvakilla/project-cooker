package file

// ProjectFiles is an array full of all found project files
var ProjectFiles []string

// CreateProjectRegistry creates the file registry
func CreateProjectRegistry(projectRootpath string) {
	ProjectFiles = ScanForFiles("./" + projectRootpath)
}
