package lorawan

import (
	"github.com/kataras/iris/v12"
)

func Router(root iris.Party) {
	otaa(root.Party("/otaa"))
}
func otaa(p iris.Party) {
	p.Post("/{dev_eui}", genOTAAKey)
	p.Get("/{dev_eui}", getOTAAKey)
}
