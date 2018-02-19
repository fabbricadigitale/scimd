package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/defaults"
	"github.com/fabbricadigitale/scimd/schemas/core"
)

func main() {
	spew.Dump(defaults.ServiceProviderConfig)
	fmt.Println()
	spew.Dump(core.GetSchemaRepository().List())
	fmt.Println()
	spew.Dump(core.GetResourceTypeRepository().List())
	fmt.Println()
	spew.Dump(config.Values)

	// Start server
	// server.Get(spc).Run(":8787")
}
