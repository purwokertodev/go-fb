package fb

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestNewFacebook(t *testing.T) {
	f := NewFacebook("001", "secret", "666666", "v1.1", "http://localhost:3000/auth/facebook/callback", false, 1000)
	if f == nil {
		t.Error("failed create new Facebook")
	}
}

func TestGetSecretProof(t *testing.T) {
	f := NewFacebook("001", "secret", "666666", "v1.1", "http://localhost:3000/auth/facebook/callback", false, 1000)

	key := []byte("secret")
	h := hmac.New(sha256.New, key)
	h.Write([]byte("666666"))
	expectedHex := hex.EncodeToString(h.Sum(nil))

	if expectedHex != f.getSecretProof() {
		t.Error("failed to verify secretProof")
	}
}
