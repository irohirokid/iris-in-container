package main

import iris "github.com/kataras/iris/v12"

func main() {
	app := iris.Default()
	app.Use(myMiddleware)

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	app.Listen(":8080")
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
