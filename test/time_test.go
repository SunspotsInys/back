package test

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	t1, _ := time.Parse("2006-01-02 15:04:05", "2021-07-17 00:00:00")
	t.Log(t1.UnixNano())
}
