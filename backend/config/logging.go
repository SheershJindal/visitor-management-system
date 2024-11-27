package config

import (
	"io"
	"log"
	"os"
	"strconv"
)

// InitializeLogging sets up logging based on the LOG_FILE_PATH environment variable.
func InitializeLogging() {
	// Get the log file path from the environment variable
	logFilePath := os.Getenv("LOG_FILE_PATH")

	var logOutput io.Writer
	if logFilePath != "" {
		// Open or create the log file
		logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}

		// Log to both the file and stdout
		logOutput = io.MultiWriter(os.Stdout, logFile)
		log.Printf("Logging to file: %s", logFilePath)
	} else {
		// Default to logging only to stdout
		logOutput = os.Stdout
		log.Println("Logging to stdout only (LOG_FILE_PATH not set)")
	}

	// Set the global log output
	log.SetOutput(logOutput)

	// Set log format
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
}

// IsDetailedLoggingEnabled checks the feature flag for detailed logging
func IsDetailedLoggingEnabled() bool {
	enabled, err := strconv.ParseBool(os.Getenv("ENABLE_DETAILED_LOGGING"))
	if err != nil {
		return false
	}
	return enabled
}
