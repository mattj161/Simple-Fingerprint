package Fingerprint

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

var fingerprint = getHash()

func FingerPrint() string {
	return fingerprint
}

func getHash() string {
	str := baseId() + biosId() + diskId()
	hasher := md5.New()
	hasher.Write([]byte(str))
	var b []byte
	hasher.Write(b)
	bt := hex.EncodeToString(hasher.Sum(nil))
	var s string
	for i := 0; i < len(bt); i++ {
		s += string(bt[i])
		if (i+1) != len(bt) && (i+1)%7 == 0 {
			s += "-"
		}
	}
	return strings.ToUpper(s)
}
