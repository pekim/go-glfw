package generate

import (
	"github.com/dave/jennifer/jen"
)

const goffiImportPath = "github.com/go-webgpu/goffi/ffi"
const typesImportPath = "github.com/go-webgpu/goffi/types"

var pointerTypeDescriptor = jen.Qual(typesImportPath, "PointerTypeDescriptor")
var voidTypeDescriptor = jen.Qual(typesImportPath, "VoidTypeDescriptor")

func (fn function) generateGoffiVar(file file) {
	file.Var().Id(fn.varName).Qual("unsafe", "Pointer")
	file.Var().Id(fn.cifVarName).Op("=").Op("&").Qual(typesImportPath, "CallInterface").Block()
	file.Line()
}

func (fn function) generateGoffiGetSymbol(g *jen.Group) {
	g.
		List(jen.Id(fn.varName), jen.Id("err")).
		Op("=").
		Qual(goffiImportPath, "GetSymbol").Call(
		jen.Id("handle"),
		jen.Lit(fn.cName),
	)

	g.If(
		jen.Id("err").Op("!=").Nil()).Block(
		jen.Return(jen.Id("err")),
	)
}

func (fn function) generateGoffiPrepareCallInterface(g *jen.Group) {
	if supported, _ := fn.supported(); !supported {
		return
	}

	g.Id("err").Op("=").Qual(goffiImportPath, "PrepareCallInterface").Call(
		jen.Line().Id(fn.cifVarName),
		jen.Line().Qual(typesImportPath, "DefaultCall"),
		jen.Line().Add(fn.result.typeDescriptor()),
		jen.Line().Index().Op("*").Qual(typesImportPath, "TypeDescriptor").ValuesFunc(func(g *jen.Group) {
			for _, param := range fn.params {
				g.Line().Add(param.typeDescriptor())
			}
			g.Line()
		}),
	)

	g.If(
		jen.Id("err").Op("!=").Nil()).Block(
		jen.Return(jen.Id("err")),
	)
}

func (functions functions) generateGoffi() {
	file := newFile("goffi.go", "glfw")
	defer file.save()

	for _, fn := range functions {
		if supported, _ := fn.supported(); supported {
			fn.generateGoffiVar(file)
		}
	}
	functions.generateGoffiInit(file)
}

func (functions functions) generateGoffiInit(file file) {
	file.Var().Id("initialised").Op("=").False()
	file.Line()

	file.Comment(`
		Initialise loads the glfw library, and initialises all of the the api functions.

		Initialise must be called before calling any other function in the package.
		If an api function is called before this has been called, a panic will result.

		Do not confuse this function with the [Init] api function.
	`)
	file.Func().Id("Initialise").Params().Error().BlockFunc(func(g *jen.Group) {
		g.If(jen.Id("initialised")).Block(
			jen.Return(jen.Nil()),
		)
		g.Defer().Func().Params().Block(
			jen.
				Id("initialised").
				Op("=").
				True(),
		).Call()
		g.Line()

		g.
			List(
				jen.Id("handle"), jen.Id("err"),
			).
			Op(":=").
			Qual(goffiImportPath, "LoadLibrary").Call(jen.Lit("libglfw.so"))
		g.If(
			jen.Id("err").Op("!=").Nil()).Block(
			jen.Return(jen.Id("err")),
		)
		g.Line()

		getVersion, foundGetVersion := functions.find("glfwGetVersion")
		if !foundGetVersion {
			panic("failed to find glfwGetVersion function")
		}
		getVersion.generateGoffiGetSymbol(g)
		getVersion.generateGoffiPrepareCallInterface(g)
		g.Line()

		// major, minor, _ := GetVersion()
		g.List(
			jen.Id("major"),
			jen.Id("minor"),
			jen.Id("_"),
		).Op(":=").
			Id("GetVersion").Call()
		// version := float64(major) + (float64(minor) / 10)
		g.
			Id("version").
			Op(":=").
			Float64().Parens(jen.Id("major")).
			Op("+").
			Parens(jen.Float64().Parens(jen.Id("minor")).Op("/").Lit(10))
		g.Line()

		for _, fn := range functions {
			if supported, _ := fn.supported(); supported {
				if fn.cName == "glfwGetVersion" {
					continue
				}

				since := float64(fn.comment.sinceMajor) + (float64(fn.comment.sinceMinor) / 10)
				g.
					// if version >= 3.4 {
					If().Id("version").Op(">=").Lit(since).
					BlockFunc(func(g *jen.Group) {
						fn.generateGoffiGetSymbol(g)
						fn.generateGoffiPrepareCallInterface(g)
					})
				g.Line()
			}
		}

		g.Return(jen.Nil())
	})
}
