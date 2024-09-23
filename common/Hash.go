package common

import (
	"crypto/sha256"
	"encoding/hex"
)

type Sha265Hash struct{}

func NewSha265Hash() *Sha265Hash {
	return &Sha265Hash{}
}
func (h *Sha265Hash) Hash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
