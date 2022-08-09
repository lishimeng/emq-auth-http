package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/emq-auth-http/internal/api/emq"
	"github.com/lishimeng/emq-auth-http/internal/api/lorawan"
)

func Route(app *iris.Application) {
	root := app.Party("/")
	emq.Router(root.Party("/emq"))
	lorawan.Router(root.Party("/lorawan"))
	return
}
