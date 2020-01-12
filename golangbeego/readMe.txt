#####################################go语法说明########################################
#环境准备
setx GOPATH F:\workspace\gospace\golangbeego
go env -w GO111MODULE=on

#golang+beego
go get github.com/astaxie/beego环境go
go list -m -versions github.com/astaxie/beego
go get github.com/beego/bee
go list -m -versions github.com/astaxie/beego
添加路径：path中添加bee的环境变量
F:\workspace\gospace\golangbeego\bin

go get gopkg.in/olivere/elastic.v5
#进入golangbeego/src下面创建项目名称
bee api esee-api
#进入golangbeego/src/esee-api
bee run -gendoc=true -downdoc=true

#go的测试框架(Convey是作为外层框架，Monkey可以为函数、方法等打桩)
go get github.com/smartystreets/goconvey
go get github.com/golang/tools.git 放入golang.org/x中
go get github.com/bouk/monkey

在测试代码目录下运行go test命令：
go test -v
在测试用例源码目录下运行goconvey：实现web端查看
goconvey -port 8081