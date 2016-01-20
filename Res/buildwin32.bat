set GOARCH=386
go build -ldflags "-s -w"
upx res.exe