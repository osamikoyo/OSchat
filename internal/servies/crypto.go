package servies

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func GeterateJWTkey() string{
	uniqueInput := fmt.Sprintf("%s-%d", "some_unique_value", time.Now().UnixNano())
	hash := sha256.Sum256([]byte(uniqueInput))
	return fmt.Sprintf("%x", hash)
}