package encrypt

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

//DoubleSha256Encryption 加密
func DoubleSha256Encryption(s string) (string, error) {
	ms := bytes.NewBufferString(s)
	h := sha256.New()
	_, err := h.Write(ms.Bytes())
	if err != nil {
		return "", err
	}
	c := h.Sum(nil)
	dq := []byte{c[2], c[4], c[6], c[10], c[12], c[24], c[18], c[7], c[5], c[21], c[30], c[11], c[9]}
	_, err = h.Write(dq)
	if err != nil {
		return "", err
	}
	ss := fmt.Sprintf("%x", h.Sum(nil))
	return ss, nil
}
