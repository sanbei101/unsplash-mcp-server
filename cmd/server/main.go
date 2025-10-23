package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/douglarek/unsplash-mcp-server/pkg/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "Unsplash MCP Server",
		Version: "1.0.0",
	}, nil)

	searchTool := tools.NewSearchPhotosTool()
	mcp.AddTool(
		server,
		searchTool,
		tools.HandleSearchPhotos,
	)
	if err := server.Run(context.Background(), &mcp.StreamableServerTransport{}); err != nil {
		slog.Error("Server error", "error", err)
		os.Exit(1)
	}
}
