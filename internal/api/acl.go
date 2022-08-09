package api

import "github.com/kataras/iris/v12"

type AclResult string

type AclReq struct {
	Username string
	Topic    string
	Action   string
}

type AclResp struct {
	Result AuthResult `json:"result,omitempty"`
}

const (
	AclAllow  AclResult = "allow"
	AclDeny             = "deny"
	AclIgnore           = "ignore"
)

func Acl(ctx iris.Context) {

	// 不控制acl
	var resp AclResp
	resp.Result = AclIgnore
	_, _ = ctx.JSON(resp)
}
