package models

type Profile struct {
	User      map[string]interface{} `json:"user"`
	Posts     []Post                 `json:"posts"`
	Bio       string                 `json:"bio"`
	Followers int                    `json:"followers"`
	Following int                    `json:"following"`
	CreatedAt string                 `json:"created_at"`
}
