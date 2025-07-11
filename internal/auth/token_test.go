package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go_social/config"

	"github.com/golang-jwt/jwt/v4"
)

func init() {
	config.SecretKey = []byte("test_secret")
}

func makeRequestWithToken(token string) *http.Request {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	return req
}

func TestCreateToken_And_ValidateToken_Success(t *testing.T) {
	token, err := CreateToken(123)
	if err != nil {
		t.Fatalf("CreateToken failed: %v", err)
	}
	req := makeRequestWithToken(token)
	if err := ValidateToken(req); err != nil {
		t.Errorf("ValidateToken failed for valid token: %v", err)
	}
}

func TestValidateToken_InvalidToken(t *testing.T) {
	req := makeRequestWithToken("invalid.token.value")
	err := ValidateToken(req)
	if err == nil {
		t.Error("ValidateToken should fail for invalid token")
	}
}

func TestValidateToken_EmptyToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil) // no Authorization header
	err := ValidateToken(req)
	if err == nil {
		t.Error("ValidateToken should fail for missing token")
	}
}

func TestExtractUserID_Success(t *testing.T) {
	token, err := CreateToken(456)
	if err != nil {
		t.Fatalf("CreateToken failed: %v", err)
	}
	req := makeRequestWithToken(token)
	uid, err := ExtractUserID(req)
	if err != nil {
		t.Fatalf("ExtractUserID failed: %v", err)
	}
	if uid != 456 {
		t.Errorf("expected userID 456, got %d", uid)
	}
}

func TestExtractUserID_InvalidToken(t *testing.T) {
	req := makeRequestWithToken("invalid.token.value")
	_, err := ExtractUserID(req)
	if err == nil {
		t.Error("ExtractUserID should fail for invalid token")
	}
}

func TestExtractUserID_EmptyToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	_, err := ExtractUserID(req)
	if err == nil {
		t.Error("ExtractUserID should fail for missing token")
	}
}

func TestExtractUserID_NoUserIDClaim(t *testing.T) {
	claims := jwt.MapClaims{
		"authorized": true,
		"exp":        float64(jwt.NewNumericDate(jwt.TimeFunc().Add(time.Hour)).Unix()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(config.SecretKey)
	req := makeRequestWithToken(tokenString)
	_, err := ExtractUserID(req)
	if err == nil || err.Error() != "userID is not a valid number" {
		t.Errorf("expected userID is not a valid number error, got: %v", err)
	}
}

func TestExtractUserID_InvalidUserIDType(t *testing.T) {
	claims := jwt.MapClaims{
		"authorized": true,
		"exp":        float64(jwt.NewNumericDate(jwt.TimeFunc().Add(time.Hour)).Unix()),
		"userID":     "notanumber",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(config.SecretKey)
	req := makeRequestWithToken(tokenString)
	_, err := ExtractUserID(req)
	if err == nil || err.Error() != "userID is not a valid number" {
		t.Errorf("expected userID is not a valid number error, got: %v", err)
	}
}

func TestExtractToken_NoBearer(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "tokenonly")
	token := extractToken(req)
	if token != "" {
		t.Errorf("expected empty token, got %q", token)
	}
}

func TestReturnKeyFunc_InvalidMethod(t *testing.T) {
	token := jwt.New(jwt.SigningMethodRS256)
	_, err := returnKeyFunc(token)
	if err == nil || err.Error() != fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]) {
		t.Errorf("expected unexpected signing method error, got: %v", err)
	}
}
