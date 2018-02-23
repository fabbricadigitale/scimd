package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fabbricadigitale/scimd/cmd"
	cobradoc "github.com/spf13/cobra/doc"
)

// Run to generate CLI markdown documentation
//
// Assumes docs/cli already exists in the root directory (first argument).
func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		log.Fatal("missing root directory")
	}

	dir, err := filepath.Abs(filepath.Join(argsWithoutProg[0], "docs/cli"))
	check(err)

	err = cobradoc.GenMarkdownTree(cmd.Get(), dir)
	check(err)

	log.Printf("CLI documentation successfully generated at \"%s\"\n", dir)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
