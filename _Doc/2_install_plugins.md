github.com/mdempsky/gocode 
github.com/uudashr/gopkgs/cmd/gopkgs 
github.com/ramya-rao-a/go-outline 
github.com/acroca/go-symbols 
golang.org/x/tools/cmd/guru 
golang.org/x/tools/cmd/gorename 
github.com/go-delve/delve/cmd/dlv 
github.com/stamblerre/gocode 
github.com/rogpeppe/godef 
github.com/sqs/goreturns 
golang.org/x/lint/golint 

1 安装vscode 的golang插件
** 提前下好需要的插件
  cd $GOPATH/src/golang.org/x
  git clone git@github.com:golang/tools.git 
  git clone https://github.com/golang/lint.git
2 打开一个.go文件,会出现安装对应插件的提示,install all
3 有些应为是在golang.org上出现无法下载的情况
