# lss: 
list size of file/folder in the current working directory,
while operating system can not provide folder size
## how to test
```
go run main.go
```

## how to deploy
```
go build -ldflags "-w"
# for linux
env GOOS=linux GOARCH=amd64 go build -ldflags "-w"
add lss.exe folder to envrionment $PATH 
lss
```

## demo
windows

![](dist/demo.png)

linux

![](dist/demo_linux.png)

