package requests

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

// MakeAuthenticatedRequest sends an HTTP request with the user's authentication cookie.
func MakeAuthenticatedRequest(r *http.Request, method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	cookie, _ := cookies.Read(r)
	req.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
