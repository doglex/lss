# lss: 
list **size** of file/**folder** in the current working directory,
while operating system can not provide folder size

+ sort from big to small 
+ also list count of files 


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
lss # display with path
lss -s # display without path
```

## demo
windows

![](dist/demo.png)

linux

![](dist/demo_linux.png)

