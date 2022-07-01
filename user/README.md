## Run  
windows
```sh
$ set MICRO_REGISTRY_ADDRESS=127.0.0.1:2379
$ set MICRO_REGISTRY=etcd
$ go run main.go user.go  
```

linux
```sh
$ MICRO_REGISTRY_ADDRESS=127.0.0.1:2379 \
  MICRO_REGISTRY=etcd \
  go run main.go user.go  
```