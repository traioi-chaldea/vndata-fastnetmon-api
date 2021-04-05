package util

import (
  "log"
  "os"
)

var (
  WarningLogger *log.Logger
  InfoLogger    *log.Logger
  ErrorLogger   *log.Logger
)

func init() {
  file, err := os.OpenFile("/opt/fastnetmon/fastnetmon_api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
  if err != nil {
    log.Fatal(err)
  }

  InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
  WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
  ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func PrintInfoLogger(e error, logger string) bool {
  var check bool
  if e != nil {
    ErrorLogger.Println(e)
    check = false
  } else {
    InfoLogger.Println(logger)
    check = true
  }
  return check
}

