## Air自动重载
[github.com/cosmtrek/air]( github.com/cosmtrek/air )

>它足够稳定、功能齐全、活跃更新
> 

### 安装air
```bash
GO111MODULE=on  go install github.com/cosmtrek/air@latest
```

安装成功后的效果：
```bash
➜ air -v                         

  __    _   ___  
 / /\  | | | |_) 
/_/--\ |_| |_| \_ , built with Go 

```

### 使用air
在项目根目录下运行以下命令：
```bash
air
```
开始自动重载
后续开发`air`是一直处于打开的状态。

### air配置信息
我们可以使用 .air.toml 文件来配置 air 的行为。下面的配置项 include_ext 里，加了两个后缀：env 和 gohtml 
```toml
# https://github.com/cosmtrek/air/blob/master/air_example.toml TOML 格式的配置文件

# 工作目录
# 使用 . 或绝对路径，请注意 `tmp_dir` 目录必须在 `root` 目录下
root = "."
tmp_dir = "tmp"

[build]
  # 由`cmd`命令得到的二进制文件名
  # Windows平台示例：bin = "./tmp/main.exe"
  bin = "./tmp/main"
  # 只需要写你平常编译使用的shell命令。你也可以使用 `make`
  # Windows平台示例: cmd = "go build -o ./tmp/main.exe ."
  cmd = "go build -o ./tmp/main ."
  # 如果文件更改过于频繁，则没有必要在每次更改时都触发构建。可以设置触发构建的延迟时间
  delay = 1000
  # 忽略这些文件扩展名或目录
  exclude_dir = ["assets", "tmp", "vendor","public/uploads"]
  # 忽略以下文件
  exclude_file = []
  # 使用正则表达式进行忽略文件设置
  exclude_regex = []
  # 忽略未变更的文件
  exclude_unchanged = false
  # 监控系统链接的目录
  follow_symlink = false
  # 自定义参数，可以添加额外的编译标识，例如添加 GIN_MODE=release
  full_bin = ""
  # 监听以下指定目录的文件
  include_dir = []
  # 监听以下文件扩展名的文件.
  include_ext = ["go", "tpl", "tmpl", "html", "gohtml", "env"]
  # kill 命令延迟
  kill_delay = "0s"
  # air的日志文件名，该日志文件放置在你的`tmp_dir`中
  log = "build-errors.log"
  # 在 kill 之前发送系统中断信号，windows 不支持此功能
  send_interrupt = false
  # error 发生时结束运行
  stop_on_error = true

[color]
  # 自定义每个部分显示的颜色。如果找不到颜色，使用原始的应用程序日志。
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  # 显示日志时间
  time = false

[misc]
  # 退出时删除tmp目录
  clean_on_exit = false
```

## 集成gin

导入gin
```bash
go get github.com/gin-gonic/gin
```

编写`main.go`

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化gin实例
	r := gin.Default()

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {
		// 以JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	// 运行服务
	r.Run()
}

```

`air`自动重载，并监听`8080`端口

可以访问地址：[http://localhost:8080](http://localhost:8080)

修改一下`main.go`

```go
// new 一个 Gin Engine 实例
r := gin.New()

// 注册中间件
r.Use(gin.Logger(), gin.Recovery())
```
上面两句等于我们初始化的
```go
r := gin.Default()
```

```go
r.Run(":8000")
```
可以重新监听新的端口

使用`go mod tidy`
处理一些不要的warning

## 初始化路由
新建两个目录：`boostrap`和`routes`
`bootstrap`目录存放程序初始化的代码；`routes`存放我们所有的路由文件

## 配置方案
```bash
go get github.com/spf13/cast

go get github.com/spf13/viper
```

## 集成GORM
```bash
go get gorm.io/gorm
```

引入驱动
```bash
go get gorm.io/driver/mysql
go get gorm.io/driver/sqlite
```