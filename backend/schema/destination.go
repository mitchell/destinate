package schema

import (
	"context"

	"googlemaps.github.io/maps"
)

var mapsAPIKey = maps.WithAPIKey("AIzaSyCWhNQ3cPHw7qFiiRTucVd61W4ZSzeUyTI")

// AllDestinations is used across the service to get all destinations based on a query, and returns results and pagination token.
func AllDestinations(ctx context.Context, query string) (*maps.PlacesSearchResponse, error) {
	c, err := maps.NewClient(mapsAPIKey)
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

// FindDestination returns a single destination with more details than AllDestinations.
func FindDestination(ctx context.Context, pid string) (*maps.PlaceDetailsResult, error) {
	c, err := maps.NewClient(mapsAPIKey)
	if err != nil {
		return nil, err
	}

	pdr := &maps.PlaceDetailsRequest{
		PlaceID: pid,
	}
	resp, err := c.PlaceDetails(ctx, pdr)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
