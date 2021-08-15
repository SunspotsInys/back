package logs

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/rs/zerolog"
)

type h1 struct {
}

func (h1) Run(e *zerolog.Event, _ zerolog.Level, _ string) {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		files := strings.Split(file, "/")
		funcNames := strings.Split(runtime.FuncForPC(pc).Name(), ".")
		file = fmt.Sprintf("%s/%s:%s:%d",
			files[len(files)-2],
			files[len(files)-1],
			funcNames[len(funcNames)-1],
			line,
		)
		e.Str("caller", file)
	}
}
