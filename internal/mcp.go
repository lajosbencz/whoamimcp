package internal

import (
	"context"
	"net/http"
	"time"

	"github.com/lajosbencz/whoamimcp/internal/keys"
	"github.com/lajosbencz/whoamimcp/internal/prompts"
	"github.com/lajosbencz/whoamimcp/internal/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func NewMcpHandler(name string, options *mcp.StreamableHTTPOptions) *mcp.StreamableHTTPHandler {
	return mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
		requestCtx := context.WithValue(req.Context(), keys.HttpRequest, req)
		requestCtx = context.WithValue(requestCtx, keys.Name, name)

		*req = *req.WithContext(requestCtx)

		options := mcp.ServerOptions{
			KeepAlive: time.Minute * 5,
		}
		server := mcp.NewServer(&mcp.Implementation{
			Name:    name,
			Version: "0.1.0",
		}, &options)

		mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: "Say hi with hostname and server info"}, tools.SayHi)
		mcp.AddTool(server, &mcp.Tool{Name: "whoami_text", Description: "Get text format system and request information"}, tools.WhoamiText)
		mcp.AddTool(server, &mcp.Tool{Name: "whoami_struct", Description: "Get structured system and request information"}, tools.WhoamiStruct)
		mcp.AddTool(server, &mcp.Tool{Name: "raise_error", Description: "Simulates an error on the MCP server"}, tools.RaiseError)
		server.AddPrompt(&mcp.Prompt{Name: "greet", Description: "Greeting prompt"}, prompts.PromptHi)

		return server
	}, options)
}
