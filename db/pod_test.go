package db

import (
  "fmt"
  "testing"
)

func TestGetPods(t *testing.T) {
  var pods []Pod
  pods = GetPods(1)
  fmt.Println(pods)
}

func TestGetPodsBySvc(t *testing.T) {
  pods := GetPodsBySvc("ocx-svc")
  fmt.Println(pods)
}

func TestUpdatePod(t *testing.T) {
  /**
  p := Pod{
   Did:    18,
   Name:   "ocx-svc",
   Eid:    9,
   Ename:  "65-121-file-serv",
  }
  */
  p := Pod{}
  //Db.Where("id = ?", 22).Take(&p)
  Db.Where("name = ? and eid = ?", "ocx-svc", 9).Take(&p)

  p.Cid = "xxxxx"
  err := UpdatePod(&p)
  fmt.Println(err)
}

func TestGetPodsByPort(t *testing.T) {
  pods := GetPodsByPort("9999")
  fmt.Println(pods)

  pods = GetPodsByPort("9700")
  fmt.Println(pods)
}

func TestGetPodPortBySvc(t *testing.T) {
  pods := GetPodPortBySvc("db-space")
  fmt.Println(pods)
}









