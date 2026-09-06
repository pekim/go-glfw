package generate

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var titleCaser = cases.Title(language.English, cases.NoLower)
var lowerCaser = cases.Lower(language.English, cases.NoLower)

func goName(name string) string {
	name = strings.TrimPrefix(name, "GLFW_")
	name = strings.TrimPrefix(name, "GLFW")
	name = strings.TrimPrefix(name, "glfw")
	name = titleCaser.String(name)
	return name
}

func goNameUnexported(name string) string {
	return lowerCaser.String(goName(name))
}
