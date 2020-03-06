package gwx

import (
  "fmt"
  "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"
  "io/ioutil"
  "regexp"
  "strings"
  "sync"
  "testing"
  "time"
)

var ngx Ngx

/**
func init()  {
  ngx.Conf = "./conf/nginx.conf"
  ngx.ServerFlag = "#serverProc--"
  ngx.UpstreamFlag = "#upstreamProc--"
  ngx.NgxRlCmd = "pwd"
}
*/

func init()  {
  ngx = Ngxx
}

func TestNgx_GetFlagLine(t *testing.T) {
  return
  flag1 := ngx.GetFlagLine(ngx.ServerFlag)
  fmt.Println(flag1)

  flag2 := ngx.GetFlagLine(ngx.UpstreamFlag)
  fmt.Println(flag2)
}

func TestNgx_ReplCfg(t *testing.T) {
  return
  flagLine := ngx.GetFlagLine(ngx.ServerFlag)
  cfgCtx := "\t\t\t\tlocation /xxdevyy { proxy_pass http://xxdevyy/; }";
  newCfg := MakeCfg(flagLine, cfgCtx)
  ngx.ReplCfg(flagLine, newCfg)

  time.Sleep(1 * time.Second)

  flagLine2 := ngx.GetFlagLine(ngx.UpstreamFlag)
  cfgCtx2 := "\t\tupstream  xxdevyy { server  8.8.8.8:5001  weight=1; }";
  newCfg2 := MakeCfg(flagLine2, cfgCtx2)
  ngx.ReplCfg(flagLine2, newCfg2)


}

func TestRegSearch(t *testing.T) {
  return
  r, err := ioutil.ReadFile("/pathto/nginx.conf")
  fmt.Println(ngx.Conf)
  fmt.Println(string(r))
  if err != nil {
    fmt.Println(err)
  }
  //rege, err := regexp.Compile(`upstream(.*){ server(.*) weight=1; }`)
  rege := regexp.MustCompile(`upstream(.*){ server(.*) weight=1; }`)
  ps := rege.FindAllStringSubmatch(string(r), -1)
  fmt.Println(ps)
  fmt.Println(strings.Trim(ps[0][1], " "))
  fmt.Println(ps[0][2])
}

func TestNgx_NgxServs(t *testing.T) {
  s := ngx.NgxServs()
  fmt.Println(s)
}

func TestNgx_NgxUpstrs(t *testing.T) {
  u := ngx.NgxUpstrs()
  fmt.Println(u)
}

func TestNgx_SetServer(t *testing.T) {
  s := "xxx"
  //s := "abc"
  ngx.SetServer(s)
}

func TestNgx_SetUpstream(t *testing.T) {
  u := "xxx"
  ngx.SetUpstream(u, []string{"9.9.9.9:3333", "8.8.8.8:2222"})
}

func TestGetYaml(t *testing.T) {
  home, err := homedir.Dir()
  fmt.Println(err, home)
  //viper.AddConfigPath(home)
  //viper.SetConfigName(".barge")

  viper.SetConfigFile(home + "\\ocx.yaml")
  //viper.SetConfigFile("C:/Users/Jimmy.Li/ocx.yaml")
  fmt.Println(viper.ConfigFileUsed())

  viper.ReadInConfig()
  url := viper.GetString("portainer_url") + "/api"
  fmt.Println(url)
}

func TestNgx_ReloadNgx(t *testing.T) {
  ngx.ReloadNgx()
}

func TestComm(t *testing.T)  {
  s := " wsshe "
  fmt.Println(len(s))
  sx := strings.Trim(s, " ")
  //fmt.Println(strings.Trim(s, "  "))
  fmt.Println(sx)
  fmt.Println(len(s))
  fmt.Println(len(sx))

  sl := make([]int, 5)
  //sl = []int{1, 2, 3}
  sl = append(sl, 1)
  sl = append(sl, 2)
  sl = append(sl, 3)
  sl = append(sl, 4)
  sl = append(sl, 5)
  sl = append(sl, 6)
  //i, ok := sl[3]
  //fmt.Println(sl[3])
  fmt.Println(len(sl))
  fmt.Println(cap(sl))

  //ColorfulRabbit.OsExecOut("docker exec  dnginx-gw  /usr/sbin/nginx -s reload")
}

func TestLock(t *testing.T) {
  var m sync.Mutex
  m.Lock()
  //m.Lock()
  fmt.Println("locked")
  m.Unlock()
  m.Unlock()
}

func TestNgx_GetFlagLine2(t *testing.T) {
  s := ngx.GetFlagLine("yy-svc")
  fmt.Println(s)

  rs := "upstream yy-svc { server  8.8.8.8:99 weight=1; server  6.6.6.6:77 weight=1; }"
  err := ngx.ReplCfg(s, "\t" + rs)
  fmt.Println(err)
}



