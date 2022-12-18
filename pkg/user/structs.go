package user

// Name represents a Facebook user's name in different formats.
type Name struct {
	Full   string `json:"full_name"`
	First  string `json:"first_name"`
	Middle string `json:"middle_name"`
	Last   string `json:"last_name"`
}

// Profile represents a Facebook user.
type Profile struct {
	Name         *Name  `json:"name"`
	Username     string `json:"username"`
	RegisteredAt int64  `json:"registration_timestamp"`
}
