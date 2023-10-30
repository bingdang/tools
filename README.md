# GO GUI Demo
## dev run
进入项目目录执行平常的运行命令即可

```
go run . 
```

## build

```cmd
go build -ldflags -H=windowsgui .
```

```shell
go install fyne.io/fyne/v2/cmd/fyne@latest
# windows 打包
fyne package -os windows -icon logo.png
# mac 打包
fyne package -os darwin -icon logo.png
# linux 打包
fyne package -os linux -icon myapp.png 
# mac编译windows
brew install mingw-w64
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC="x86_64-w64-mingw32-gcc" fyne package -os windows -icon logo.png
```
