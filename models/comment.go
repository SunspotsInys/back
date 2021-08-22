package models

import "time"

type Comment struct {
	ID         uint64    `db:"id"         json:"id,string"`
	PID        uint64    `db:"pid"        json:"pid,string,omitempty"`
	FID        uint64    `db:"fid"        json:"fid,string,omitempty"`
	Content    string    `db:"content"    json:"content"`
	CreateTime time.Time `db:"createtime" json:"createTime"`
	Name       string    `db:"name"       json:"name"`
	EMail      string    `db:"email"      json:"email"`
	Site       string    `db:"site"       json:"site"`
}

type Comments struct {
	Comment
	Children []Comment `json:"children"`
}
