GOARCH = amd64
GIT_COMMIT=$(shell git rev-list -1 HEAD)
LDFLAGS = -ldflags "-X main.GitCommit=${GIT_COMMIT}"


linux-user:
	CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=${GOARCH} go build -o ./user/linux_${GOARCH}/user ./user/main.go

windows-user:
	CGO_ENABLED=0 GO111MODULE=on GOOS=windows GOARCH=${GOARCH} go build -o ./user/windows_${GOARCH}/user.exe ./user/main.go

build-user: linux-user windows-user

build-all: build-user

all: build-all