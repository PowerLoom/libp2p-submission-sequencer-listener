package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func InitLogger() {
	// Set log level from environment
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info" // Default to info level instead of debug
	}
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
	
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	
	// Check if LOG_FILE environment variable is set
	logFile := os.Getenv("LOG_FILE")
	if logFile != "" {
		// Ensure log directory exists
		logDir := filepath.Dir(logFile)
		if err := os.MkdirAll(logDir, 0755); err != nil {
			fmt.Printf("Failed to create log directory: %v\n", err)
			// Fall back to stdout/stderr split
			setupConsoleLogging()
			return
		}
		
		// Parse rotation settings from environment
		maxSize := 100 // Default 100MB
		if val := os.Getenv("LOG_MAX_SIZE_MB"); val != "" {
			if size, err := strconv.Atoi(val); err == nil {
				maxSize = size
			}
		}
		
		maxBackups := 5 // Default keep 5 old files
		if val := os.Getenv("LOG_MAX_BACKUPS"); val != "" {
			if backups, err := strconv.Atoi(val); err == nil {
				maxBackups = backups
			}
		}
		
		maxAge := 30 // Default 30 days
		if val := os.Getenv("LOG_MAX_AGE_DAYS"); val != "" {
			if age, err := strconv.Atoi(val); err == nil {
				maxAge = age
			}
		}
		
		compress := false
		if val := os.Getenv("LOG_COMPRESS"); val == "true" || val == "1" {
			compress = true
		}
		
		// Create rotating log writer
		rotator := &lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    maxSize,    // megabytes
			MaxBackups: maxBackups, // number of old files
			MaxAge:     maxAge,     // days
			Compress:   compress,   // compress old files
			LocalTime:  true,
		}
		
		// Write to both rotating file and stdout
		multiWriter := io.MultiWriter(rotator, os.Stdout)
		log.SetOutput(multiWriter)
		
		log.WithFields(log.Fields{
			"file":       logFile,
			"maxSize":    maxSize,
			"maxBackups": maxBackups,
			"maxAge":     maxAge,
			"compress":   compress,
			"logLevel":   logLevel,
		}).Info("Initialized rotating logger")
	} else {
		// If no log file, use stdout/stderr split
		setupConsoleLogging()
		log.WithField("logLevel", logLevel).Info("Initialized console logger")
	}
}

func setupConsoleLogging() {
	log.SetOutput(io.Discard) // Send all logs to nowhere by default
	
	log.AddHook(&writer.Hook{ // Send logs with level higher than warning to stderr
		Writer: os.Stderr,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
		},
	})
	log.AddHook(&writer.Hook{ // Send info and debug logs to stdout
		Writer: os.Stdout,
		LogLevels: []log.Level{
			log.TraceLevel,
			log.InfoLevel,
			log.DebugLevel,
		},
	})
}
