package interfaces

// File is a interface with all available functions of the package file
type File interface {
	WriteStringToFile(logLine string)
	WriteBytesToFile(logline []byte)
}
