package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_tagresources "github.com/apimanagementclient/mcp-server/tools/tagresources"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_tagresources.CreateTagresource_listbyserviceTool(cfg),
	}
}
