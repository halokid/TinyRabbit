package cmd

import (
  "fmt"
  "math/rand"
  "testing"
  "time"
)

func TestScale_Random2(t *testing.T) {
  n := make([]int, 0)
  for i := 0; i < 5; i++ {
    n = append(n, i)
    fmt.Println(i)
  }
  fmt.Println(n)
}

func TestC1(t *testing.T) {
  rand.Seed(time.Now().Unix())
  //k := rand.Intn(2)
  k := rand.Intn(1)
  fmt.Println(k)
}

func TestGetRandomPort(t *testing.T) {
  hostPort := GetRandomPort()
  fmt.Println(hostPort)
}