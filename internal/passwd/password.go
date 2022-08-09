package passwd

import (
	"encoding/base64"
	"fmt"
	"github.com/lishimeng/emq-auth-http/internal/db/model"
	"strings"
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
	bs, err := hmacSha256(plain)
	if err != nil {
		return
	}
	r = strings.ToLower(base64.StdEncoding.EncodeToString(bs))
	return
}
