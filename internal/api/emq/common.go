package emq

import (
	"github.com/kataras/iris/v12"
	"time"
)

const (
	RespCodeSuccess      = 200
	RespCodeNotFound     = 404
	RespCodeUnauthorized = 401
)

const (
	RespMsgNotFount = "not found"
	RespMsgIdNum    = "id must be a int value"
)

func ResponseJSON(ctx iris.Context, j interface{}) {
	_, _ = ctx.JSON(j)
}

const (
	DefaultTimeFormatter = "2006-01-02:15:04:05"
	DefaultCodeLen       = 16
)

func FormatTime(t time.Time) (s string) {
	s = t.Format(DefaultTimeFormatter)
	return
}
