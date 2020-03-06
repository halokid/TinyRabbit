package db

import "strings"

type Pod struct {
  ID      int
  Did     int
  Name    string
  Eid     int
  Ename   string
  Addr    string
  Env     string
  Cid     string
}

func (p Pod) TableName() string {
  return "pod"
}

func GetPods(did int) []Pod {
  // 获取所有pod信息
  var pods []Pod
  Db.Where("did = ?", did).Find(&pods)
  return pods
}

func AddPods(ps []Pod) error {
  /**
  p := Pod{
    Did:    did,
    Name:   name,
    Eid:    eid,
    Cid:    cid,
  }
  */
  for _, p := range ps {
    if err := Db.Create(&p).Error; err != nil {
      return err
    }
  }
  return nil
}

func AddPod(p Pod) error {
  if err := Db.Create(&p).Error; err != nil {
    return err
  }
  return nil
}

func GetPodsBySvc(svcName string, gwEnv string) []Pod {
  var pods []Pod
  Db.Where("name = ? and env = ?", svcName, gwEnv).Find(&pods)
  return pods
}

func UpdatePod(p *Pod) error {
  err := Db.Save(p).Error
  return err
}

func GetPodsByPort(hostPort string) []Pod {
  // 根据端口信息获取pods
  var pods []Pod
  Db.Where("addr LIKE ?", "%:" + hostPort).Find(&pods)
  return pods
}

func GetPodPortBySvc(svc string) string {
  // 根据svc获取pod的port, svc一定存在的情况下
  var pods []Pod
  Db.Where("name = ?", svc).Find(&pods)
  if len(pods) == 0 {
    return ""
  }
  addr := pods[0].Addr
  addrSpl := strings.Split(addr, ":")
  if len(addrSpl) > 1 {
    return addrSpl[1]
  } else {
    return ""
  }
}





