package lorawan

import (
	"github.com/lishimeng/app-starter/tool"
	"testing"
)

func TestGenOtaa(t *testing.T) {
	var s = tool.GetRandomString(16)
	t.Log(s, len(s), len(s) == 16*2)
}
