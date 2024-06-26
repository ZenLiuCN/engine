// Code generated by define_gene; DO NOT EDIT.
package jade

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	"github.com/Joker/jade"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/http"
)

var (
	//go:embed github_com_joker_jade.d.ts
	GithubComJokerJadeDefine   []byte
	GithubComJokerJadeDeclared = map[string]any{
		"config":                  jade.Config,
		"parse":                   jade.Parse,
		"parseFile":               jade.ParseFile,
		"parseFileFromFileSystem": jade.ParseFileFromFileSystem,
		"parseWithFileSystem":     jade.ParseWithFileSystem,
		"ReadFunc":                jade.ReadFunc,
		"TabSize":                 jade.TabSize,

		"emptyReplaseTokens":    engine.Empty[jade.ReplaseTokens],
		"emptyRefReplaseTokens": engine.EmptyRefer[jade.ReplaseTokens],
		"refOfReplaseTokens":    engine.ReferOf[jade.ReplaseTokens],
		"unRefReplaseTokens":    engine.UnRefer[jade.ReplaseTokens]}
)

func init() {
	engine.RegisterModule(GithubComJokerJadeModule{})
}

type GithubComJokerJadeModule struct{}

func (S GithubComJokerJadeModule) Identity() string {
	return "github.com/Joker/jade"
}
func (S GithubComJokerJadeModule) TypeDefine() []byte {
	return GithubComJokerJadeDefine
}
func (S GithubComJokerJadeModule) Exports() map[string]any {
	return GithubComJokerJadeDeclared
}
