package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/go-log"
)

type AuthReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// Auth
// http://host:port/mqtt/auth
// clientid=%c,username=%u,password=%P
func Auth(ctx iris.Context) {

	var err error
	var req AuthReq
	err = ctx.ReadJSON(&req)

	if err != nil {
		ctx.StatusCode(RespCodeUnauthorized)
	}

	log.Info("u:%s, p:%s", req.Username, req.Password)

	// TODO check
	ctx.StatusCode(RespCodeSuccess)
}
