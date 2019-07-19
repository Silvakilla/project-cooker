package file

import (
	"flag"
	"log"
	"os"

	"project-cooker/configuration"
)

var config configuration.Configuration

// WriteStringToFile is used to write a string to file
func WriteStringToFile(logLine string) {

	flag.Parse()

	var file, openError = os.OpenFile(config.GetLogFilePath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if openError != nil {
		panic(openError)
	}

	defer file.Close()

	logger := log.New(file, "", log.LstdFlags|log.Lshortfile)
	logger.Println(logLine)
}

// WriteBytesToFile is used to write a byte array to file
func WriteBytesToFile(logLine []byte) {
	flag.Parse()

	var file, openError = os.OpenFile(config.GetLogFilePath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if openError != nil {
		panic(openError)
	}

	defer file.Close()

	logger := log.New(file, "", log.LstdFlags|log.Lshortfile)
	logger.Println(logLine)
}
