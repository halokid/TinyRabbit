package db

import (
  "fmt"
  "testing"
)

func TestGetDeploys(t *testing.T) {
  deploy := GetDeploys()
  fmt.Println(deploy)
  for _, d := range deploy {
    fmt.Println(d.ID, d.Name)
  }
}

func TestGetOneDeploy(t *testing.T) {
  //d := GetOneDeploy("xxx")
  d := GetOneDeploy("xx-dev")
  fmt.Println(d)
}

func TestAddDeploy(t *testing.T) {
  d := Deploy{Name: "yy-test"}
  fmt.Println(AddDeploy(&d))
  fmt.Println(d.ID)
}