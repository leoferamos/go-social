package models

// PublicUser represents a public view of a user's profile.
type PublicUser struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url,omitempty"`
	BannerURL string `json:"banner_url,omitempty"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	CreatedAt string `json:"created_at"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
}

// Profile represents a user's profile in the application.
type Profile struct {
	User        PublicUser `json:"user"`
	IsFollowing bool       `json:"is_following"`
	Posts       []Post     `json:"posts"`
}
