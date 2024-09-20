## 区块链非物质文化遗产数字版权保护平台
<div>

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.22-9cf)
![Release](https://img.shields.io/badge/release-1.0.0-green.svg)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
</div>
一个基于Golang Gin的非遗版权保护平台

- 后端基于 [golang](https://go.dev/) + [gin](https://gin-gonic.com/zh-cn/)
- 前端基于 [VUE3](https://vuejs.org/)
## 项目结构 🧐

| 子项目名 | 项目路径                                |
|------|-------------------------------------|
| 后端服务 | [/my_finished/back](./back)         |
| 智能合约 | [/my_finished/contract](./contract) |
| 前端服务 | [/my_finished/front](./front)       |
## 技术栈
- mysql
- redis
- rabbitmq
- gin
- jwt

## 一键部署
1. 直接拉去镜像
```bash
docker pull registry.cn-hangzhou.aliyuncs.com/ljn_docker_hub/finish:heritage
```
2. 构建容器
```bash
docker run -it --name=heritage -p 8080:8080 -p 8081:8081 -p 23306:3306 -p 26379:6379 -p 15672:15672 -d registry.cn-hangzhou.aliyuncs.com/ljn_docker_hub/finish:heritage 
```
3. 进入容器自行修改/root/deploy/back/conf/conf.json 合约基础配置
## 系统截图

