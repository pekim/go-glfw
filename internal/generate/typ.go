package generate

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type scalar struct {
	cType          jen.Code
	goType         jen.Code
	typeDescriptor jen.Code
}

var scalars = map[clang.CursorKind]scalar{
	clang.Type_Bool: {
		cType:          jen.Byte(),
		goType:         jen.Id("Bool"),
		typeDescriptor: jen.Qual(typesImportPath, "UInt8TypeDescriptor"),
	},

	clang.Type_Float: {
		cType:          jen.Float32(),
		goType:         jen.Id("Float"),
		typeDescriptor: jen.Qual(typesImportPath, "FloatTypeDescriptor"),
	},

	clang.Type_Double: {
		cType:          jen.Float64(),
		goType:         jen.Id("Double"),
		typeDescriptor: jen.Qual(typesImportPath, "DoubleTypeDescriptor"),
	},

	clang.Type_Int: {
		cType:          jen.Int32(),
		goType:         jen.Id("Int"),
		typeDescriptor: jen.Qual(typesImportPath, "SInt32TypeDescriptor"),
	},

	// clang.Type_Long: {
	// cType:          jen.Int64(),
	// goType:         jen.Int64(),
	// 	typeDescriptor: jen.Qual(typesImportPath, "SInt64TypeDescriptor"),
	// },

	clang.Type_UChar: {
		cType:          jen.Uint8(),
		goType:         jen.Id("UChar"),
		typeDescriptor: jen.Qual(typesImportPath, "UInt8TypeDescriptor"),
	},

	clang.Type_UInt: {
		cType:          jen.Uint32(),
		goType:         jen.Id("UInt"),
		typeDescriptor: jen.Qual(typesImportPath, "UInt32TypeDescriptor"),
	},

	clang.Type_UShort: {
		cType:          jen.Uint16(),
		goType:         jen.Id("UShort"),
		typeDescriptor: jen.Qual(typesImportPath, "UInt16TypeDescriptor"),
	},

	clang.Type_ULong: {
		cType:          jen.Uint64(),
		goType:         jen.Id("ULong"),
		typeDescriptor: jen.Qual(typesImportPath, "UInt64TypeDescriptor"),
	},
}

type typ struct {
	typ             clang.Type
	isArray         bool
	isCallback      bool
	isPointer       bool
	isScalar        bool
	isString        bool
	isStringPointer bool
	isStruct        bool
	isVoid          bool
	arraySize       int
	callback        *callback
	scalar          scalar
	struct_         *struct_
}

func newTyp(gen *gen, typ_ clang.Type) typ {
	typ := typ{
		typ: typ_,
	}

	typ.callback, typ.isCallback = gen.callbacks.find(typ.typ.Spelling())
	typ.scalar, typ.isScalar = scalars[clang.CursorKind(typ.typ.CanonicalType().Kind())]
	// typ.struct_, typ.isStruct = gen.structs.findByGoName(goName(typ.typ.Spelling()))
	typ.isVoid = typ.typ.CanonicalType().Kind() == clang.Type_Void

	// if strings.HasSuffix(typ.typ.Spelling(), "Callback") {
	// 	typ.isCallback = true
	// 	typ.callback = goName(typ.typ.Spelling())
	// }

	typ.isPointer = typ.typ.Kind() == clang.Type_Pointer
	if typ.isPointer {
		pointeeType := typ.typ.PointeeType()
		pointeeKind := pointeeType.CanonicalType().Kind()

		possibleStructName := pointeeType.Spelling()
		possibleStructName = strings.TrimPrefix(possibleStructName, "const ")

		typ.struct_, typ.isStruct = gen.structs.findByGoName(goName(possibleStructName))
		typ.scalar, typ.isScalar = scalars[clang.CursorKind(pointeeKind)]
		typ.isString = pointeeKind == clang.Type_Char_S || pointeeKind == clang.Type_UChar
		typ.isVoid = pointeeKind == clang.Type_Void

		if pointeeKind == clang.Type_Pointer {
			pointeePointeeKind := pointeeType.PointeeType().CanonicalType().Kind()
			typ.isStringPointer = pointeePointeeKind == clang.Type_Char_S || pointeePointeeKind == clang.Type_UChar
		}
	}

	typ.isArray = typ.typ.Kind() == clang.Type_ConstantArray
	if typ.isArray {
		typ.arraySize = int(typ.typ.ArraySize())
		elementType := typ.typ.ArrayElementType()
		typ.scalar, typ.isScalar = scalars[clang.CursorKind(elementType.Kind())]
	}

	return typ
}

func (typ typ) cDecl() jen.Code {
	if typ.isScalar {
		if typ.isPointer {
			return jen.Op("*").Add(typ.scalar.cType)
		}
		return typ.scalar.cType
	} else if typ.isString {
		return jen.Op("*").Byte()
	} else if typ.isStringPointer {
		return jen.Op("**").Byte()
	}

	return typ.goDecl()
}

func (typ typ) goDecl() jen.Code {
	if typ.isVoid {
		if typ.isPointer {
			return jen.Qual("unsafe", "Pointer")
		}
		return jen.Null()
	}
	if typ.isCallback {
		return jen.Id(typ.callback.name)
	}
	if typ.isScalar {
		if typ.isPointer {
			return jen.Op("*").Add(typ.scalar.goType)
		}
		return typ.scalar.goType
	}
	if typ.isStruct {
		if typ.isPointer {
			return jen.Op("*").Id(typ.struct_.name)
		}
		return jen.Id(typ.struct_.name)
	}
	if typ.isString {
		return jen.String()
	}

	// fmt.Printf("%#v\n", typ)
	panic(fmt.Sprintf("unhandled type : %s", typ.typ.Spelling()))
}

func (typ typ) typeDescriptor() jen.Code {
	if typ.isCallback {
		return pointerTypeDescriptor
	}
	if typ.isPointer {
		return pointerTypeDescriptor
	}
	if typ.isVoid {
		return voidTypeDescriptor
	}
	if typ.isScalar {
		return typ.scalar.typeDescriptor
	}
	if typ.isString {
		return pointerTypeDescriptor
	}

	panic("unhandled type")
}
