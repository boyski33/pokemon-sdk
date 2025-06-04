//go:build live

package test

import (
	"context"
	_ "embed"
	"encoding/json"
	"github.com/boyski33/pokemon-sdk/v2"
	"github.com/boyski33/pokemon-sdk/v2/model"
	"github.com/stretchr/testify/require"
	"slices"
	"testing"
)

//go:embed testdata/pikachu-stub.json
var actualPikachu []byte

//go:embed testdata/gen-7.json
var actualGen7 []byte

// This runs live tests against the public Pokémon API
func Test_Live_GetPokemon(t *testing.T) {
	t.Run("given pokemon exists when get pokemon then success", func(t *testing.T) {
		resolver := pokemon.NewResolver().WithConfig(pokemon.Config{
			CacheEnabled: true,
		})
		pikachu, err := resolver.Pokemon("pikachu").Get()
		require.NoError(t, err)
		expected := &model.Pokemon{}
		require.NoError(t, json.Unmarshal(actualPikachu, expected))
		require.Equal(t, expected, pikachu)
	})

	t.Run("given pokemon does not exist when get pokemon then return 404", func(t *testing.T) {
		resolver := pokemon.NewResolver()
		pikachu, err := resolver.Pokemon("some random pokemon name").Get()
		require.Nil(t, pikachu)
		require.ErrorIs(t, err, model.ErrNotFound)
	})
}

func Test_Live_GetPokemonList(t *testing.T) {
	t.Run("given 10 pokemon names when calling get list and next then return correct results", func(t *testing.T) {
		resolver := pokemon.NewResolver().WithConfig(pokemon.Config{
			CacheEnabled: true,
		})

		// GIVEN the first 10 Pokemon
		first10, err := resolver.PokemonList(1, 10).Get()
		require.NoError(t, err)
		require.Len(t, first10, 10)

		// GIVEN new list with page size 5
		list := resolver.PokemonList(1, 5)

		// WHEN getting the first page of 5 Pokemon
		first5, err := list.Next(context.Background())
		require.NoError(t, err)

		// WHEN getting the second page of 5 Pokemon
		second5, err := list.Next(context.Background())
		require.NoError(t, err)
		require.Len(t, second5, 5)

		// THEN firstFive + secondFive = first10
		require.Equal(t, first10, slices.Concat(first5, second5))
	})
}

func Test_Live_GetGeneration(t *testing.T) {
	t.Run("given generation exists when get generation then success", func(t *testing.T) {
		resolver := pokemon.NewResolver().WithConfig(pokemon.Config{
			CacheEnabled: true,
		})
		seven, err := resolver.Generation("7").Get()
		require.NoError(t, err)
		expected := &model.Generation{}
		require.NoError(t, json.Unmarshal(actualGen7, expected))
		require.Equal(t, expected, seven)
	})

	t.Run("given generation does not exist when get generation then return 404", func(t *testing.T) {
		resolver := pokemon.NewResolver().WithConfig(pokemon.Config{
			CacheEnabled: true,
		})
		seven, err := resolver.Generation("777").Get()
		require.Nil(t, seven)
		require.ErrorIs(t, err, model.ErrNotFound)
	})
}

func Test_Live_GetGenerationList(t *testing.T) {
	t.Run("given valid page when get generation list then success", func(t *testing.T) {
		resolver := pokemon.NewResolver().WithConfig(pokemon.Config{
			CacheEnabled: true,
		})

		page, err := resolver.GenerationList(1, 5).Get()
		require.NoError(t, err)
		require.Len(t, page, 5)
		require.Equal(t, "generation-i", page[0])
		require.Equal(t, "generation-ii", page[1])
	})
}
