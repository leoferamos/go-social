package models

// User represents a user in the system.
type User struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Bio       string `json:"bio,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	BannerURL string `json:"banner_url,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
