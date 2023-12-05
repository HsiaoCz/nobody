package logger

import (
	"log"
	"os"
)

func LoggerFile() *os.File {
	file, err := os.OpenFile("./nobody.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
