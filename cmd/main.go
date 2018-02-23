package main

import (
	"log"

	"github.com/fabbricadigitale/scimd/cmd"
	cobradoc "github.com/spf13/cobra/doc"
)

// Install and use to generate CLI markdown documentation
func main() {
	err := cobradoc.GenMarkdownTree(cmd.Get(), "/tmp")
	if err != nil {
		log.Fatal(err)
	}
}
