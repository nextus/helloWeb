package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
)

func setVerboseLevel(level uint) {
	switch level {
	// Info
	case 4:
		logInit(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	// Warning
	case 3:
		logInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	// Error
	case 2:
		logInit(ioutil.Discard, ioutil.Discard, os.Stdout, os.Stderr)
	// Fatal
	case 1:
		logInit(ioutil.Discard, ioutil.Discard, ioutil.Discard, os.Stderr)
	default:
		log.Fatalf("[Log handler] Unknown debug level: %d", level)
	}
}

func logInit(infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer, fatalHandle io.Writer) {
	timePattern := log.Ldate | log.Ltime | log.Lshortfile

	Info = log.New(infoHandle,
		"INFO: ", timePattern)

	Warning = log.New(warningHandle,
		"WARNING: ", timePattern)

	Error = log.New(errorHandle,
		"ERROR: ", timePattern)

	Fatal = log.New(errorHandle,
		"FATAL: ", timePattern)
}
