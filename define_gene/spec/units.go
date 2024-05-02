package spec

import (
	"github.com/ZenLiuCN/fn"
	"golang.org/x/tools/go/packages"
	"log"
	"os"
)

var (
	IdentityLoadMode = fn.Identity[packages.LoadMode]
)

func IsDir(name string) bool {
	if info, err := os.Stat(name); err != nil {
		log.Fatal(err)
	} else {
		return info.IsDir()
	}
	return false
}

// ParseTypeInfo files to packages with tags
// flags: optional compile flags;
// files: required file targets;
// mode: required mode rewriter;
// log: optional debug output;
func ParseTypeInfo(flags, files []string, mode func(packages.LoadMode) packages.LoadMode, log func(format string, args ...any)) []*packages.Package {
	return fn.Panic1(packages.Load(&packages.Config{
		Mode:       mode(packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax),
		BuildFlags: flags,
		Tests:      false,
		Logf:       log,
	}, files...))
}
