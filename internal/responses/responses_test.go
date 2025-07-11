package responses

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestJSON checks if JSON writes correct status, header and body.
func TestJSON(t *testing.T) {
	rec := httptest.NewRecorder()
	data := map[string]string{"foo": "bar"}
	JSON(rec, http.StatusOK, data)

	assertStatus(t, rec, http.StatusOK)
	assertContentType(t, rec, "application/json")

	var resp map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if resp["foo"] != "bar" {
		t.Errorf("expected foo=bar, got foo=%s", resp["foo"])
	}
}

// TestJSONError checks if JSONError writes correct status and error message.
func TestJSONError(t *testing.T) {
	rec := httptest.NewRecorder()
	errMsg := "something went wrong"
	JSONError(rec, http.StatusBadRequest, errors.New(errMsg))

	assertStatus(t, rec, http.StatusBadRequest)
	assertContentType(t, rec, "application/json")

	var resp struct {
		Error string `json:"error"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode error response: %v", err)
	}
	if resp.Error != errMsg {
		t.Errorf("expected error message %q, got %q", errMsg, resp.Error)
	}
}

// TestJSON_EncodingError checks if JSON handles encoding errors gracefully.
func TestJSON_EncodingError(t *testing.T) {
	rec := httptest.NewRecorder()
	data := make(chan int) // cannot be encoded to JSON
	JSON(rec, http.StatusOK, data)
	assertStatus(t, rec, http.StatusOK)
	assertContentType(t, rec, "application/json")
	// No need to decode body, just ensure no panic
}

// Helper functions
func assertStatus(t *testing.T, rec *httptest.ResponseRecorder, want int) {
	t.Helper()
	if rec.Code != want {
		t.Errorf("expected status %d, got %d", want, rec.Code)
	}
}

func assertContentType(t *testing.T, rec *httptest.ResponseRecorder, want string) {
	t.Helper()
	if ct := rec.Header().Get("Content-Type"); ct != want {
		t.Errorf("expected Content-Type %s, got %s", want, ct)
	}
}
