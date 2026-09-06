package generate

import "github.com/dave/jennifer/jen"

func (gen gen) generateStructTest() {
	file := newFile("internal/structtest/struct.go", "structtest")
	defer file.save()

	file.CgoPreamble(`#include "../generate/glfw3.h"`)

	file.
		Var().Id("structs").
		Op("=").
		Index().
		Struct(
			jen.Id("name").String(),
			jen.Id("c").Any(),
			jen.Id("go_").Any(),
		).
		ValuesFunc(func(g *jen.Group) {
			for i, struct_ := range gen.structs {
				g.
					Do(func(s *jen.Statement) {
						if i == 0 {
							s.Line()
						}
					}).
					Values(jen.DictFunc(func(d jen.Dict) {
						d[jen.Id("name")] = jen.Lit(struct_.cName)
						d[jen.Id("c")] = jen.Qual("C", struct_.cName).Block()
						d[jen.Id("go_")] = jen.Qual("github.com/pekim/glfw", struct_.name).Block()
					}))
			}
			g.Line()
		})
}
