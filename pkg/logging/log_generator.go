package logging

import (
	"fmt"
	"log"
	"os"
)

// NewLogFile Creates a log file and clears existing one
func NewLogFile(name string) *os.File {
	logPath := "./var/log/" + name + ".log"
	if _, err := os.Stat(logPath); err == nil {
		fmt.Printf("Removing existing Log file %v", logPath)
		e := os.Remove(logPath)
		if e != nil {
			log.Fatalf("Error removing log file: %v", e)
		}
	} else {
		fmt.Printf("Log file does not exist, creating new one: %v", logPath)
	}

	file, err := os.Create(logPath)

	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	return file
}
