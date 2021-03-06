package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/cache"
	"github.com/kataras/iris/v12/middleware/recover"
	"io"
	"iris-project/middleware"
	"os"
	"time"
)

func main() {
	app := iris.New()

	// recover 中间件从任何异常中恢复，如果有异常，则写入500状态码（服务器内部错误）。
	app.Use(recover.New())

	// 同时写文件日志与控制台日志
	f := middleware.NewLogFile()
	defer f.Close()
	app.Logger().SetOutput(io.MultiWriter(f, os.Stdout))
	logConfig := middleware.LogHandler()
	app.Use(logConfig)

	// Load all templates from the "./templates" folder
	// where extension is ".html" and parse them
	// using the standard `html/template` package
	// Reload 用来配置是否动态加载html模板.
	app.RegisterView(iris.HTML("./templates", ".html").Reload(true))

	app.Get("/", cache.Handler(60*time.Second), func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./templates/hello.html
		ctx.View("hello.html")
	})

	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

	// 从settings包里载入参数配置
	config := iris.WithConfiguration(iris.TOML("./conf/iris.tml"))

	app.Listen(":8080", config)
	// Start the server using a network address.
}
