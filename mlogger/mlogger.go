package mlogger

import (
	"log"
	"os"
	"sync"
)

type mlLogger struct {
	*log.Logger
	filename string
}

var mlogger *mlLogger
var once sync.Once

//GetInstance create a singleton instance of the hydra logger
func GetInstance() *mlLogger {
	once.Do(func() {
		mlogger = createLogger("mlogger.log")
	})
	return mlogger
}

//Create a logger instance
func createLogger(fname string) *mlLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &mlLogger{
		filename: fname,
		Logger:   log.New(file, "Medaliving ", log.Lshortfile),
	}
}
