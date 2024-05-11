package main

import (
	"errors"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
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

		&cli.StringFlag{
			Name:    "cond",
			Usage:   "condition and suffix for module",
			Aliases: []string{"c"},
		},
		&cli.BoolFlag{
			Name:    "debug",
			Usage:   "debug generator",
			Aliases: []string{"d"},
		},

		&cli.StringSliceFlag{
			Name:    "env",
			Usage:   "envs to apply",
			Aliases: []string{"e"},
		},
		&cli.StringSliceFlag{
			Name:    "tags",
			Usage:   "tags to apply",
			Aliases: []string{"g"},
		},
		&cli.StringSliceFlag{
			Name:    "ignore",
			Usage:   "ignore packages (package path)",
			Aliases: []string{"i"},
		},
		&cli.StringFlag{
			Name:    "out",
			Usage:   "output path or current path",
			Aliases: []string{"o"},
		},
		&cli.BoolFlag{
			Name:    "print",
			Usage:   "print generate results only",
			Aliases: []string{"p"},
		},
		&cli.BoolFlag{
			Name:    "rewrite",
			Usage:   "rewrite exist file",
			Aliases: []string{"r"},
		},
		&cli.BoolFlag{
			Name:    "errors",
			Usage:   "generate with error as secondary result",
			Aliases: []string{"s"},
		},
		&cli.BoolFlag{
			Name:    "test",
			Usage:   "rewrite any exists test file",
			Aliases: []string{"t"},
		},

		&cli.BoolFlag{
			Name:   "trace",
			Usage:  "trace",
			Hidden: true,
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
	var dir string
	if len(file) == 1 && IsDir(file[0]) {
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
		g.out = c.String("out")
		g.over = c.Bool("rewrite")
		g.overTest = c.Bool("test")
		g.print = c.Bool("print")
		g.trace = c.Bool("trace")
		g.suffix = c.String("cond")
		g.env = c.StringSlice("env")
		g.errors = c.Bool("errors")
		g.ignores = c.StringSlice("ignore")
		if c.Bool("debug") {
			g.log = log.Printf
		}
	}
	return g.generate()
}

func IsDir(s string) bool {
	i, err := os.Stat(s)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}
	if err != nil {
		return false
	}
	return i.IsDir()
}

var (
	debug = false
)

func debugf(format string, args ...any) {
	if debug {
		_ = log.Output(2, "[DEBUG] "+fmt.Sprintf(format, args...))
	}
}
func tracef(n int, format string, args ...any) {
	if debug {
		_ = log.Output(n+2, "[TRACE] "+fmt.Sprintf(format, args...))
	}
}
