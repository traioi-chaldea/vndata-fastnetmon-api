package util

import (
  "bytes"
  "encoding/json"
)

func FormatData(data interface{}) *bytes.Buffer {
  j_data, err := json.Marshal(data)
  if err != nil {
    ErrorLogger.Print(err)
  }
  return bytes.NewBuffer(j_data)
}

