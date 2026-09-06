package generate

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type result struct {
	typ
	haveCountParam bool
	countParam     *param
}

func newResult(gen *gen, typ clang.Type, haveCountParam bool, countParam *param) result {
	return result{
		typ:            newTyp(gen, typ),
		haveCountParam: haveCountParam,
		countParam:     countParam,
	}
}

func (result result) supported() (bool, string) {
	if result.typ.isVoid {
		return true, ""
	}
	if result.isScalar {
		return true, ""
	}
	if result.isString {
		return true, ""
	}
	if result.isPointer && result.isStruct && !result.haveCountParam {
		return true, ""
	}
	if result.isCallback {
		return true, ""
	}

	return false, fmt.Sprintf("result type is %q", result.typ.typ.Spelling())
}

func (result result) resultVar(g *jen.Group) {
	if result.isVoid() {
		return
	}

	if result.haveCountParam {
		g.Var().Id("result").Add(result.typ.goDecl())
	} else if result.isString {
		g.Var().Id("result").Op("*").Byte()
	} else {
		g.Var().Id("result").Add(result.typ.goDecl())
	}
}

func (result result) isVoid() bool {
	return result.typ.isVoid && !result.isPointer
}

func (result result) returnVar(g *jen.Group) {
	if result.haveCountParam {
		g.Id("slice").Op(":=").Make(
			jen.Index().Add(result.scalar.goType),
			jen.Id(result.countParam.name+"_"),
		)
		g.Copy(
			jen.Id("slice"),
			jen.Qual("unsafe", "Slice").Call(
				jen.Id("result"),
				jen.Id(result.countParam.name+"_"),
			),
		)
	}
}

func (result result) returnValue(g *jen.Group) {
	if result.isVoid() {
		return
	}

	if result.haveCountParam {
		g.Id("slice")

	} else if result.isString {
		g.Id("goString").Call(jen.Id("result"))

	} else {
		g.Id("result")
	}
}

func (result result) returnValuePointer() jen.Code {
	if result.typ.isVoid {
		return jen.Nil()
	}

	return jen.Op("&").Id("result")
}

func (result result) goDecl() jen.Code {
	if result.haveCountParam {
		return jen.Index().Add(result.scalar.goType)
	}
	return result.typ.goDecl()
}
