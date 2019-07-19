package file

import (
	"flag"
	"log"
	"os"
	"project-cooker/constants"
)

// WriteStringToFile is used to write a string to file
func WriteStringToFile(logLine string, level string) {
	flag.Parse()
	var file, openError = os.OpenFile(constants.LOGFILEPATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if openError != nil {
		panic(openError)
	}

	defer file.Close()

	if level == "DEBUG" {
		log.Println(logLine)
		logger := log.New(file, level+" ", log.LstdFlags|log.Lshortfile)
		logger.Println(logLine)
	} else {
		log.Println(logLine)
		logger := log.New(file, level+" ", log.LstdFlags)
		logger.Println(logLine)
	}

}

// WriteScannedFilesToFile is used to write all found files to the scanner log
func WriteScannedFilesToFile(logLine string) {
	flag.Parse()
	var file, openError = os.OpenFile(constants.SCANNEDFILESPATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if openError != nil {
		panic(openError)
	}

	defer file.Close()
	logger := log.New(file, "", log.LstdFlags)
	logger.Println(logLine)
}

// WriteBytesToFile is used to write a byte array to file
func WriteBytesToFile(logLine []byte, level string) {
	flag.Parse()
	var file, openError = os.OpenFile(constants.LOGFILEPATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if openError != nil {
		panic(openError)
	}

	defer file.Close()

	if level == "DEBUG" {
		log.Println(logLine)
		logger := log.New(file, level+" ", log.LstdFlags|log.Lshortfile)
		logger.Println(logLine)
	} else {
		log.Println(logLine)
		logger := log.New(file, level+" ", log.LstdFlags)
		logger.Println(logLine)
	}
}
