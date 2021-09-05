package test

import (
	"testing"

	_ "github.com/SunspotsInys/thedoor/configs"

	"github.com/SunspotsInys/thedoor/db"
	"github.com/SunspotsInys/thedoor/models"
)

func TestUpdatPost(t *testing.T) {
	p := models.Post{
		ID:      16212958658533785605,
		Title:   "更新博文测试",
		Content: "# 测试\n\n供博文更新的测试 使用 \n\n  json",
		Public:  false,
		Top:     0,
	}
	tags := []models.Tag{
		{
			ID:   11313042263954685953,
			Name: "14234",
		},
	}
	err := db.UpdatePost(&p, &tags)
	if err != nil {
		t.Log(err)
	}
}
