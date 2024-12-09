# build 生成文件到 bin/ 目录下
build:
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build  -o bin/feishu-gin
