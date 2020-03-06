package gwx

import (
  "bufio"
  "fmt"
  . "github.com/halokid/ColorfulRabbit"
  "github.com/mitchellh/go-homedir"
  "github.com/pkg/errors"
  "github.com/spf13/viper"
  "io"
  "os"
  "strings"
  "sync"

  //"time"
)

/**
const (
  NgxConf = "./conf/nginx.conf"
  ServerFlag = "#serverProc--"
  UpstreamFlag = "#upstreamProc--"
  NgxRlCmd = "docker exec  dnginx-gw  /usr/sbin/nginx -s reload"
)
*/

type NgxI interface {
  SetServer() error
  SetUpstream() error
}

type Ngx struct {
  Conf         string         // 配置文件的路径
  ServerFlag   string         // 配置server的标识
  UpstreamFlag string         // 配置upstream的标识
  RpcAddr      string         // ngx服务器上的rpc服务端地址
  NgxRlCmd     string         // ngx重载的命令
  RpcAddrTest  string         // ngx 测试服务器上的rpc服务端地址
  //NgxRlCmdTest string         // ngx 测试服务器重载的命令
  mx           sync.Mutex     // 文件读写锁
}

type NgxServ struct {
  Name  string
  Nodes string
}

type NgxUpstr struct {
  Name string
}

var Ngxx Ngx
//var sm sync.Mutex

func init() {
  //viper.SetConfigFile("./tinyrabbit_gwx.yaml")
  home, err := homedir.Dir()
  CheckError(err, "$HOME目录读取失败")
  viper.AddConfigPath(home)
  viper.SetConfigName("gwx")
  err = viper.ReadInConfig()
  CheckFatal(err, "读取配置文件错误")
  fmt.Println("使用配置文件为:", viper.ConfigFileUsed())

  Ngxx.Conf = viper.GetString("gwx_conf")
  Ngxx.UpstreamFlag = "#upstreamProc--"
  Ngxx.ServerFlag = "#serverProc--"
  //Ngxx.NgxRlCmd = "docker exec  dnginx-gw  /usr/sbin/nginx -s reload"
  Ngxx.NgxRlCmd = viper.GetString("gwxReload")

  Ngxx.RpcAddr = viper.GetString("addr")
  Ngxx.RpcAddrTest = viper.GetString("addrTest")
  //Ngxx.NgxRlCmdTest = viper.GetString("gwxReloadTest")
}

func (n *Ngx) SetServer(servName string) error {
  // 检查是否已有server
  n.mx.Lock()
  defer n.mx.Unlock()
  servs := n.NgxServs()
  fmt.Println(servs)
  for _, s := range servs {
    fmt.Println(s.Name)
    if servName == s.Name {
      fmt.Println(errors.New("已存在 " + servName + " server"))
      // fixme: 如果并发， 逻辑走到这里就return， 并没有lock就触发了 Unlock, 会出现 fatal error: sync: unlock of unlocked mutex错误
      return nil
    }
  }

  flagLine := n.GetFlagLine(n.ServerFlag)
  cfgCtx := "\t\t location /" + servName + " { proxy_pass http://" + servName + "/; }"
  // 添加节点信息
  newCfg := MakeCfg(flagLine, cfgCtx)
  //sm.Lock()
  // 替换原有内容
  err := n.ReplCfg(flagLine, newCfg)
  if err == nil {
    fmt.Println("添加server", servName, "成功")
  }
  return err
}

func (n *Ngx) SetUpstream(servName string, nodes []string) error {
  // 检查是否已有server
  n.mx.Lock()
  defer n.mx.Unlock()
  ups := n.NgxUpstrs()
  fmt.Println(ups)
  exit := false
  for _, u := range ups {
    fmt.Println(u.Name)
    if servName == u.Name {
      fmt.Println(errors.New("已存在 " + servName + " 负载"))
      exit = true
      //return nil
    }
  }
  flagLine2 := n.GetFlagLine(n.UpstreamFlag)
  cfgCtx2 := "\t upstream " + servName + " { "
  //nodesUp := ""
  for _, node := range nodes {
    cfgCtx2 += "server  " + node + " weight=1; "
  }
  cfgCtx2 += "}"

  if exit {
    oldUps := n.GetFlagLine(servName)
    err := n.ReplCfg(oldUps, cfgCtx2)
    if err != nil {
      fmt.Println("修改负载", servName, "成功")
    }
    return err
  } else {
    newCfg2 := MakeCfg(flagLine2, cfgCtx2)
    err := n.ReplCfg(flagLine2, newCfg2)
    if err == nil {
      fmt.Println("添加负载", servName, "成功")
    }
    return err
  }
}

func MakeCfg(flagLine string, cfgCtx string) string {
  // 整合旧的标识 和 新的配置内容
  newCfg := cfgCtx + "\n\n" + flagLine
  return newCfg
}

func (n *Ngx) ReplCfg(flagLine string, newCfg string) error {
  // 替换配置文件的内容
  // @param:  flag 更改的标识
  // @param:  cfg  更改后的内容
  ReplCtx(n.Conf, flagLine, newCfg)
  return nil
}

func (n *Ngx) GetFlagLine(flag string) string {
  // 返回标示所在行的内容
  fi, err := os.Open(n.Conf)
  CheckFatal(err, "打开nginx配置文件错误")
  defer fi.Close()

  br := bufio.NewReader(fi)
  for {
    a, _, c := br.ReadLine()
    if strings.Contains(string(a), flag) {
      return string(a)
    }
    if c == io.EOF {
      return ""
    }
  }
}

func (n *Ngx) ReloadNgx() error {
  // reload nginx
  fmt.Println("gwx reload cmd----", n.NgxRlCmd)
  OsExecOut(n.NgxRlCmd)
  //OsExecOut("ls")
  return nil
}

func (n *Ngx) NgxServs() []NgxServ {
  // 获取nginx所有的server
  var ngxServs []NgxServ
  ps := GetMatCtx(n.Conf, `location /(.*) { proxy_pass http://(.*)/; }`)
  for _, p := range ps {
    ngxServs = append(ngxServs, NgxServ{Name: strings.Trim(p[1], " ")})
  }
  return ngxServs
}

func (n *Ngx) NgxUpstrs() []NgxUpstr {
  // 获取nginx所有的upstream
  var ngxUpstr []NgxUpstr
  ps := GetMatCtx(n.Conf, `upstream(.*){ server(.*) weight=1; }`)
  for _, p := range ps {
    ngxUpstr = append(ngxUpstr, NgxUpstr{Name: strings.Trim(p[1], " ")})
  }
  return ngxUpstr
}




