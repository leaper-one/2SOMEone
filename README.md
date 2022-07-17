# 2SOMEone
2SOMEone is a brand new social media platform.  

## (安装命令行工具)[https://go-zero.dev/cn/docs/goctl/installation]
```sh
$ go install github.com/zeromicro/go-zero/tools/goctl@latest
```
## 构建镜像  
根目录下：
```sh
$ docker build -t message:v1 -f 2someone/message/rpc/message/Dockerfile .
$ docker build -t message:<tag> -f 2someone/<your_path>/Dockerfile .
```
