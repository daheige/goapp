# goapp

    go api/grpc/worker

# layout参考

    https://www.jianshu.com/p/1c47d99f33ed
    https://draveness.me/golang-101/


# go version选择
    推荐使用go v1.16.15+版本
# golang linux环境安装

golang下载地址:
https://golang.google.cn/dl/

以go最新版本go1.16.15版本为例
https://golang.google.cn/dl/go1.20.2.linux-amd64.tar.gz
1. linux环境(centos,ubuntu操作系统)，下载
```shell
cd /usr/local/
    sudo wget https://golang.google.cn/dl/go1.16.15.linux-amd64.tar.gz
    sudo tar zxvf go1.16.15.linux-amd64.tar.gz
    # 创建golang需要的目录
    sudo mkdir ~/go
    sudo mkdir ~/go/bin
    sudo mkdir ~/go/src
    sudo mkdir ~/go/pkg
```
2. 设置环境变量vim ~/.bashrc 或者sudo vim /etc/profile
```shell
    export GOROOT=/usr/local/go
    export GOOS=linux
    export GOPATH=~/go
    export GOSRC=$GOPATH/src
    export GOBIN=$GOPATH/bin
    export GOPKG=$GOPATH/pkg
    
    #开启go mod机制
    export GO111MODULE=on

    #禁用cgo模块
    export CGO_ENABLED=0
    export GOPROXY=https://goproxy.cn,direct

    export PATH=$GOROOT/bin:$GOBIN:$PATH
```
:wq 保存退出
3. source ~/.bashrc 生效配置

# golang mac系统安装
只需要下载 https://golang.google.cn/dl/go1.16.15.darwin-amd64.pkg 然后点击下一步，下一步就可以安装完毕
环境变量配置：
vim ~/.bash_profile
```shell
    export GOROOT=/usr/local/go
    export GOOS=linux
    export GOPATH=~/go
    export GOSRC=$GOPATH/src
    export GOBIN=$GOPATH/bin
    export GOPKG=$GOPATH/pkg
    #开启go mod机制
    export GO111MODULE=on
    
    #禁用cgo模块
    export CGO_ENABLED=0
    
    #配置goproxy代理
    export GOPROXY=https://goproxy.cn,direct
    export PATH=$GOROOT/bin:$GOBIN:$PATH
```

:wq 退出即可，然后执行 source ~/.bash_profile 生效
