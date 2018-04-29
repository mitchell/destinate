package schema

import (
	"context"

	"googlemaps.github.io/maps"
)

// AllDestinations is used across the service to get all destinations based on a query, and returns results and pagination token.
func AllDestinations(ctx context.Context, query string) (*maps.PlacesSearchResponse, error) {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyCWhNQ3cPHw7qFiiRTucVd61W4ZSzeUyTI"))
	if err != nil {
		return nil, err
	}

	tsr := &maps.TextSearchRequest{
		Query: query,
	}

	resp, err := c.TextSearch(ctx, tsr)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
