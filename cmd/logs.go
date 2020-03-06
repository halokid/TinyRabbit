package cmd

import (
	//"encoding/json"
	//"fmt"
	//"strconv"
  //"fmt"
  "fmt"
)

func (p Portainer) getCrLog(eId string, cId string) string {
  log := p.fetch("endpoints/" + eId + "/docker/containers/" + cId + "/logs?since=0&stderr=1&stdout=1&tail=100&timestamps=0")
  return log
}

func printLog(eId string, cId string) {
  p := NewPortainer()
  log := p.getCrLog(eId, cId)
  fmt.Println(log)
}


