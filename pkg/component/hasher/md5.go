package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

// Hash performs MD5 encryption on a byte slice
func Hash(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
