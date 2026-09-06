package generate

import (
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type macroConstant struct {
	cursor  clang.Cursor
	name    string
	cName   string
	value   string
	comment comment
}

func (constant macroConstant) generate(file file) {
	file.Comment(constant.comment.text())
	file.Const().Id(constant.name).Op("=").Id(constant.value)
}

type macroConstants []macroConstant

func (gen *gen) newMacroConstant(cursor clang.Cursor) {
	_, _, _, start := cursor.Extent().Start().FileLocation()
	_, _, _, end := cursor.Extent().End().FileLocation()
	extent := gen.contents[start:end]
	value := strings.TrimPrefix(extent, cursor.Spelling())
	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, "GLFW_", "")

	if value == "" {
		return
	}

	cName := cursor.Spelling()
	name := goName(cName)
	gen.macroConstants = append(gen.macroConstants, macroConstant{
		cursor: cursor,
		name:   name,
		cName:  cName,
		value:  value,
	})
}

func (constants macroConstants) createComments(gen gen) {
	for i, constant := range constants {
		constants[i].comment = gen.newComment(constant.cursor)
	}
}

func (constants macroConstants) find(cName string) (*macroConstant, bool) {
	for i, constant := range constants {
		if constant.cName == cName {
			return &constants[i], true
		}
	}
	return nil, false
}

func (constants macroConstants) generate() {
	file := newFile("constant.go", "glfw")
	defer file.save()

	for _, constant := range constants {
		constant.generate(file)
	}
}
