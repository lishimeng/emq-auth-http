package passwd

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash"
)

var hmacProvider hash.Hash // TODO 并发处理

func init() {
	InitHmac("UBNKU_THINGPLE^&*VHV")
}

func InitHmac(key string) {
	hmacProvider = hmac.New(sha256.New, []byte(key))
}

func hmacSha256(content string) (r []byte, err error) {
	hmacProvider.Reset()
	_, err = hmacProvider.Write([]byte(content))
	if err != nil {
		return
	}
	r = hmacProvider.Sum(nil)
	return
}
