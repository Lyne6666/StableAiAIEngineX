// cmd/stableaiaienginex/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"stableaiaienginex/internal/stableaiaienginex"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := stableaiaienginex.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
