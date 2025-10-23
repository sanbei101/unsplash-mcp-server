package tools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/douglarek/unsplash-mcp-server/internal/api"
	"github.com/douglarek/unsplash-mcp-server/internal/config"
	"github.com/douglarek/unsplash-mcp-server/internal/models"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func NewSearchPhotosTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "search_photos",
		Description: "Search for Unsplash photos",
	}
}

func HandleSearchPhotos(ctx context.Context, req *mcp.CallToolRequest, input models.SearchPhotosRequest) (*mcp.CallToolResult, *models.SearchResult, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load configuration: %v", err)
	}

	client := api.NewClient(cfg)
	if input.Query == "" {
		return nil, nil, fmt.Errorf("query arg is required")
	}

	params := buildSearchParams(input)
	photos, err := client.SearchPhotos(ctx, params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to search photos: %v", err)
	}
	result := models.SearchResult{
		Photos: photos,
	}
	return nil, &result, nil
}

func buildSearchParams(args models.SearchPhotosRequest) url.Values {
	params := url.Values{}

	params.Add("query", args.Query)
	if args.Page == 0 {
		args.Page = 1
	}
	params.Add("page", strconv.Itoa(args.Page))
	if args.PerPage == 0 {
		args.PerPage = 5
	}
	params.Add("per_page", strconv.Itoa(args.PerPage))
	if args.OrderBy == "" {
		args.OrderBy = "relevant"
	}
	params.Add("order_by", args.OrderBy)
	if args.Color != "" {
		params.Add("color", args.Color)
	}
	if args.Orientation != "" {
		params.Add("orientation", args.Orientation)
	}

	return params
}
