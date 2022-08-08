package api

import (
	"github.com/kataras/iris/v12"
)

func Route(app *iris.Application) {
	root := app.Party("/")
	router(root)
	return
}

func router(root iris.Party) {
	mqtt(root.Party("/mqtt"))

}
func mqtt(p iris.Party) {
	p.Post("/auth", Auth)
	p.Post("/acl", Acl)
	p.Post("/superuser", SuperUser)
}
