
# download package

## go env
``` 
# ~/.bashrc
export GOROOT=/usr/local/go
export GOPATH=$GOROOT/bin:~/go
export PATH=$PATH:$GOROOT:$GOPATH
```

# 文档
运行一下命令查看 [查看文档](http://localhost:8080)

```
godoc -http=:8080
```