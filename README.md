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
```
