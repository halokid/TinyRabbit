
package cmd

// Portainer is an instance of Portainer
type Portainer struct {
	URL       string
	username  string
	password  string
	token     string
	verbose   bool
	Endpoints []Endpoint
}

// NewPortainer returns a new Portainer instance
func NewPortainer() Portainer {
	// 从配置文件读取这些参数
	//url := viper.GetString("portainer_url") + "/api"
	// fixme: 统一更改配置, just for test
	url :=  "http://8.8.8.8:9000/api"
	//url :=  "http://192.168.1.129:9000/api"

	//username := viper.GetString("portainer_username")
	username := "admin"
	//password := viper.GetString("portainer_password")
	password := "xxx"

	portainer := Portainer{
		URL: url,
		username: username,
		password: password,
	}
	portainer.token = portainer.login()

	return portainer
}
