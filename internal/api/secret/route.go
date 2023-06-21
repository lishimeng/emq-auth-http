package secret

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/emq-auth-http/internal/midware"
)

func Route(root iris.Party) {

	root.Post("/", midware.WithAuth(add)...)
	root.Get("/", midware.WithAuth(get)...)
	//root.Post("/", add)
	//root.Get("/{sn}", get)
}
