package logger

import (
	"fmt"
	"log"
	"sync"
)

const (
	// Silent is a level that logs nothing.
	Silent int = iota
	// Error is a level that logs error logs.
	Error
	// Info is a level that logs error, and info logs.
	Info
	// Debug is a level that logs error, info, and debug logs.
	Debug
)

var (
	lock sync.Mutex
)

// Infof logs an info message.
func Infof(format string, args ...interface{}) {
	lock.Lock()
	defer lock.Unlock()

	msg := fmt.Sprintf("INFO: %s", format)
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	log.Println(msg)
}

// Debugf logs a debug message.
func Debugf(format string, args ...interface{}) {
	lock.Lock()
	defer lock.Unlock()

	msg := fmt.Sprintf("DEBUG: %s", format)
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	log.Println(msg)
}

// Errorf logs an error message.
func Errorf(format string, args ...interface{}) {
	lock.Lock()
	defer lock.Unlock()

	msg := fmt.Sprintf("ERROR: %s", format)
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	log.Println(msg)
}
