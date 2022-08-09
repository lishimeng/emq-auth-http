package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/go-log"
)

type AuthResult string

type AuthReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type AuthResp struct {
	Result      AuthResult `json:"result,omitempty"`
	IsSuperuser bool       `json:"is_superuser,omitempty"`
}

const (
	AuthAllow  AuthResult = "allow"
	AuthDeny              = "deny"
	AuthIgnore            = "ignore"
)

// Auth
// http://host:port/mqtt/auth
// clientid=%c,username=%u,password=%P
func Auth(ctx iris.Context) {

	var err error
	var req AuthReq
	var resp AuthResp
	err = ctx.ReadJSON(&req)

	if err != nil {
		ctx.StatusCode(RespCodeUnauthorized)
	}

	log.Info("u:%s, p:%s", req.Username, req.Password)

	// TODO check
	resp.Result = AuthAllow
	resp.IsSuperuser = false
	_, _ = ctx.JSON(resp)
}
