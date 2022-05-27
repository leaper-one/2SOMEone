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
