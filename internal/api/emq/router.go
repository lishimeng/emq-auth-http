package emq

import (
	"github.com/kataras/iris/v12"
)

// emq中使用设备id为username
// id根据不同设备的规则创建，
// 密码规则：TODO
//

func Router(root iris.Party) {
	mqtt(root.Party("/v5.0"))
}
func mqtt(p iris.Party) {
	p.Post("/auth", Auth)
	p.Post("/auth/{client_id}", Auth)
	p.Post("/acl", Acl)
	p.Post("/acl/{client_id}", Acl)
}
