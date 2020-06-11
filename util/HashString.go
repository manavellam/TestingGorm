package util

import (
	"crypto/sha256"
	"fmt"
)

//HashString hashes a given string
func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))

}
