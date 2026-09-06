package generate

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type param struct {
	name      string
	cName     string
	direction paramDirection
	typ
}

func newParam(gen *gen, cursor clang.Cursor, commentParams commentParams) param {
	name := cursor.Spelling()
	param := param{
		name:  name,
		cName: "c_" + name,
		typ:   newTyp(gen, cursor.Type()),
	}

	if commentParam, ok := commentParams[param.name]; ok {
		param.direction = commentParam.direction
	}

	return param
}

func (param param) supported() (bool, string) {
	if param.isCallback {
		return true, ""
	}
	if param.isScalar {
		return true, ""
	}
	if param.isStruct && param.isPointer {
		return true, ""
	}
	if param.isString {
		return true, ""
	}
	if param.isPointer && param.isVoid {
		return true, ""
	}

	return false, fmt.Sprintf("param %s is %q", param.name, param.typ.typ.Spelling())
}

func (param param) isCount() bool {
	return param.name == "count" && param.direction == out && param.isPointer && param.isScalar
}

func (param param) cDecl() jen.Code {
	return jen.Id(param.cName).Add(param.typ.cDecl())
}

func (param param) goDecl() jen.Code {
	return jen.Id(param.name).Add(param.typ.goDecl())
}

func (param param) outGoDecl() jen.Code {
	if param.isCount() {
		return jen.Null()
	}

	if !param.isPointer {
		panic(fmt.Sprintf("out param %s is not a pointer : %s", param.name, param.typ.typ.Spelling()))
	}

	if param.isScalar {
		return param.scalar.goType
	}
	if param.isStruct {
		return jen.Id(param.struct_.name)
	}

	panic(fmt.Sprintf("out param %s type is unhandled : %s", param.name, param.typ.typ.Spelling()))
}

func (param param) cArgVar(g *jen.Group) {
	if param.isString {
		g.Id("c_" + param.name).Op(":=").Id("cString").Call(jen.Id(param.name))
	}
	if param.isCount() {
		g.Var().Id(param.name + "_").Add(param.scalar.goType)
		g.Id(param.name).Op(":=").Op("&").Id(param.name + "_")
	} else if param.direction == out {
		if param.isScalar {
			g.Var().Id(param.name).Add(param.scalar.goType)
		}
		if param.isStruct {
			g.Var().Id(param.name).Id(param.struct_.name)
		}
	}
}

func (param param) cArgName() jen.Code {
	if param.direction == out {
		return jen.New(jen.Op("&").Id(param.name))
	}

	if param.isCallback {
		return jen.Op("&").Id(param.name)
	}
	if param.isScalar {
		return jen.Op("&").Id(param.name)
	}
	if param.isStruct {
		return jen.Op("&").Id(param.name)
	}
	if param.isString {
		return jen.Op("&").Id("c_" + param.name)
	}
	if param.isPointer && param.isVoid {
		return jen.Op("&").Id(param.name)
	}

	panic("param type")
}

type params []param

func newParams(gen *gen, cursor clang.Cursor, commentParams commentParams) params {
	params := make(params, cursor.NumArguments())
	for i := range len(params) {
		params[i] = newParam(gen, cursor.Argument(uint32(i)), commentParams)
	}
	return params
}

func (params params) supported() (bool, string) {
	for _, param := range params {
		if supported, reason := param.supported(); !supported {
			return supported, reason
		}
	}

	return true, ""
}

func (params params) cDecl(g *jen.Group) {
	for _, param := range params {
		g.Add(param.cDecl())
	}
}

func (params params) goDecl(g *jen.Group) {
	for _, param := range params {
		if param.direction == in {
			g.Add(param.goDecl())
		}
	}
}

func (params params) outGoDecl(g *jen.Group) {
	for _, param := range params {
		if param.direction == out {
			g.Add(param.outGoDecl())
		}
	}
}

func (params params) cArgVars(g *jen.Group) {
	for _, param := range params {
		param.cArgVar(g)
	}
}

func (params params) count() (bool, *param) {
	for i, param := range params {
		if param.isCount() {
			return true, &params[i]
		}
	}
	return false, nil
}

func (params params) haveOut() bool {
	for _, param := range params {
		if param.direction == out {
			return true
		}
	}
	return false
}

func (params params) returnValues(g *jen.Group) {
	for _, param := range params {
		if param.direction == out && !param.isCount() {
			g.Id(param.name)
		}
	}
}
