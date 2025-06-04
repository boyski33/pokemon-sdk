package pokemon

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/boyski33/pokemon-sdk/v2/model"
	"io"
	"net/http"
)

type client struct {
	baseURL    string
	httpClient *http.Client
	cache      *Cache
}

func newClient(baseURL string, cl *http.Client, c *Cache) *client {
	return &client{
		baseURL:    baseURL,
		httpClient: cl,
		cache:      c,
	}
}

// GetPokemonByIDOrName return a model.Pokemon pointer for the provided identifier or an error.
//
// A model.ErrNotFound is returned if the Pokemon does not exist.
func (c *client) GetPokemonByIDOrName(ctx context.Context, idOrName string) (*model.Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", c.baseURL, idOrName)

	var result model.Pokemon
	if err := c.loadFromCache(url, &result); err == nil {
		return &result, nil
	}

	body, err := c.fetchFromURL(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch generation: %w", err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.saveToCache(url, body)

	return &result, nil
}

// GetPokemonNamesList return a list of Pokemon names, a flag if there are more Pokemon to be fetched and an error.
func (c *client) GetPokemonNamesList(ctx context.Context, limit, offset int) (names []string, hasMore bool, err error) {
	type response struct {
		Count   int                   `json:"count"`
		Results []model.NamedResource `json:"results"`
	}

	url := fmt.Sprintf("%s/pokemon?limit=%d&offset=%d", c.baseURL, limit, offset)

	var resp response

	body, err := c.fetchFromURL(ctx, url)
	if err != nil {
		return nil, false, fmt.Errorf("failed to fetch generation: %w", err)
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, false, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	hasMore = resp.Count > offset+limit

	names = make([]string, len(resp.Results))
	for i, res := range resp.Results {
		names[i] = res.Name
	}

	return names, hasMore, nil
}

// GetGenerationByIDOrName return a model.Generation pointer for the provided identifier or an error.
//
// A model.ErrNotFound is returned if the Generation does not exist.
func (c *client) GetGenerationByIDOrName(ctx context.Context, idOrName string) (*model.Generation, error) {
	url := fmt.Sprintf("%s/generation/%s", c.baseURL, idOrName)

	var result model.Generation
	if err := c.loadFromCache(url, &result); err == nil {
		return &result, nil
	}

	body, err := c.fetchFromURL(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch generation: %w", err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.saveToCache(url, body)

	return &result, nil
}

// GetGenerationNamesList return a list of Generation names, a flag if there are more to be fetched and an error.
func (c *client) GetGenerationNamesList(ctx context.Context, limit, offset int) (names []string, hasMore bool, err error) {
	type response struct {
		Count   int                   `json:"count"`
		Results []model.NamedResource `json:"results"`
	}

	url := fmt.Sprintf("%s/generation?limit=%d&offset=%d", c.baseURL, limit, offset)

	var resp response

	body, err := c.fetchFromURL(ctx, url)
	if err != nil {
		return nil, false, fmt.Errorf("failed to fetch generation: %w", err)
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, false, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	hasMore = resp.Count > offset+limit

	names = make([]string, len(resp.Results))
	for i, res := range resp.Results {
		names[i] = res.Name
	}

	return names, hasMore, nil
}

func (c *client) fetchFromURL(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, model.ErrNotFound
		}
		return nil, fmt.Errorf("something went wrong")
	}

	return io.ReadAll(resp.Body)
}

func (c *client) loadFromCache(url string, v any) error {
	if c.cache == nil {
		return fmt.Errorf("cache disabled")
	}

	data := c.cache.GetResponseBodyForURL(url)
	if data == nil {
		return fmt.Errorf("cache miss")
	}

	return json.Unmarshal(data, v)
}

func (c *client) saveToCache(url string, data []byte) {
	if c.cache != nil {
		c.cache.CacheResponseForURL(url, data)
	}
}
