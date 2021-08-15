package models

import "time"

type Post struct {
	ID         uint64    `db:"id"         json:"id,string"`
	Title      string    `db:"title"      json:"title"`
	Content    string    `db:"content"    json:"content"`
	CreateTime time.Time `db:"createTime" json:"createTime"`
	Public     bool      `db:"public"     json:"public"`
	Top        int64     `db:"top"        json:"top"` // 默认为0，即不置顶，
}

type PostWithTag struct {
	Post
	Tags []Tag `json:"tags"`
}

type PostSimplicity struct {
	ID         uint64    `db:"id"         json:"id,string"`
	Title      string    `db:"title"      json:"title"`
	CreateTime time.Time `db:"createTime" json:"createTime"`
}
