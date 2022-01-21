package main

import (
	"os"

  "gitlab.com/tokend/blob-serv/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
