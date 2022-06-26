package models

import "time"

type Url struct {
	Id              string    `json:"id"`
	Url             string    `json:"url"`
	Slug            string    `json:"slug"`
	UserId          string    `json:"user_id"`
	RedirectEnabled bool      `json:"enabled"`
	Clicks          int64     `json:"clicks"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type UrlDetails struct {
	Url             string    `json:"url"`
	Slug            string    `json:"slug"`
	RedirectEnabled bool      `json:"enabled"`
	Clicks          int64     `json:"clicks"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
