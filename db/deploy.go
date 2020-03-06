package db

type Deploy struct {
  ID        int
  Name      string
  Env       string
}

func (d Deploy) TableName() string {
  return "deploy"
}

func GetDeploys() []Deploy {
  // 获取所有deploy信息
  var deploys []Deploy
  Db.Find(&deploys)
  return deploys
}

func GetOneDeploy(name string, gwEnv string) Deploy {
  var deploy Deploy
  Db.Where("name = ? and env = ?", name, gwEnv).Find(&deploy)
  return deploy
}

func AddDeploy(d *Deploy) error {
  if err := Db.Create(d).Error; err != nil {
    return err
  }
  return nil
}






