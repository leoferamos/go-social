package models

type Profile struct {
	User      User    `json:"user"`
	Posts     []Posts `json:"posts"`
	Bio       string  `json:"bio"`
	Followers int     `json:"followers"`
	Following int     `json:"following"`
	CreatedAt string  `json:"created_at"`
}
