package util

import (
  "bytes"
  "reflect"
  "encoding/json"
)

func FormatData(data interface{}) *bytes.Buffer {
  j_data, err := json.Marshal(data)
  if err != nil {
    ErrorLogger.Print(err)
  }
  return bytes.NewBuffer(j_data)
}

func IsNil(content interface{}) bool {
  return content == nil || reflect.ValueOf(content).IsNil()
}
