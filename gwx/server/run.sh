go build -o gwx_server
sleep 2
nohup ./gwx_server  > gwx_serv.log 2>&1 &