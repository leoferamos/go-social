package security

import (
	"testing"
)

func TestHashPasswordAndCheckPasswordHash(t *testing.T) {
	password := "SecurityTest123"

	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	if err := CheckPasswordHash(string(hashed), password); err != nil {
		t.Errorf("CheckPasswordHash failed for correct password: %v", err)
	}

	if err := CheckPasswordHash(string(hashed), "SecurityTestWrongPassword"); err == nil {
		t.Error("CheckPasswordHash should have failed for incorrect password, but it didn't")
	}
}
