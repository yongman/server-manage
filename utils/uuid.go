package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateUUID() string {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return ""
	}
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo

	return hex.EncodeToString(uuid)
}
