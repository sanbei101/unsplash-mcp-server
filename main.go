package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "Unsplash MCP Server",
		Version: "1.0.0",
	}, nil)

	searchTool := NewSearchPhotosTool()
	mcp.AddTool(
		server,
		searchTool,
		HandleSearchPhotos,
	)
	// handler := mcp.NewStreamableHTTPHandler(
	// 	func(r *http.Request) *mcp.Server {
	// 		return server
	// 	},
	// 	&mcp.StreamableHTTPOptions{
	// 		JSONResponse: true,
	// 		Stateless:    true,
	// 	},
	// )
	// http.ListenAndServe(":8080", handler)
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
