package cmd
/*
生成创建容器的map

{
    "Image": "xxxxxx:5000/xxxx",
    "Env": [],
    "Cmd": [],
    "MacAddress": "",
    "ExposedPorts": {
        "8080/tcp": {}
    },
    "HostConfig": {
        "RestartPolicy": {
            "Name": "no"
        },
        "PortBindings": {
            "8080/tcp": [
                {
                    "HostPort": "5001"
                }
            ]
        },
        "PublishAllPorts": false,
        "Binds": [],
        "AutoRemove": false,
        "NetworkMode": "bridge",
        "Privileged": false,
        "Runtime": "",
        "ExtraHosts": [],
        "Devices": [],
        "CapAdd": [
            "AUDIT_WRITE",
            "CHOWN",
            "DAC_OVERRIDE",
            "FOWNER",
            "FSETID",
            "KILL",
            "MKNOD",
            "NET_BIND_SERVICE",
            "NET_RAW",
            "SETFCAP",
            "SETGID",
            "SETPCAP",
            "SETUID",
            "SYS_CHROOT"
        ],
        "CapDrop": [
            "AUDIT_CONTROL",
            "BLOCK_SUSPEND",
            "DAC_READ_SEARCH",
            "IPC_LOCK",
            "IPC_OWNER",
            "LEASE",
            "LINUX_IMMUTABLE",
            "MAC_ADMIN",
            "MAC_OVERRIDE",
            "NET_ADMIN",
            "NET_BROADCAST",
            "SYSLOG",
            "SYS_ADMIN",
            "SYS_BOOT",
            "SYS_MODULE",
            "SYS_NICE",
            "SYS_PACCT",
            "SYS_PTRACE",
            "SYS_RAWIO",
            "SYS_RESOURCE",
            "SYS_TIME",
            "SYS_TTY_CONFIG",
            "WAKE_ALARM"
        ]
    },
    "NetworkingConfig": {
        "EndpointsConfig": {
            "bridge": {
                "IPAMConfig": {
                    "IPv4Address": "",
                    "IPv6Address": ""
                }
            }
        }
    },
    "Labels": {},
    "name": "xx-dev",
    "OpenStdin": false,
    "Tty": false,
    "Volumes": {}
}
 */

func makeCtMap(imgName string, cName string, cPort string, hostPort string) string {
  /**
  data := make(map[string]interface{})
  data["Image"] = imgName
  data["Env"] = "[]"
  data["Cmd"] = "[]"
  data["MacAddress"] = ""
  data["ExposedPorts"] = `{"8080/tcp": {}"}`
  data["HostConfig"] = `{"RestartPolicy": {"Name": "no"}"}, "PortBindings": {"8080/tcp": [{
                        "HostPort": "5001"}]}`
  data["PublishAllPorts"] = false
  data["Runtime"] = ""
  data["ExtraHosts"] = "[]"
  */

  s := `{"Image":"` + imgName + `","Env":["TZ=Asia/Shanghai"],"Cmd":[],"MacAddress":"","ExposedPorts":{"` + cPort + `/tcp":{}},"HostConfig":{"RestartPolicy":{"Name":"no"},"PortBindings":{"` + cPort + `/tcp":[{"HostPort":"` + hostPort + `"}]},"PublishAllPorts":false,"Binds":[],"AutoRemove":false,"NetworkMode":"bridge","Privileged":false,"Runtime":"","ExtraHosts":[],"Devices":[],"CapAdd":["AUDIT_WRITE","CHOWN","DAC_OVERRIDE","FOWNER","FSETID","KILL","MKNOD","NET_BIND_SERVICE","NET_RAW","SETFCAP","SETGID","SETPCAP","SETUID","SYS_CHROOT"],"CapDrop":["AUDIT_CONTROL","BLOCK_SUSPEND","DAC_READ_SEARCH","IPC_LOCK","IPC_OWNER","LEASE","LINUX_IMMUTABLE","MAC_ADMIN","MAC_OVERRIDE","NET_ADMIN","NET_BROADCAST","SYSLOG","SYS_ADMIN","SYS_BOOT","SYS_MODULE","SYS_NICE","SYS_PACCT","SYS_PTRACE","SYS_RAWIO","SYS_RESOURCE","SYS_TIME","SYS_TTY_CONFIG","WAKE_ALARM"]},"NetworkingConfig":{"EndpointsConfig":{"bridge":{"IPAMConfig":{"IPv4Address":"","IPv6Address":""}}}},"Labels":{},"name":"` + cName + `","OpenStdin":false,"Tty":false,"Volumes":{}}`
  return s

}







