package api

import "github.com/kataras/iris/v12"

// Acl
// http://host:port/mqtt/acl
// access=%A,username=%u,clientid=%c,ipaddr=%a,topic=%t,mountpoint=%m
func Acl(ctx iris.Context) {

	// 不控制acl
	ctx.StatusCode(RespCodeSuccess)
	_, _ = ctx.WriteString("ignore")
}
