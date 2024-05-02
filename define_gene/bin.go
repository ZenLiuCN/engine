package main

import (
	"github.com/ZenLiuCN/engine/define_gene/spec"
	"github.com/ZenLiuCN/fn"
	"github.com/urfave/cli/v2"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Usage:   "show version",
		Aliases: []string{"v"},
	}
	app := cli.NewApp()
	app.UseShortOptionHandling = true
	app.Name = "Define Generator"
	app.Version = "v0.0.1"
	app.Usage = "Generate engine define from go source"
	app.Flags = []cli.Flag{
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
		&cli.BoolFlag{
			Name:    "override",
			Usage:   "override exist file",
			Aliases: []string{"r"},
		},
		&cli.StringSliceFlag{
			Name:    "tags",
			Usage:   "tags to apply",
			Aliases: []string{"g"},
		},
		&cli.StringFlag{
			Name:    "out",
			Usage:   "output path or current path",
			Aliases: []string{"o"},
		},
		&cli.StringFlag{
			Name:    "go",
			Usage:   "go root to walk all embed types",
			Aliases: []string{"x"},
		},
	}
	app.Suggest = true
	app.EnableBashCompletion = true
	app.Action = action
	fn.Panic(app.Run(os.Args))
}

func action(c *cli.Context) error {
	var file []string
	if c.Args().Len() == 0 {
		file = append(file, ".")
	} else {
		file = append(file, c.Args().Slice()...)
	}
	tags := c.StringSlice("tags")
	if c.String("go") == "" {
		var dir string
		if len(file) == 1 && spec.IsDir(file[0]) {
			dir = file[0]
		} else if len(tags) != 0 {
			log.Fatal("--tags can only applies with directory")
		} else {
			dir = filepath.Dir(file[0])
		}
		g := new(Generator)
		{
			g.dir = dir
			g.over = c.Bool("r")
			g.tags = tags
			g.files = file
			g.types = c.StringSlice("type")
			g.out = c.String("out")
			if c.Bool("debug") {
				g.log = log.Printf
			}
		}
		g.generate()
	} else {
		root := filepath.Join(c.String("go"), "src")
		return filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() || root == path {
				return nil
			}
			if strings.HasSuffix(path, "builtin") ||
				strings.HasSuffix(path, "go") ||
				strings.HasSuffix(path, "internal") ||
				strings.HasSuffix(path, "debug") ||
				strings.HasSuffix(path, "embed") ||
				strings.HasSuffix(path, "runtime") ||
				strings.HasSuffix(path, "vendor") ||
				strings.HasSuffix(path, "heap") ||
				strings.HasSuffix(path, "unsafe") ||
				strings.HasSuffix(path, "testing") ||
				strings.HasSuffix(path, "testdata") ||
				strings.HasSuffix(path, "reflect") ||
				strings.HasSuffix(path, "plugin") ||
				strings.HasSuffix(path, "cmd") ||
				strings.HasSuffix(path, "sync") ||
				strings.HasSuffix(path, "syscall") ||
				strings.HasSuffix(path, "pprof") ||
				strings.HasSuffix(path, "flag") ||
				strings.HasSuffix(path, "text") ||
				strings.HasSuffix(path, "log") ||
				strings.HasSuffix(path, "cmp") ||
				strings.HasSuffix(path, "cgi") ||
				strings.HasSuffix(path, "fcgi") ||
				strings.HasSuffix(path, "database") ||
				strings.HasSuffix(path, "crypto") {
				log.Printf("ignore: %s", path)
				return filepath.SkipDir
			}
			log.Printf("generate: %s", path)
			g := new(Generator)
			g.tags = tags
			g.files = []string{path}
			g.types = c.StringSlice("type")
			g.out = c.String("out")
			g.over = c.Bool("r")
			if c.Bool("debug") {
				g.log = log.Printf
			}
			g.generate()
			return nil
		})
	}
	return nil
}
