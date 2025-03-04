package function

import (
	"html/template"
	"math/rand"

	"github.com/spf13/cast"
)

var mathFuncs = map[string]interface{}{
	"add1": func(i interface{}) int64 { return cast.ToInt64(i) + 1 },
	"add": func(i ...interface{}) int64 {
		var a int64 = 0
		for _, b := range i {
			a += cast.ToInt64(b)
		}
		return a
	},
	"sub": func(a, b interface{}) int64 { return cast.ToInt64(a) - cast.ToInt64(b) },
	"div": func(a, b interface{}) int64 { return cast.ToInt64(a) / cast.ToInt64(b) },
	"mod": func(a, b interface{}) int64 { return cast.ToInt64(a) % cast.ToInt64(b) },
	"mul": func(a interface{}, v ...interface{}) int64 {
		val := cast.ToInt64(a)
		for _, b := range v {
			val = val * cast.ToInt64(b)
		}
		return val
	},
	"rand": func(min, max int) int { return rand.Intn(max-min) + min },
}

func GetMathFuncs() template.FuncMap {
	return template.FuncMap(mathFuncs)
}
