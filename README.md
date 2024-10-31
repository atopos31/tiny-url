# 一个demo短链接服务
![](./image/短链接服务.png)
# Docker 部署
修改docker-compose.yaml文件，设置你的。
```bash
 # /bin/bash
 make docker
```
使用Nginx将域名反向代理到[SERVER_PORT]