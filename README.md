# 2SOMEone
2SOMEone is a brand new social media platform.  

# Proto Buffer complie  
Install `vscode-proto3` extension in vscode. And make setting file as follow:
`<root_path>/.vscode/settings.json`  
``` json
{
    "protoc": {
        "path": "<path of protoc.exe>",
        "compile_on_save": true,
        "options": [
            "--go_out=.",
            "--go_opt=paths=import",
            "--go-grpc_out=.",
            "--go-grpc_opt=paths=import",
            "${fileBaseName}/${fileBaseName}.proto",
            "--proto_path=${workspaceRoot}",
            "--proto_path=${workspaceRoot}/grpc",
        ]
    }
}
```

# UserService  
### Start a UserService grpc service  
> Defalut port: `http://127.0.0.1:50051`  

```sh  
$ make all
$ cd user
$ ls
```  

# Start a service
安装命令行工具
```sh
$ go install github.com/go-micro/cli/cmd/go-micro@v1.1.1
```

```
$ cd /user
$ go run main.go user.go
$ go run api/api.go  // API 依赖底层 go.micro.srv.greeter 服务
$ micro api --handler=api // 启动 API 网关处理 HTTP 请求，--handle 参数不能为空，否则可能报错
```
