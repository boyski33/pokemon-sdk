package main

import (
	"fmt"
	"github.com/boyski33/pokemon-sdk/v2"
	"log"
	"time"
)

func main() {
	resolver := pokemon.NewResolver().WithConfig(pokemon.Config{
		ClientTimeout: 10 * time.Second,
		CacheEnabled:  true,
		CacheTTL:      1 * time.Minute,
	})

	start := time.Now()

	list := resolver.PokemonList(1, 5)
	firstPage, err := list.Get()
	if err != nil {
		log.Fatal(err)
	}

	first, err := resolver.Pokemon(firstPage[0]).Get()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("time elapsed: ", time.Since(start))

	fmt.Println("Name:", first.Name)
	fmt.Println("Weight:", first.Weight)
}
