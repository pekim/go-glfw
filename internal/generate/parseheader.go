package generate

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"

	"github.com/go-clang/clang-v15/clang"
)

var clangResourceDir = sync.OnceValue(func() string {
	out, err := exec.Command("clang", "-print-resource-dir").Output()
	fatalOnError(err)

	resDir := strings.TrimSpace(string(out))
	parts := strings.Split(resDir, "\n")
	resDir = parts[0]

	if resDir == "" {
		fatal("no output when getting clang resource dir")
	}
	if !strings.HasPrefix(resDir, "/") {
		fatalf("expected clang resource dir to start with '/', but it %s", resDir)
	}

	return resDir
})

func (gen *gen) parseHeaderFile() {
	resourcesDir := clangResourceDir()
	parseArgs := []string{
		"-I", path.Join(resourcesDir, "include"),
		"-x", "c-header",
	}

	headerData, err := os.ReadFile(gen.headerFilename)
	fatalOnError(err)
	gen.contents = string(headerData)
	gen.lines = strings.Split(gen.contents, "\n")

	index := clang.NewIndex(0, 1)
	errCode := index.ParseTranslationUnit2(
		gen.headerFilename,
		parseArgs,
		nil,
		uint32(clang.TranslationUnit_SkipFunctionBodies|clang.TranslationUnit_DetailedPreprocessingRecord),
		&gen.tu,
	)
	if errCode != clang.Error_Success {
		fatal(errCode)
	}
}
