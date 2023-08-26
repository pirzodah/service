package log

import (
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	// app logging to file
	log.SetOutput(&lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 10,
		MaxAge:     90, //days
	})
}

// Log ...
func Log(packageName string, functionName string, positionLabel string, v interface{}) {
	log.Printf("%s %s %s %+v\n", packageName, functionName, positionLabel, v)
}

// Fatal ...
func Fatal(packageName string, functionName string, positionLabel string, v interface{}) {
	log.Fatalf("%s %s %s %+v\n", packageName, functionName, positionLabel, v)

}
