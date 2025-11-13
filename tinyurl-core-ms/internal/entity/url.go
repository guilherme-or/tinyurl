package entity

import (
	"net/url"
	"time"
)

type TinyURL struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"` // Generated short code
	RawURL    string    `json:"url"`
	URL       *url.URL  `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *TinyURL) SinceCreated() time.Duration {
	return time.Since(t.CreatedAt)
}
