package encoding

import (
	"crypto/sha256"
	"fmt"
	"io"
)

func Encrypt(password string) string {
	h := sha256.New()
	io.WriteString(h, "BeeGo"+password+"GoFrame")
	return fmt.Sprintf("%x", h.Sum(nil))
}
