package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/emq-auth-http/internal/api/emq"
	"github.com/lishimeng/emq-auth-http/internal/api/lorawan"
	"github.com/lishimeng/emq-auth-http/internal/api/password"
	"github.com/lishimeng/emq-auth-http/internal/api/secret"
)

func Route(app *iris.Application) {
	root := app.Party("/")
	emq.Router(root.Party("/emq"))
	lorawan.Router(root.Party("/lorawan"))
	password.Router(root.Party("/password"))
	secret.Route(root.Party("/secret"))
	return
}
