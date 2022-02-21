package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

func Init() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   os.Getenv("LOG_PATH") + "/app.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
}
