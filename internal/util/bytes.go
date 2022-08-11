package util

import (
	"bytes"
	"encoding/hex"
	"strconv"
	"strings"
)

func BytesToHexStr(data []byte) string {
	buffer := new(bytes.Buffer)
	for index, b := range data {

		s := strconv.FormatInt(int64(b&0xff), 16)
		if index > 0 {
			buffer.WriteString(" ")
		}
		if len(s) == 1 {
			buffer.WriteString("0x0")
		} else {
			buffer.WriteString("0x")
		}
		buffer.WriteString(s)
	}
	return buffer.String()
}

func BytesToHex(data []byte) string {
	return strings.ToLower(hex.EncodeToString(data))
}
