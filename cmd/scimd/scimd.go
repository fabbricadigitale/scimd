package main

import (
	"fmt"

	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/schemas/core"
)

func main() {
	// Initialize configurations
	spc := config.Get()

	fmt.Println(spc)
	fmt.Println(core.GetSchemaRepository().List())
	fmt.Println(core.GetResourceTypeRepository().List())
	fmt.Println(config.Values)

	// Start server
	// server.Get(spc).Run(":8787")
}
