package generate

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-clang/clang-v15/clang"
	"github.com/pekim/go-glfw/internal/tagfile"
)

type gen struct {
	headerFilename  string
	tu              clang.TranslationUnit
	contents        string
	lines           []string
	entityURLs      map[string]string
	pageSectionURLs map[string]tagfile.URL
	groupURLs       map[string]tagfile.URL
	macroConstants  macroConstants
	callbacks       callbacks
	structs         structs
	functions       functions
}

func Generate() {
	gen := gen{
		headerFilename:  "internal/generate/glfw3.h",
		entityURLs:      tagfile.EntityURLs(),
		pageSectionURLs: tagfile.PageSectionURLs(),
		groupURLs:       tagfile.GroupURLs(),
	}

	timeFunction("TOTAL", func() {
		timeFunction("parse header", gen.parseHeaderFile)
		timeFunction("find entities", gen.findEntities)
		timeFunction("generate", gen.generateFiles)
	})
	fmt.Println()

	gen.printStatistics()
	fmt.Println()
}

func (gen *gen) findEntities() {
	gen.tu.TranslationUnitCursor().Visit(func(cursor, _parent clang.Cursor) clang.ChildVisitResult {
		// ignore cursors that are not from the header file
		file, _, _, _ := cursor.Location().FileLocation()
		if file.Name() != gen.headerFilename {
			return clang.ChildVisit_Continue
		}

		switch cursor.Kind() {

		case clang.Cursor_MacroDefinition:
			gen.newMacroConstant(cursor)

		case clang.Cursor_FunctionDecl:
			gen.functions = append(gen.functions, newFunction(gen, cursor))

		case clang.Cursor_StructDecl:
			cName := cursor.Spelling()
			gen.structs = append(gen.structs, struct_{
				gen:    gen,
				cursor: cursor,
				cName:  cName,
				name:   goName(cName),
			})

		case clang.Cursor_TypedefDecl:
			if cursor.TypedefDeclUnderlyingType().Kind() == clang.Type_Pointer {
				gen.callbacks = append(gen.callbacks, newCallback(gen, cursor))
			}

			// default:
			// 	fmt.Println(cursor.Kind().Spelling(), cursor.Spelling())
		}

		return clang.ChildVisit_Continue
	})
}

func (gen gen) generateFiles() {
	gen.macroConstants.createComments(gen)
	gen.callbacks.createComments(gen)
	gen.functions.createComments(gen)
	gen.structs.createComments(gen)

	gen.macroConstants.generate()
	gen.callbacks.generate()
	gen.functions.generate()
	gen.structs.generate()

	gen.generateStructTest()
}

func (gen gen) resolveEntityReference(ref string) (string, bool) {
	ref = strings.TrimSuffix(ref, "_attrib")
	ref = strings.TrimSuffix(ref, "_hint")

	if callback, ok := gen.callbacks.find(ref); ok {
		return callback.name, true
	} else if constant, ok := gen.macroConstants.find(ref); ok {
		return constant.name, true
	} else if function, ok := gen.functions.find(ref); ok {
		return function.name, true
	} else if function, ok := gen.functions.find(ref); ok {
		return function.name, true
	} else if struct_, ok := gen.structs.find(ref); ok {
		return struct_.name, true
	}

	return "", false
}

func (gen *gen) printStatistics() {
	functionCount := len(gen.functions)
	supportedFunctionCount := 0
	for _, fn := range gen.functions {
		if supported, _ := fn.supported(); supported {
			supportedFunctionCount++
		}
	}
	fmt.Printf("functions supported  : %d/%d  %1.1f%%\n",
		supportedFunctionCount,
		functionCount,
		(float64(supportedFunctionCount)/float64(functionCount))*100,
	)
}

func timeFunction(title string, fn func()) {
	start := time.Now()
	fn()
	fmt.Printf("%-20s : %.0fms\n",
		title,
		time.Since(start).Seconds()*1_000,
	)
}
