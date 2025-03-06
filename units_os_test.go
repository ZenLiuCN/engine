package engine

import (
	"github.com/ZenLiuCN/fn"
	"os"
	"testing"
)

func TestEnvExpand(t *testing.T) {
	println(EnvExpand("@/env/some"))
	println(EnvSetPath("SOME_ENV", "@/env/some"))
	println(EnvExpand("$SOME_ENV/env/some"))
}
func TestCreateSymlink(t *testing.T) {
	WriteTextFile("./md", "sample")
	fn.Panic(CreateSymlink("./md", "./sym"))
	if ReadTextFile("./sym") != "sample" {
		panic("symlink not created")
	}
	os.Remove("./sym")
	os.Remove("./md")
}
