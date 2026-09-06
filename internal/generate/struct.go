package generate

import (
	"slices"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type struct_ struct {
	gen     *gen
	cursor  clang.Cursor
	cName   string
	name    string
	comment comment
}

func (struct_ struct_) generate(file file) {
	file.Comment(struct_.comment.text())
	file.Type().Id(struct_.name).StructFunc(func(g *jen.Group) {
		g.Id("_").Qual("structs", "HostLayout")

		struct_.cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
			if cursor.Kind() == clang.Cursor_FieldDecl {
				struct_.generateField(g, cursor)
			}

			return clang.ChildVisit_Continue
		})
	})
}

func (struct_ struct_) generateField(g *jen.Group, cursor clang.Cursor) {
	g.Comment(struct_.gen.newComment(cursor).text())

	cName := cursor.Spelling()
	name := goName(cName)
	if slices.Contains([]string{"GLFWgammaramp", "GLFWimage"}, struct_.cName) {
		name = goNameUnexported(cName) // accessor methods will provide access to the unexported fields
	}

	typ := newTyp(struct_.gen, cursor.Type())

	if typ.isPointer {
		if typ.isScalar {
			g.Id(name).Op("*").Add(typ.scalar.goType)
		} else {
			g.Id(name).Qual("unsafe", "Pointer")
		}

	} else if typ.isArray {
		g.Id(name).Index(jen.Lit(typ.arraySize)).Add(typ.scalar.goType)

	} else if typ.isScalar {
		g.Id(name).Add(typ.scalar.goType)

	} else if typ.isCallback {
		g.Id(name).Id(typ.callback.name)

	} else {
		fatalf("UNSUPPORTED : field %s of type %q", goName(cursor.Spelling()), cursor.Type().Spelling())
	}
}

type structs []struct_

func (structs structs) generate() {
	file := newFile("struct.go", "glfw")
	defer file.save()

	for _, struct_ := range structs {
		struct_.generate(file)
		file.Line()
	}
}

func (structs structs) createComments(gen gen) {
	for i, struct_ := range structs {
		structs[i].comment = gen.newComment(struct_.cursor)
	}
}

func (structs structs) find(cName string) (*struct_, bool) {
	for i, struct_ := range structs {
		if struct_.cName == cName {
			return &(structs[i]), true
		}
	}

	return nil, false
}

func (structs structs) findByGoName(name string) (*struct_, bool) {
	for i, struct_ := range structs {
		if struct_.name == name {
			return &structs[i], true
		}
	}

	return nil, false
}
