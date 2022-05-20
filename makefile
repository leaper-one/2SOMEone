GOARCH = amd64

user:
	cd user
	go build -o ./user/user ./user/main.go

user.exe:
	go build -o ./user/user.exe ./user/main.go

build-linux: user

build-win: user.exe