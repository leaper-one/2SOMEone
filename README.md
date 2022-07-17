# 2SOMEone
2SOMEone is a brand new social media platform.  


## [安装命令行工具](https://go-zero.dev/cn/docs/goctl/installation)
```sh
$ go install github.com/zeromicro/go-zero/tools/goctl@latest
```
## 构建镜像  
根目录下：
```sh
$ docker build -t message-rpc:v1 -f /rpc/message-rpc/Dockerfile .
$ docker build -t <image-name>:<tag> -f /<your_path>/Dockerfile .
```

## 项目结构
```
├─api -- api 网关
├─core -- 数据库定义
├─rpc -- rpc 层
├─service -- 业务代码
├─store -- CRUD
└─util -- 相关工具
```