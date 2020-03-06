package cmd

import (
  "encoding/json"
  "fmt"
  "math/rand"
  "testing"
  "time"
)


const (
  // fixme
  testImg = "service1"
)

func TestMarshalCts(t *testing.T) {
  //s := `{"Id": "xxx", "Image": "yyy", "State": "zzzz", "Ports": [{"IP": "8.8.8.8"]`
  output := `[{"Command":"/jvmIdx","Created":1575359495,"HostConfig":{"NetworkMode":"default"},"Id":"a3bd5483b93ded1401f74be2523b1ed3e949cf6c719a3fe8341a84fda559179f","Image":"10.1.1.9:5000/jvmidx:latest","ImageID":"sha256:dfa471cd036d3ea2d17d229050993f92b66cb43f268ab62a9edabf040a8ca57a","Labels":{},"Mounts":[{"Destination":"/etc/localtime","Mode":"ro","Propagation":"rprivate","RW":false,"Source":"/etc/localtime","Type":"bind"}],"Names":["/jvmIdx-test"],"NetworkSettings":{"Networks":{"bridge":{"Aliases":null,"DriverOpts":null,"EndpointID":"a7879b0a54b3047d2cc837523809e788533b83b2ee2e63682e58dc30b0fca590","Gateway":"172.17.0.1","GlobalIPv6Address":"","GlobalIPv6PrefixLen":0,"IPAMConfig":null,"IPAddress":"172.17.0.3","IPPrefixLen":16,"IPv6Gateway":"","Links":null,"MacAddress":"02:42:ac:11:00:03","NetworkID":"92d9381add3ad16bd3652a8cae697b765885c963796cb1d4047ccc68a5078348"}}},"Ports":[{"IP":"0.0.0.0","PrivatePort":8080,"PublicPort":8089,"Type":"tcp"}],"State":"running","Status":"Up 19 hours"}]`
  containers := make([]Container, 0)

  json.Unmarshal([]byte(output), &containers)

  fmt.Println(containers)
}

func TestPortainer_PullImage(t *testing.T) {
  return
  //p := NewPortainer()
  //data := make(map[string]interface{})
  //data["fromImage"] = "10.1.1.9:5000/micro_demo"
  //data["tag"] = "latest"
  //path := "http://10.1.1.40:9000/api/endpoints/11/docker/images/create?fromImage=10.1.1.9:5000%2Fmicro_demo&tag=latest"
  //p.PullImage(data, path)

  p := NewPortainer()
  pullRes := p.PullImage(11, testImg)
  fmt.Println(pullRes)
}

func TestPortainer_CreateCt(t *testing.T) {
  return
  p := NewPortainer()
  cId, err := p.CreateCt(11, testImg, "xx-dev",
                          "8080", "5001")
  fmt.Println(cId, err)
}

func TestPortainer_StartCt(t *testing.T) {
  return
  p := NewPortainer()
  res := p.StartCt(11, "9fb22912c536ffa9515ad9b2aa6b14958fc8aff3fdf95669342c73a113d49548", "")
  fmt.Println(res)
}

func TestPortainer_RmCt(t *testing.T) {
  return
  p := NewPortainer()
  p.RmCt(11, "xx-dev")
}

func TestPortainer_DeployCt(t *testing.T) {
  return
  p := NewPortainer()
  err := p.DeployCt(11, "ocx-test", testImg, "8080", "9090", "")
  fmt.Println(err)
}

func TestGetRandom(t *testing.T) {
  return
  sl := MakeRange(0, 4)
  fmt.Println(sl)

  rand.Seed(time.Now().Unix())
  r := rand.Intn(4)
  fmt.Println(r)
  
  x := rand.Intn(4)
  fmt.Println(x)
}

func TestRandomSl(t *testing.T) {
  return
  n := 2
  sl := []int{0, 1, 2, 3, 4}
  rn := RandomSl(n, sl)
  fmt.Println(rn)
}

func TestScale_Random(t *testing.T) {
  //esSl := MakeRange(0, 5 - 1)
  //hitEsIdx := RandomSl(2, esSl)
  //fmt.Println(hitEsIdx)

  var s Scalex        // 可以继承 Scalex 的方法
  s = &Scale{}        // 可以接续 Scale  的方法
  //s := Scalex()
  h := s.Random(2, 4)
  fmt.Println(h)
}

func TestGenCtState(t *testing.T) {
  p := NewPortainer()
  endpoints := p.getEndpoints()
  fmt.Println("len ed -------------", len(endpoints))
  GenCtState(endpoints)
  fmt.Println(len(CtState))
}








