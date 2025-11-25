package tools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lajosbencz/whoamimcp/internal/keys"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type WhoamiTextArgs struct {
}

func WhoamiText(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[WhoamiTextArgs]) (*mcp.CallToolResultFor[any], error) {

	name, ok := ctx.Value(keys.Name).(string)
	if !ok {
		return &mcp.CallToolResultFor[any]{}, fmt.Errorf("Internal error: failed to get name")
	}
	req, ok := ctx.Value(keys.HttpRequest).(*http.Request)
	if !ok {
		return &mcp.CallToolResultFor[any]{}, fmt.Errorf("Internal error: failed to get request")
	}

	info := GetWhoamiStructInfo(req, name)
	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: info.String(),
			},
		},
	}, nil
}
