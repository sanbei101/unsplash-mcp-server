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
	inputSchema := map[string]any{
		"type":     "object",
		"required": []string{"query"},
		"properties": map[string]any{
			"query": map[string]any{
				"type":        "string",
				"description": "搜索关键词（例如：'山峰'、'日落'）",
			},
			"page": map[string]any{
				"type":        "integer",
				"minimum":     1,
				"default":     1,
				"description": "页码（从 1 开始）",
			},
			"per_page": map[string]any{
				"type":        "integer",
				"minimum":     1,
				"maximum":     30,
				"default":     5,
				"description": "每页返回的照片数量（最多 30 张）",
			},
			"order_by": map[string]any{
				"type":        "string",
				"enum":        []string{"relevant", "latest"},
				"default":     "relevant",
				"description": "结果排序方式",
			},
			"color": map[string]any{
				"type":        "string",
				"description": "按颜色筛选（例如：'黑色'、'红色'、'蓝色'）",
			},
			"orientation": map[string]any{
				"type":        "string",
				"enum":        []string{"landscape", "portrait", "squarish"},
				"description": "照片方向",
			},
		},
		"additionalProperties": false,
	}

	outputSchema := map[string]any{
		"type":     "object",
		"required": []string{"photos"},
		"properties": map[string]any{
			"photos": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type":     "object",
					"required": []string{"id", "description", "urls", "width", "height"},
					"properties": map[string]any{
						"id": map[string]any{
							"type": "string",
						},
						"description": map[string]any{
							"type": []any{"string", nil},
						},
						"width": map[string]any{
							"type": "integer",
						},
						"height": map[string]any{
							"type": "integer",
						},
						"urls": map[string]any{
							"type":     "object",
							"required": []string{"raw", "full", "regular", "small", "thumb"},
							"properties": map[string]any{
								"raw":     map[string]any{"type": "string"},
								"full":    map[string]any{"type": "string"},
								"regular": map[string]any{"type": "string"},
								"small":   map[string]any{"type": "string"},
								"thumb":   map[string]any{"type": "string"},
							},
							"additionalProperties": false,
							"description":          "不同尺寸的照片 URL",
						},
					},
					"additionalProperties": false,
					"description":          "单张照片信息",
				},
			},
		},
		"additionalProperties": false,
		"description":          "Unsplash 搜索结果",
	}
	return &mcp.Tool{
		Name:         "search_photos",
		Description:  "Search for Unsplash photos",
		InputSchema:  inputSchema,
		OutputSchema: outputSchema,
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
