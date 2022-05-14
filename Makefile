build:
	CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -ldflags "-s -w" -o tunnel cmd/tunnel/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o tunnel.exe cmd/tunnel/main.go

upx:
	upx tunnel
	upx tunnel.exe