package password

import "github.com/kataras/iris/v12"

func Router(p iris.Party) {
	p.Post("/", gen)
}
