package models

type TagList struct {
	ID  uint64 `db:"id"`
	TID uint64 `db:"tid"`
	PID uint64 `db:"pid"`
}
