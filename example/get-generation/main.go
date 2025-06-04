package main

import (
	"fmt"
	pokemon "github.com/boyski33/pokemon-sdk/v2"
	"log"
	"time"
)

func main() {
	resolver := pokemon.NewResolver().WithConfig(pokemon.Config{
		ClientTimeout: 10 * time.Second,
		CacheEnabled:  true,
		CacheTTL:      1 * time.Minute,
	})

	list := resolver.GenerationList(1, 5)
	firstPage, err := list.Get()
	if err != nil {
		log.Fatal(err)
	}

	gen, err := resolver.Generation(firstPage[0]).Get()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generation name: ", gen.Name)
}
