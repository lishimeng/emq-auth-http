package passwd

import (
	"github.com/lishimeng/app-starter/digest"
	"github.com/lishimeng/emq-auth-http/internal/db/model"
)

func Generate(plaintext string, ds model.DeviceSecret) (r string) {
	r = digest.Generate(plaintext, ds.CreateTime.UnixNano())
	return
}

func Verify(plaintext string, ds model.DeviceSecret) (r bool) {
	r = digest.Verify(plaintext, ds.Password, ds.CreateTime.UnixNano())
	return
}
