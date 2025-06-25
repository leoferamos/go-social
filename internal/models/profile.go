package models

type PublicUser struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	CreatedAt string `json:"created_at"`
}

type Profile struct {
	User      PublicUser `json:"user"`
	Posts     []Posts    `json:"posts"`
	Followers int        `json:"followers"`
	Following int        `json:"following"`
}
