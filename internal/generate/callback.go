package generate

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type callback struct {
	gen     *gen
	cursor  clang.Cursor
	cName   string
	name    string
	result  result
	params  params
	comment comment
}

func newCallback(gen *gen, cursor clang.Cursor) callback {
	cName := cursor.Spelling()
	name := goName(cName)
	name = strings.Replace(name, "fun", "Callback", 1)
	name = strings.Replace(name, "proc", "Proc", 1)

	var params params
	cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() == clang.Cursor_ParmDecl {
			params = append(params, newParam(gen, cursor, nil))
		}
		return clang.ChildVisit_Continue
	})

	cb := callback{
		gen:    gen,
		cursor: cursor,
		cName:  cName,
		name:   name,
		result: newResult(gen, cursor.TypedefDeclUnderlyingType().PointeeType().ResultType(), false, nil),
		params: params,
	}

	return cb
}

func (callback callback) supported() (bool, string) {
	if supported, reason := callback.result.supported(); !supported {
		return supported, reason
	}
	return callback.params.supported()
}

func (callback callback) generate(file file) {
	// callback type
	file.Comment(callback.comment.text())
	file.Type().Id(callback.name).Uintptr()

	if supported, reason := callback.supported(); !supported {
		file.Commentf("UNSUPPORTED %s : %s", callback.cName, reason)
		return
	}

	// new callback function
	file.
		Func().
		Id(callback.name + "New").
		Params(
			jen.Id("callback").Func().
				ParamsFunc(callback.params.goDecl).
				Params(jen.Add(callback.result.goDecl())),
		).
		Id(callback.name).
		Block(
			jen.
				Return().
				Id(callback.name).Parens(
				jen.Qual(goffiImportPath, "NewCallback").Call(
					jen.
						Func().
						ParamsFunc(callback.params.cDecl).
						BlockFunc(func(g *jen.Group) {
							g.Id("callback").CallFunc(func(g *jen.Group) {
								for _, param := range callback.params {
									if param.isString {
										g.Id("goString").Call(jen.Id(param.cName))
									} else if param.isScalar {
										g.Add(param.scalar.goType).Parens(jen.Id(param.cName))
									} else {
										g.Id(param.cName)
									}
								}
							})
						}),
				),
			),
		)
}

type callbacks []callback

func (callbacks callbacks) generate() {
	file := newFile("callback.go", "glfw")
	defer file.save()

	for _, callback := range callbacks {
		callback.generate(file)
		file.Line()
	}
}

func (callbacks callbacks) createComments(gen gen) {
	for i, callback := range callbacks {
		callbacks[i].comment = gen.newComment(callback.cursor)
	}
}

func (callbacks callbacks) find(name string) (*callback, bool) {
	for i, callback_ := range callbacks {
		if callback_.cName == name {
			return &(callbacks[i]), true
		}
	}

	return nil, false
}
