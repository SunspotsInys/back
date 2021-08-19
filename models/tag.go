package models

type Tag struct {
	ID   uint64 `db:"id" json:"id,string"`
	Name string `db:"name" json:"name"`
}

type Tags struct {
	Tag
	Num int `db:"num" json:"num"`
}
