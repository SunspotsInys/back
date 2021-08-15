package test

import (
	"testing"

	"github.com/SunspotsInys/thedoor/utils"
)

func TestJwtGen(t *testing.T) {
	str, err := utils.GenToken("Username")
	if err != nil {
		panic(err)
	}
	t.Log(str)
}

func TestJwtParse(t *testing.T) {
	s := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlVzZXJuYW1lIiwiZXhwIjoxNjI2NjEyOTA4LCJpc3MiOiJ0aGVkb29yIn0.Sqzita1pJZARiAHKyVXdb-tMvJGVtXOnsIZDJKZgKN0"
	s1 := utils.ParseToken(s)
	t.Log(s1)
}
