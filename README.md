
## swag 中文官网
```
https://github.com/swaggo/swag/blob/master/README_zh-CN.md
```
## swag init command not found 
```
// 则运行 
go install github.com/swaggo/swag/cmd/swag
```


## swagger注解
```
// 修改完注释和配置要记得重新生成docs

swag init
```


## 编译发布到linux
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```
编译完后会生成一个geekdemo的go 文件，直接拷贝到服务器/usr/local/sunlight/go

## 部署到linux做成systemd服务
```
// geek.service

[Unit]
Description=dvs-basic
After=network-online.target
Wants=network-online.target

[Service]
# modify when deploy in prod env

Type=simple
#Environment="GIN_MODE=release"
ExecStart=/usr/local/sunlight/dvs/dvs-basic/dvs-basic
WorkingDirectory=/usr/local/sunlight/dvs/dvs-basic

Restart=always
RestartSec=1
StartLimitInterval=0

[Install]
WantedBy=multi-user.target

```

## 设置go服务
```
// 设置开机启动
systemctl enable geekdemo.service

// 启动服务
systemctl start geekdemo.service

// 停止服务
systemctl stop geekdemo.service

// 重新加载配置文件
sytemctl daemon-reload

// 查看服务状态
systemctl status geekdemo.service

// 查看运行日志
journalctl -u geekdemo -f
```