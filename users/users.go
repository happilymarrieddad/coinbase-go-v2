package coinbasev2

import "time"

// Country - country
type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// ReadPermission - read permissions
type ReadPermission struct {
	TimeZone       string     `json:"time_zone"`
	NativeCurrency string     `json:"native_currency"`
	BitcoinUnit    string     `json:"bitcoin_unit"`
	Country        Country    `json:"country"`
	CreatedAt      *time.Time `json:"created_at"`
}

// EmailPermission - email permission
type EmailPermission struct {
	Email string `json:"email"`
}

// User - coinbase user
type User struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	Username        string         `json:"username"`
	ProfileLocation string         `json:"profile_location"`
	ProfileBio      string         `json:"profile_bio"`
	ProfileURL      string         `json:"profile_url"`
	AvatarURL       string         `json:"avatar_url"`
	Resource        string         `json:"resource"`
	ResourcePath    string         `json:"resource_path"`
	ReadPermission  ReadPermission `json:"read_permission"`
}
