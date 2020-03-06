# 配置网关工具

### 设计
独立的配置网关地址的工具, 网关反代用nginx来做, 所以这个工具是管理nginx反代的工具

- 运行在 nginx 所在的系统航
- 以RPC协议来进行通信
- 管理nginx配置文件, 主要针对反代的配置
- 管理nginx配置重载



### 通信数据定义
```markdown

# 范例
reqData1 = {
  "act":    "setServer",
  "val":    <servName>
}


reqData2 = {
  "act":    "setUpstream",
  "val":    <servName>-!-<nodeIp>:<nodePort>
}

```

