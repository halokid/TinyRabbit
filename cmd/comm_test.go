package cmd

import (
  "fmt"
  "github.com/spf13/cast"
  "testing"
  "time"
)

func TestWriteToken(t *testing.T) {
  tm := time.Now().Unix()
  fmt.Println(tm)
  WriteToken(cast.ToString(tm) + ":xxxxxxxxx")
}

func TestReadToken(t *testing.T) {
  tk := ReadToken()
  fmt.Println(tk)
  
}