package util

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"math/rand"
	"time"
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

func GetRandomString(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}
