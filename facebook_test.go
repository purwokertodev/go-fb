package fb

import (
	"testing"
)

func testNewFacebook(t *testing.T) {
	f := NewFacebook("001", "secret", "v1.1", "http://localhost:3000/auth/facebook/callback", false, 1000)
	if f == nil {
		t.Error("failed create new Facebook")
	}
}
