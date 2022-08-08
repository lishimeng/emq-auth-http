package api

import (
	"github.com/kataras/iris/v12"
)

// SuperUser
// http://host:port/mqtt/superuser
// clientid=%c,username=%u
func SuperUser(ctx iris.Context) {

	userName := ctx.URLParam("username")

	// is superuser?

	var isSuperuser = false
	if userName == "admin_thingple_iot" {
		isSuperuser = true
	}

	if isSuperuser {
		ctx.StatusCode(RespCodeSuccess)
	} else {
		ctx.StatusCode(RespCodeUnauthorized)
	}

}
