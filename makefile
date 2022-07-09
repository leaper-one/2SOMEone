GOARCH = amd64
GIT_COMMIT=$(shell git rev-list -1 HEAD)
LDFLAGS = -ldflags "-X main.GitCommit=${GIT_COMMIT}"


linux-user:
	CGO_ENABLED=1 GO111MODULE=on GOOS=linux GOARCH=${GOARCH} go build -o ./user/linux_${GOARCH}/user ./user/main.go ./user/user.go
	CGO_ENABLED=1 GO111MODULE=on GOOS=linux GOARCH=${GOARCH} go build -o ./user/linux_${GOARCH}/api ./user/api/api.go

windows-user:
	CGO_ENABLED=1 GO111MODULE=on GOOS=windows GOARCH=${GOARCH} go build -o ./user/windows_${GOARCH}/user.exe ./user/main.go ./user/user.go
	CGO_ENABLED=1 GO111MODULE=on GOOS=windows GOARCH=${GOARCH} go build -o ./user/windows_${GOARCH}/api.exe ./user/api/api.go

build-user: linux-user windows-user

linux-message:
	CGO_ENABLED=1 GO111MODULE=on GOOS=linux GOARCH=${GOARCH} go build -o ./message/linux_${GOARCH}/message ./message/main.go ./message/message.go
	CGO_ENABLED=1 GO111MODULE=on GOOS=linux GOARCH=${GOARCH} go build -o ./message/linux_${GOARCH}/api ./message/api/api.go

windows-message:
	CGO_ENABLED=1 GO111MODULE=on GOOS=windows GOARCH=${GOARCH} go build -o ./message/windows_${GOARCH}/message.exe ./message/main.go ./message/message.go
	CGO_ENABLED=1 GO111MODULE=on GOOS=windows GOARCH=${GOARCH} go build -o ./message/windows_${GOARCH}/api.exe ./message/api/api.go

build-message: linux-message windows-message

build-all: build-user build-message

all: build-all
