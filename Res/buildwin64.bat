set GOARCH=amd64
go build -ldflags "-s -w"
upx Res.exe

