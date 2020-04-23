package controllers

import (
	"IRIS_WEB/controllers/http"
	"github.com/kataras/iris"
)

// 定义500错误处理函数
func err500(ctx iris.Context) {
	ctx.WriteString("CUSTOM 500 ERROR")
}

// 定义404错误处理函数
func err404(ctx iris.Context) {
	ctx.WriteString("CUSTOM 404 ERROR")
}

// 注入路由
func InnerRoute(app *iris.Application) {
	app.OnErrorCode(iris.StatusInternalServerError, err500)
	app.OnErrorCode(iris.StatusNotFound, err404)
	app.Get("/ping", func(ctx iris.Context) { ctx.WriteString("pong") })

	//app.Get("/user/p", http.ActionGetUser)  						// /path?user_id=1
	//app.Get("/user/{id:int min(1)}", http.ActionGetRestUser)  	// /path/1
	//app.Post("/user", http.ActionCreateUser)
	//app.Put("/user", http.ActionUpdateUser)

	app.Get("/get_user", http.ActionGetUser)              // /path?user_id=1
	app.Any("/user/{id:int min(1)}", http.ActionCrudUser) // /path/1(GET DELETE)
	app.Any("/user", http.ActionCrudUser)                 // /path(POST PUT)
	app.Any("/users", http.ActionGetAllUsers)
	app.Any("/users/auth", jwtHandler.Serve, http.ActionGetAllUsers)
}
