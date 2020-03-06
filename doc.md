```markdown

# 启动一个容器三连

1. pull image
POST http://10.1.1.40:9000/api/endpoints/11/docker/images/create?fromImage=smicro_demo&tag=latest
request payload：  {"fromImage":"micro_demo","tag":"latest"}


2. create容器
POST http://10.1.1.40:9000/api/endpoints/11/docker/containers/create?name=xx-dev
request payload：  {"Image":"10.1.1.9:5000/micro_demo","Env":[],"Cmd":[],"MacAddress":"","ExposedPorts":{"8080/tcp":{}},"HostConfig":{"RestartPolicy":{"Name":"no"},"PortBindings":{"8080/tcp":[{"HostPort":"5001"}]},"PublishAllPorts":false,"Binds":[],"AutoRemove":false,"NetworkMode":"bridge","Privileged":false,"Runtime":"","ExtraHosts":[],"Devices":[],"CapAdd":["AUDIT_WRITE","CHOWN","DAC_OVERRIDE","FOWNER","FSETID","KILL","MKNOD","NET_BIND_SERVICE","NET_RAW","SETFCAP","SETGID","SETPCAP","SETUID","SYS_CHROOT"],"CapDrop":["AUDIT_CONTROL","BLOCK_SUSPEND","DAC_READ_SEARCH","IPC_LOCK","IPC_OWNER","LEASE","LINUX_IMMUTABLE","MAC_ADMIN","MAC_OVERRIDE","NET_ADMIN","NET_BROADCAST","SYSLOG","SYS_ADMIN","SYS_BOOT","SYS_MODULE","SYS_NICE","SYS_PACCT","SYS_PTRACE","SYS_RAWIO","SYS_RESOURCE","SYS_TIME","SYS_TTY_CONFIG","WAKE_ALARM"]},"NetworkingConfig":{"EndpointsConfig":{"bridge":{"IPAMConfig":{"IPv4Address":"","IPv6Address":""}}}},"Labels":{},"name":"xx-dev","OpenStdin":false,"Tty":false,"Volumes":{}}


3. start容器
POST http://10.1.1.40:9000/api/endpoints/11/docker/containers/3ee6b1fee994b338c3a1b1518d872f5c2ff071023869bd70b47c1dcb3f5b317e/start

 */
// api/endpoints/11/docker/images/create?fromImage=8.8.8.8:5000%2Fxxxx&tag=latest

```