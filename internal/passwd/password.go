package passwd

import (
	"fmt"
	"github.com/lishimeng/emq-auth-http/internal/db/model"
	"github.com/lishimeng/emq-auth-http/internal/util"
)

func Generate(plaintext string, ds model.DeviceSecret) (r string) {
	r = genPass(plaintext, ds)
	return
}

func Verify(plaintext string, ds model.DeviceSecret) (r bool) {
	encoded := genPass(plaintext, ds)
	r = encoded == ds.Password
	return
}

func genPass(password string, ds model.DeviceSecret) (r string) {
	s := ds.CreateTime.UnixNano()
	plain := fmt.Sprintf("%d.%s_%d", s, password, s)
	bs, err := gen([]byte(plain))
	if err != nil {
		return
	}
	r = util.BytesToHex(bs)
	return
}
