package emq

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/emq-auth-http/internal/passwd"
	"github.com/lishimeng/emq-auth-http/internal/util"
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
		resp.Result = AuthDeny
		util.ResponseJSON(ctx, resp)
		return
	}

	log.Info("u:%s, p:%s", req.Username, req.Password)

	u, err := getUser(req.Username)
	if err != nil {
		log.Info(err)
		resp.Result = AuthDeny
		resp.IsSuperuser = false
		tool.ResponseJSON(ctx, resp)
		return
	}

	ok := passwd.Verify(req.Password, u)
	if !ok {
		resp.Result = AuthDeny
	} else {
		resp.Result = AuthAllow
	}
	util.ResponseJSON(ctx, resp)
}
