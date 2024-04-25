package lab3

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
)

func GenerateUniqueToken(username string) int {
	rawToken := username + strconv.Itoa(rand.Intn(20))
	hash := sha256.Sum256([]byte(rawToken))
	hexToken := hex.EncodeToString(hash[:])
	intToken, _ := strconv.ParseInt(hexToken[:15], 16, 64)
	return int(intToken)
}
