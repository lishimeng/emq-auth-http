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
	mqtt(root.Party("/v5.0"))
}
func mqtt(p iris.Party) {
	p.Post("/auth", Auth)
	p.Post("/auth/{client_id}", Auth)
	p.Post("/acl", Acl)
	p.Post("/acl/{client_id}", Acl)
}
