package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// OperationTagResourceContractProperties represents the OperationTagResourceContractProperties schema from the OpenAPI specification
type OperationTagResourceContractProperties struct {
	Name string `json:"name,omitempty"` // Operation name.
	Urltemplate string `json:"urlTemplate,omitempty"` // Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
	Apiname string `json:"apiName,omitempty"` // Api Name.
	Apirevision string `json:"apiRevision,omitempty"` // Api Revision.
	Apiversion string `json:"apiVersion,omitempty"` // Api Version.
	Description string `json:"description,omitempty"` // Operation Description.
	Id string `json:"id,omitempty"` // Identifier of the operation in form /operations/{operationId}.
	Method string `json:"method,omitempty"` // A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
}

// ProductTagResourceContractProperties represents the ProductTagResourceContractProperties schema from the OpenAPI specification
type ProductTagResourceContractProperties struct {
	State string `json:"state,omitempty"` // whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
	Subscriptionrequired bool `json:"subscriptionRequired,omitempty"` // Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
	Subscriptionslimit int `json:"subscriptionsLimit,omitempty"` // Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of true.
	Terms string `json:"terms,omitempty"` // Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
	Approvalrequired bool `json:"approvalRequired,omitempty"` // whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of true.
	Description string `json:"description,omitempty"` // Product description. May include HTML formatting tags.
}

// TagResourceCollection represents the TagResourceCollection schema from the OpenAPI specification
type TagResourceCollection struct {
	Count int64 `json:"count,omitempty"` // Total record count number across all pages.
	Nextlink string `json:"nextLink,omitempty"` // Next page link if any.
	Value []TagResourceContract `json:"value,omitempty"` // Page values.
}

// TagResourceContract represents the TagResourceContract schema from the OpenAPI specification
type TagResourceContract struct {
	Operation OperationTagResourceContractProperties `json:"operation,omitempty"` // Operation Entity contract Properties.
	Product ProductTagResourceContractProperties `json:"product,omitempty"` // Product profile.
	Tag TagTagResourceContractProperties `json:"tag"` // Contract defining the Tag property in the Tag Resource Contract
	Api ApiTagResourceContractProperties `json:"api,omitempty"` // API contract properties for the Tag Resources.
}

// TagTagResourceContractProperties represents the TagTagResourceContractProperties schema from the OpenAPI specification
type TagTagResourceContractProperties struct {
	Id string `json:"id,omitempty"` // Tag identifier
	Name string `json:"name,omitempty"` // Tag Name
}

// ApiTagResourceContractProperties represents the ApiTagResourceContractProperties schema from the OpenAPI specification
type ApiTagResourceContractProperties struct {
	Apiversion string `json:"apiVersion,omitempty"` // Indicates the Version identifier of the API if the API is versioned
	Description string `json:"description,omitempty"` // Description of the API. May include HTML formatting tags.
	Iscurrent bool `json:"isCurrent,omitempty"` // Indicates if API revision is current api revision.
	Apirevisiondescription string `json:"apiRevisionDescription,omitempty"` // Description of the Api Revision.
	Apiversiondescription string `json:"apiVersionDescription,omitempty"` // Description of the Api Version.
	Authenticationsettings interface{} `json:"authenticationSettings,omitempty"` // API Authentication Settings.
	TypeField string `json:"type,omitempty"` // Type of API.
	Apiversionsetid string `json:"apiVersionSetId,omitempty"` // A resource identifier for the related ApiVersionSet.
	Isonline bool `json:"isOnline,omitempty"` // Indicates if API revision is accessible via the gateway.
	Subscriptionkeyparameternames interface{} `json:"subscriptionKeyParameterNames,omitempty"` // Subscription key parameter names details.
	Apirevision string `json:"apiRevision,omitempty"` // Describes the Revision of the Api. If no value is provided, default revision 1 is created
}
