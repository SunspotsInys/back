package models

type Tag struct {
	ID   uint64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
