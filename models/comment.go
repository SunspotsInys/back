package models

import "time"

type Comment struct {
	ID         uint64    `db:"id"         json:"id,string"`
	PID        uint64    `db:"pid"        json:"pid,string,omitempty"`
	FID        uint64    `db:"fid"        json:"fid,string"`
	Content    string    `db:"content"    json:"content"`
	CreateTime time.Time `db:"createtime" json:"createtime"`
	CName      string    `db:"cname"      json:"cname"`
	CEMail     string    `db:"cemail"     json:"cemail"`
	CSite      string    `db:"csite"      json:"csite"`
}

type Comments struct {
	Comment
	ChildComment []Comment `json:"childComment"`
}
