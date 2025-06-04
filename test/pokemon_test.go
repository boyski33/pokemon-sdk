//go:build integration

package test

import (
	_ "embed"
	"encoding/json"
	pokemon "github.com/boyski33/poke-sdk/v2"
	"github.com/boyski33/poke-sdk/v2/model"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed testdata/gen-7.json
var genSevenStub []byte

//go:embed testdata/pikachu-stub.json
var pikachuStub []byte

func TestResolver_Pokemon(t *testing.T) {
	t.Run("given pokemon exists with cache enabled when getting by name then only call server once", func(t *testing.T) {
		var mockServerInvocations int
		// GIVEN a mock server returning a stub (in case we want to run tests without internet access)
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mockServerInvocations++

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write(pikachuStub)
			require.NoError(t, err)
		}))

		const existingName = "pikachu"

		// GIVEN config with enabled cache
		config := pokemon.Config{
			BaseURL:      mockServer.URL,
			CacheEnabled: true,
		}

		// GIVEN resolver with config
		resolver := pokemon.NewResolver().WithConfig(config)

		// WHEN calling for the first time
		pikachuFromServer, err := resolver.Pokemon(existingName).Get()

		// THEN success
		require.NoError(t, err)
		require.NotNil(t, pikachuFromServer)

		// WHEN calling again
		pikachuFromCache, err := resolver.Pokemon(existingName).Get()

		// THEN success
		require.NoError(t, err)
		require.Equal(t, pikachuFromServer, pikachuFromCache)

		// THEN server only called once
		require.Equal(t, 1, mockServerInvocations)

		// THEN response equal to stub
		expected := &model.Pokemon{}
		require.NoError(t, json.Unmarshal(pikachuStub, expected))
		require.Equal(t, expected, pikachuFromServer)
	})

	t.Run("given pokemon exists with cache disabled when getting by name then call server twice", func(t *testing.T) {
		var mockServerInvocations int
		// GIVEN a mock server returning a stub (in case we want to run tests without internet access)
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mockServerInvocations++

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write(pikachuStub)
			require.NoError(t, err)
		}))

		const existingName = "pikachu"

		// GIVEN config with disabled cache
		config := pokemon.Config{
			BaseURL:      mockServer.URL,
			CacheEnabled: false,
		}

		// GIVEN resolver with config
		resolver := pokemon.NewResolver().WithConfig(config)

		// WHEN calling for the first time
		pikachuFromServer1, err := resolver.Pokemon(existingName).Get()

		// THEN success
		require.NoError(t, err)
		require.NotNil(t, pikachuFromServer1)

		// WHEN calling again
		pikachuFromServer2, err := resolver.Pokemon(existingName).Get()

		// THEN success
		require.NoError(t, err)
		require.Equal(t, pikachuFromServer1, pikachuFromServer2)

		// THEN server called twice
		require.Equal(t, 2, mockServerInvocations)
	})
}

func TestResolver_Generation(t *testing.T) {
	t.Run("given generation exists with cache enabled when getting by name then only call server once", func(t *testing.T) {
		var mockServerInvocations int
		// GIVEN a mock server returning a stub (in case we want to run tests without internet access)
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mockServerInvocations++

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write(genSevenStub)
			require.NoError(t, err)
		}))

		const existingGeneration = "7"

		// GIVEN config with enabled cache
		config := pokemon.Config{
			BaseURL:      mockServer.URL,
			CacheEnabled: true,
		}

		// GIVEN resolver with config
		resolver := pokemon.NewResolver().WithConfig(config)

		// WHEN calling for the first time
		generationFromServer, err := resolver.Generation(existingGeneration).Get()

		// THEN success
		require.NoError(t, err)
		require.NotNil(t, generationFromServer)

		// WHEN calling again
		generationFromCache, err := resolver.Generation(existingGeneration).Get()

		// THEN success
		require.NoError(t, err)
		require.Equal(t, generationFromServer, generationFromCache)

		// THEN server only called once
		require.Equal(t, 1, mockServerInvocations)

		// THEN response equal to stub
		expected := &model.Generation{}
		require.NoError(t, json.Unmarshal(genSevenStub, expected))
		require.Equal(t, expected, generationFromServer)
	})

	t.Run("given generation exists with cache disabled when getting by name then call server twice", func(t *testing.T) {
		var mockServerInvocations int
		// GIVEN a mock server returning a stub (in case we want to run tests without internet access)
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mockServerInvocations++

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write(genSevenStub)
			require.NoError(t, err)
		}))

		const existingGeneration = "7"

		// GIVEN config with disabled cache
		config := pokemon.Config{
			BaseURL:      mockServer.URL,
			CacheEnabled: false,
		}

		// GIVEN resolver with config
		resolver := pokemon.NewResolver().WithConfig(config)

		// WHEN calling for the first time
		genFromServer1, err := resolver.Generation(existingGeneration).Get()

		// THEN success
		require.NoError(t, err)
		require.NotNil(t, genFromServer1)

		// WHEN calling again
		genFromServer2, err := resolver.Generation(existingGeneration).Get()

		// THEN success
		require.NoError(t, err)
		require.Equal(t, genFromServer1, genFromServer2)

		// THEN server called twice
		require.Equal(t, 2, mockServerInvocations)
	})
}
