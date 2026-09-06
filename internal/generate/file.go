package generate

import (
	"fmt"
	"os/exec"

	"github.com/dave/jennifer/jen"
)

type file struct {
	*jen.File
	filename string
}

func newFile(filename string, pkg string) file {
	f := file{
		File:     jen.NewFile(pkg),
		filename: filename,
	}

	f.HeaderComment("This is a generated file. DO NOT EDIT.")
	f.Line()

	return f
}

func (f file) save() {
	f.NoFormat = true
	err := f.Save(f.filename)
	if err != nil {
		fatalOnError(err)
	}

	cmd := exec.Command("goimports", "-w", f.filename)
	fmtOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("failed to format source, writing unformatted source to %s\n", f.filename)
		f.NoFormat = true
		err := f.Save(f.filename)
		fatalOnError(err)
	}
	if len(fmtOutput) > 0 {
		fmt.Println(string(fmtOutput))
	}
}
