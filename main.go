package main

import (
	"github.com/alecthomas/kong"
)

var CLI struct {
	Extract ExtractCommand `cmd help:"Extract data from Hugo frontmatter"`
}

func (ex *ExtractCommand) Run(ctx *kong.Context) error {
	return ex.HandleExtractCommand(ctx)
}

func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
