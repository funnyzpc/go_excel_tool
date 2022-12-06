
## excel批量解密工具
具体使用请见[README(使用说明).md](./document/README(使用说明).md)

## 添加代理
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

# M1程序 (arm) 不可用
set CGO_ENABLED=0
set GOOS=darwin
set GOARCH=arm64
go build -o ./go_excel_tool_macos_m1   ./main.go
go build -o ./go_excel_tool_macos_m1 -ldflags "-w -s -X main.env=prod" ./main.go
upx --backup go_excel_tool_macos_m1

# Mac程序 (intel)
set CGO_ENABLED=0
set GOOS=darwin
set GOARCH=amd64
go build -o ./go_excel_tool_macos_intel -ldflags "-w -s -X main.env=prod" ./main.go
upx --backup go_excel_tool_macos_intel

# window程序
set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -o ./go_excel_tool_window.exe -ldflags "-w -s -X main.env=prod" ./main.go
upx --backup go_excel_tool_window.exe


# linux程序
set CGO_ENABLED=0
set GOARCH=amd64
set GOOS=linux
go build -o ./go_excel_tool_linux  ./main.go
upx --backup go_excel_tool_linux
