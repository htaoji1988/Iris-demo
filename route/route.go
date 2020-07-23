package route

import (
	"github.com/kataras/iris/v12"
	"iris-project/controllers"
)

var animalRoutes iris.Party

// 为外提供一个设置主路由的方法
func (u *controllers.Animal) SetUserRouter(app *iris.Application, path string) {
	testRoutes = app.Party(path)
	//路由分发,这里再次路由分发，将功能块再次细化
	setLoginRoute()
	setUserInfoRoute()
}
