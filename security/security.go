package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

//PassSHA256 you can use this function to hash your password whit sha256 and encoding to hexadecimal finaly return a string format
func PassSHA256(payload string) string {

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte("secret"))

	// Write Data to it
	h.Write([]byte(payload))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
