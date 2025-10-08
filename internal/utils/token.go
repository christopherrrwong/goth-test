package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateToken() string {
	uniqueID1 := generateUniqID()
	hash1 := md5Hash(uniqueID1)
	uniqueID2 := generateUniqID()
	hash2 := md5Hash(uniqueID2)

	return hash1 + hash2
}

func generateUniqID() string {
	now := time.Now().UnixNano() / 1000
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return fmt.Sprintf("%d", now)
	}

	return fmt.Sprintf("%d%x", now, b)
}

func md5Hash(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
