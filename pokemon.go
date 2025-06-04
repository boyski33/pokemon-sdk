package pokemon

import (
	"context"
	"github.com/boyski33/poke-sdk/v2/model"
	"io"
	"net/http"
)

// Pokemon is a helper type with a reference to the Resolver making the requests.
type Pokemon struct {
	resolver *Resolver
	id       string
}

// Get returns the model.Pokemon by the ID set on the Pokemon type.
func (p *Pokemon) Get() (*model.Pokemon, error) {
	return p.GetWithContext(context.Background())
}

// GetWithContext returns the model.Pokemon by the ID set on the Pokemon type. You can pass a context.Context if you
// want more granular control over the lifecycle of the request i.e., setting timeouts.
func (p *Pokemon) GetWithContext(ctx context.Context) (*model.Pokemon, error) {
	data, err := p.resolver.getPokemon(ctx, p.id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type List struct {
	page     int
	pageSize int
	limit    int
	done     bool
}

// PokemonList is a helper type with a reference to the Resolver making the requests. It is used to make paginated
// requests to the Pokemon API.
type PokemonList struct {
	resolver *Resolver
	List
}

// Get returns a specific page based on the PokemonList. You can use Next if you need more results.
func (l *PokemonList) Get() ([]string, error) {
	return l.GetWithContext(context.Background())
}

// GetWithContext returns a specific page based on the PokemonList. You can use Next if you need more results. You can pass a
// context.Context if you want more granular control over the lifecycle of the request i.e., setting timeouts.
func (l *PokemonList) GetWithContext(ctx context.Context) ([]string, error) {
	names, _, err := l.resolver.getPokemonNames(ctx, l.pageSize, (l.page-1)*l.pageSize)
	if err != nil {
		return nil, err
	}

	return names, nil
}

// The Next method allows for cursor pagination. Defined by the pageSize field in the PokemonList object, Next fetches the next
// N Pokemon names from the public API. The names can later be passed in the Pokemon function to retrieve data about a
// specific Pokemon.
//
// The function returns an io.EOF error when the last page is hit.
func (l *PokemonList) Next(ctx context.Context) ([]string, error) {
	if l.done {
		return nil, io.EOF
	}

	names, hasMore, err := l.resolver.getPokemonNames(ctx, l.pageSize, (l.page-1)*l.pageSize)
	if err != nil {
		return nil, err
	}

	if !hasMore {
		l.done = true
	} else {
		l.page++
	}

	return names, nil
}

// Generation is a helper type with a reference to the Resolver making the requests
type Generation struct {
	resolver *Resolver
	id       string
}

// Get returns the model.Generation by the ID set on the Generation type.
func (g *Generation) Get() (*model.Generation, error) {
	return g.GetWithContext(context.Background())
}

// GetWithContext returns the model.Generation by the ID set on the Generation type. You can pass a context.Context
// if you want more granular control over the lifecycle of the request i.e., setting timeouts.
func (g *Generation) GetWithContext(ctx context.Context) (*model.Generation, error) {
	data, err := g.resolver.getGeneration(ctx, g.id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type GenerationList struct {
	resolver *Resolver
	List
}

// Get returns a specific page based on the GenerationList. You can use Next if you need more results.
func (l *GenerationList) Get() ([]string, error) {
	return l.GetWithContext(context.Background())
}

// GetWithContext returns a specific page based on the GenerationList. You can use Next if you need more results.
// You can pass a context.Context if you want more granular control over the lifecycle of the request
// i.e., setting timeouts.
func (l *GenerationList) GetWithContext(ctx context.Context) ([]string, error) {
	names, _, err := l.resolver.getGenerationNames(ctx, l.pageSize, (l.page-1)*l.pageSize)
	if err != nil {
		return nil, err
	}

	return names, nil
}

// The Next method allows for cursor pagination. Defined by the pageSize field in the GenerationList object,
// Next fetches the next N Generation names from the public API. The names can later be passed in the Generation
// function to retrieve data about a specific Generation.
//
// The function returns an io.EOF error when the last page is hit.
func (l *GenerationList) Next(ctx context.Context) ([]string, error) {
	if l.done {
		return nil, io.EOF
	}

	names, hasMore, err := l.resolver.getGenerationNames(ctx, l.pageSize, (l.page-1)*l.pageSize)
	if err != nil {
		return nil, err
	}

	if !hasMore {
		l.done = true
	} else {
		l.page++
	}

	return names, nil
}

// Resolver is the main type you will use when interacting with the SDK. It creates helper objects, such as Pokemon and
// Generation, and it contains the HTTP client used for fetching data from the remote Pokemon API.
type Resolver struct {
	client *client
}

// NewResolver returns a Resolver with a default client and the cache disabled. If you want to change the configuration,
// you can use the Resolver.WithConfig function.
func NewResolver() *Resolver {
	return &Resolver{
		client: newClient(defaultBaseURL, &http.Client{Timeout: defaultClientTimeout}, nil),
	}
}

// WithConfig accepts a Config object and sets the specified configuration to the Resolver.
func (r *Resolver) WithConfig(config Config) *Resolver {
	if config.BaseURL != "" {
		r.client.baseURL = config.BaseURL
	}

	if config.ClientTimeout != 0 {
		r.client.httpClient.Timeout = config.ClientTimeout
	}

	if config.CacheEnabled {
		r.client.cache = NewCache(config.CacheTTL)
	}

	return r
}

// Pokemon returns a new Pokemon object with the specified identifier (ID or name) and a reference to the Resolver.
func (r *Resolver) Pokemon(id string) *Pokemon {
	return &Pokemon{
		resolver: r,
		id:       id,
	}
}

// PokemonList returns a new PokemonList object with the specified page size and a reference to the Resolver.
func (r *Resolver) PokemonList(page, pageSize int) *PokemonList {
	return &PokemonList{
		resolver: r,
		List: List{
			page:     page,
			pageSize: pageSize,
		},
	}
}

// Generation returns a new Generation object with the specified
// identifier (ID or name) and a reference to the Resolver.
func (r *Resolver) Generation(id string) *Generation {
	return &Generation{
		resolver: r,
		id:       id,
	}
}

// GenerationList returns a new GenerationList object with the specified page size and a reference to the Resolver.
func (r *Resolver) GenerationList(page, pageSize int) *GenerationList {
	return &GenerationList{
		resolver: r,
		List: List{
			page:     page,
			pageSize: pageSize,
		},
	}
}

func (r *Resolver) getPokemon(ctx context.Context, id string) (*model.Pokemon, error) {
	return r.client.GetPokemonByIDOrName(ctx, id)
}

func (r *Resolver) getPokemonNames(ctx context.Context, limit, offset int) ([]string, bool, error) {
	return r.client.GetPokemonNamesList(ctx, limit, offset)
}

func (r *Resolver) getGeneration(ctx context.Context, id string) (*model.Generation, error) {
	return r.client.GetGenerationByIDOrName(ctx, id)
}

func (r *Resolver) getGenerationNames(ctx context.Context, limit, offset int) ([]string, bool, error) {
	return r.client.GetGenerationNamesList(ctx, limit, offset)
}
