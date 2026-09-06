package generate

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type function struct {
	gen        *gen
	cursor     clang.Cursor
	cName      string
	name       string
	result     result
	params     params
	varName    string
	cifVarName string
	comment    comment
	isMethod   bool
}

func newFunction(gen *gen, cursor clang.Cursor) function {
	cName := cursor.Spelling()
	comment := gen.newComment(cursor)
	params := newParams(gen, cursor, comment.params)
	haveCountParam, countParam := params.count()

	fn := function{
		gen:        gen,
		cursor:     cursor,
		cName:      cName,
		name:       goName(cName),
		result:     newResult(gen, cursor.ResultType(), haveCountParam, countParam),
		params:     params,
		varName:    "func_" + cName,
		cifVarName: "cif_" + cName,
	}

	fn.isMethod = len(fn.params) > 0 && fn.params[0].isPointer && fn.params[0].isStruct

	if fn.isMethod {
		fn.name = strings.Replace(fn.name, fn.params[0].struct_.name, "", 1)
	}

	return fn
}

func (fn function) generate(file file) {
	if supported, reason := fn.supported(); !supported {
		file.Commentf("UNSUPPORTED %s : %s", fn.cName, reason)
		return
	}

	file.Comment(fn.comment.text())

	file.
		Func().
		Do(func(s *jen.Statement) { // receiver
			if fn.isMethod {
				s.Parens(jen.Add(fn.params[0].goDecl()))
			}
		}).
		Id(fn.name).
		ParamsFunc(func(g *jen.Group) {
			if fn.isMethod {
				fn.params[1:].goDecl(g)
			} else {
				fn.params.goDecl(g)
			}
		}).
		ParamsFunc(func(g *jen.Group) {
			fn.params.outGoDecl(g)
			g.Add(fn.result.goDecl())
		}).
		BlockFunc(func(g *jen.Group) {
			fn.result.resultVar(g)
			fn.params.cArgVars(g)

			g.
				List(jen.Id("_"), jen.Id("err")).
				Op(":=").
				Qual(goffiImportPath, "CallFunction").
				CallFunc(func(g *jen.Group) {
					g.Line().Id(fn.cifVarName)
					g.Line().Id(fn.varName)
					g.Line().Qual("unsafe", "Pointer").Parens(fn.result.returnValuePointer())
					g.Line().Index().Qual("unsafe", "Pointer").ValuesFunc(func(g *jen.Group) {
						for _, param := range fn.params {
							g.Line().Qual("unsafe", "Pointer").Parens(param.cArgName())
						}
						g.Line()
					})
					g.Line()
				})

			g.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Panic(jen.Id("err")),
			)

			if !fn.result.isVoid() || fn.params.haveOut() {
				fn.result.returnVar(g)
				g.Return().ListFunc(func(g *jen.Group) {
					fn.params.returnValues(g)
					fn.result.returnValue(g)
				})
			}
		})
}

func (fn function) supported() (bool, string) {
	if supported, reason := fn.result.supported(); !supported {
		return supported, reason
	}
	return fn.params.supported()
}

type functions []function

func (functions functions) generate() {
	functions.generateGoffi()
	functions.generateGo()
}

func (functions functions) generateGo() {
	file := newFile("function.go", "glfw")
	defer file.save()

	for _, function := range functions {
		function.generate(file)
		file.Line()
	}
}

func (functions functions) createComments(gen gen) {
	for i, function := range functions {
		functions[i].comment = gen.newComment(function.cursor)
	}
}

func (functions functions) find(cName string) (*function, bool) {
	for i, function_ := range functions {
		if function_.cName == cName {
			return &(functions[i]), true
		}
	}

	return nil, false
}
