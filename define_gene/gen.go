package main

//https://cs.opensource.google/go/x/tools/+/master:cmd/stringer/stringer.go;drc=daf94608b5e2caf763ba634b84e7a5ba7970e155;l=382
import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go/ast"
	"log"
	"os"
	"path/filepath"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Usage:   "show version",
		Aliases: []string{"v"},
	}
	err := (&cli.App{
		UseShortOptionHandling: true,
		Name:                   "Define Generator",
		Version:                "v0.0.1",
		Usage:                  "Generate engine define from go source",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "type",
				Usage:   "type names of current file or directory",
				Aliases: []string{"t"},
			},
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "debug generator",
				Aliases: []string{"d"},
			},
			&cli.StringSliceFlag{
				Name:    "tags",
				Usage:   "tags to apply",
				Aliases: []string{"g"},
			},
		},
		Suggest:              true,
		EnableBashCompletion: true,
		Action: func(c *cli.Context) error {
			var file []string
			if c.Args().Len() == 0 {
				file = append(file, ".")
			} else {
				file = append(file, c.Args().Slice()...)
			}
			tags := c.StringSlice("tags")
			var dir string
			if len(file) == 1 && isDir(file[0]) {
				dir = file[0]
			} else if len(tags) != 0 {
				log.Fatal("--tags can only applies with directory")
			} else {
				dir = filepath.Dir(file[0])
			}
			g := new(Generator)
			{
				g.dir = dir
				g.tags = tags
				g.files = file
				g.types = c.StringSlice("type")
				g.out = c.String("out")
				if c.Bool("debug") {
					g.Logf = log.Printf
				}
			}
			g.generate()
			return nil
		},
	}).Run(os.Args)
	if err != nil {
		panic(err)
	}
}
func isDir(name string) bool {
	if info, err := os.Stat(name); err != nil {
		log.Fatal(err)
	} else {
		return info.IsDir()
	}
	return false
}

type Generator struct {
	dir    string
	tags   []string
	files  []string
	types  []string
	access bool
	SpecContext
	out      string
	typeSpec []*SpecType
	*SpecWriter
}

func (g *Generator) generate() {
	g.Parse(g.tags, g.files)
	if len(g.SpecContext.SpecPackages) != 1 {
		panic("only one package for each generation")
	}
	g.Process()
}

func (g *Generator) Process() {
	for _, p := range g.SpecPackages {
		for _, file := range p.SpecFiles {
			if file.AstFile != nil {
				ast.Inspect(file.AstFile, FindTypeSpec(file, true, func(s *SpecType) {
					g.typeSpec = append(g.typeSpec, s)
				}))
			}
		}
	}
	for _, specType := range g.typeSpec {
		fmt.Printf("%#+v\n", specType.Resolve())
	}
}
