# Pokemon SDK

A Go library for interacting with the public Pokemon API: https://pokeapi.co/docs/v2

## API

The main way to interact with the SDK is through a `Resolver`. You can create one by:

```go
package mypackage

import pokemon "github.com/boyski33/pokemon-sdk/v2"

func main() {
	resolver := pokemon.NewResolver()
}
```

You can specify optional configuration, or the resolver will be configured with sensible defaults. For example:

```go
resolver := pokemon.NewResolver().WithConfig(pokemon.Config{
ClientTimeout: 10 * time.Second,
CacheEnabled:  true,
CacheTTL:      1 * time.Minute,
})
```

### Pokemon

To get a single Pokemon by identifier (either ID or name):

```go
pikachu := resolver.Pokemon("pikachu").Get()
```

To get a specific page of Pokemon names by providing a page and page size:

```go
page := resolver.PokemonList(1, 3).Get()
```

`page` is a string slice of names e.g., `["bulbasaur", "ivysaur", "venusaur"]`. These names can be used as identifiers
when getting single Pokemon. See `example` folder for more.

Alternatively, you can use cursor-based pagination:

```go
list := resolver.PokemonList(1, 5)
page1 := list.Next(context.Background())
page2 := list.Next(context.Background())
```

If you want to iterate through all Pokemon names:

```go
ctx := context.Background()
for names, err := list.Next(ctx); err == nil; names, err = list.Next(ctx) {
...
}
```

### Generation

A generation is a grouping of the Pokémon games that separates them based on the Pokémon they include. In each
generation, a new set of Pokémon, Moves, Abilities and Types that did not exist in the previous generation are released.

The API is identical to the Pokemon one.

To get a single generation by identifier (either ID or name):

```go
pikachu := resolver.Generation("1").Get()
```

To get a specific page of Generations names by providing a page and page size:

```go
page := resolver.GenerationList(1, 3).Get()
```

`page` is a string slice of names e.g., `["generation-i", "generation-ii]`. These names can be used as identifiers
when getting single generations. See `example` folder for more.

Alternatively, you can use cursor-based pagination:

```go
list := resolver.GenerationList(1, 5)
page1 := list.Next(context.Background())
page2 := list.Next(context.Background())
```

It's good to note there are currently only nine generations, so pagination is probably unnecessary, but it is added for
consistency with the Pokemon API.

## Project structure

The `model` package contains the important resource types e.g., `Pokemon` and `Generation`.

The root `pokemon` package contains all the business logic, as well as the HTTP client and in-memory cache.
Following a flat structure is idiomatic in Go, since we don't want to export all types to the public (even if we
decide to use `internal` packages).

The `test` folder contains integration tests, which use a mock server, as well as live tests, which make actual
requests to the public API. The API doesn't require authentication and isn't rate-limited, so you shouldn't run into
issues running the live tests barring connectivity issues. There is a `testdata` folder which contains JSON stubs with
real data from the Pokemon API.

The `example` folder contains modules with main files, which can be modified and run. The existing code demonstrates
how to use the SDK.

## Technical decisions

### Pagination

I have decided to go with cursor-based pagination, as it's user-friendly. The API allows for fetching any
specific page anyway, so the client can decide to use that instead.

### Caching

The library can be configured to cache HTTP response payloads in memory. This can be specified in the `Config` object
when creating a resolver. Additionally, a TTL can be specified. Otherwise, the library caches responses without
expiration.

The reasons for implementing a cache are to reduce latency and traffic since responses can be fairly large at 7–10 KB.

I've used a popular library for this purpose: https://github.com/patrickmn/go-cache

### Timeouts

The HTTP client timeout can be specified by the `ClientTimeout` field in the `Config`.

Alternatively, API exposes both `Get()` and `GetWithContext()` functions. This allows the client to set specific
timeouts per request if needed by providing a `context.Context`.

### Versioning

Since the Pokemon API only has v2 exposed, I've set the module name to `github.com/boyski33/pokemon-sdk/v2`. If a v3 comes
out, the library can migrate to v3 and change the module name to `github.com/boyski33/pokemon-sdk/v3`. This will happen
alongside a new git tag following that convention. This approach is taken from https://github.com/redis/go-redis.

## How to run

You can test the SDK in the `example` folder by modifying any of the examples. You can run `go run main.go` from any of
the example packages.

## How to run tests

To run integration tests (with mocks): `make test-integration`
To run live tests (against the public API): `make test-live`
