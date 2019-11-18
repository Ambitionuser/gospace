#####################################go语法说明########################################
#环境准备
setx GOPATH F:\workspace\gospace\golangbeego
go get -x github.com/gin-gonic/gin
go get -x github.com/swaggo/swag/cmd/swag
go get -x github.com/swaggo/gin-swagger
go get -x github.com/go-xorm/xorm

#下载放入golang.org/x/文件夹中
go get github.com/golang/tools.git 放入golang.org/x中
go get github.com/golang/net.git
go get github.com/golang/text.git
go get github.com/alecthomas
#进入安装swag的根目录中进行编译（同main）
go install
#编译后再目录中生成swag.exe文件
#初始化文件（生成docs目录）
swag init
#

#golang+beego

#下载beego
go get github.com/astaxie/bee环境go
go get github.com/beego/bee
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

参考网站
https://blog.csdn.net/weixin_34248023/article/details/
https://blog.csdn.net/qq_34857250/article/details/102493228