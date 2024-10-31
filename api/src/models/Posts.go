package models

import "time"

type Posts struct {
	Id         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorId   uint64    `json:"author_id,omitempty"`
	AuthorNick uint64    `json:"author_nick,omitempty"`
	LikeCount  uint64    `json:"like_count"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
