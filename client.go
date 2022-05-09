package shopify

import (
	"os"

	"github.com/r0busta/graphql"
	log "github.com/sirupsen/logrus"
	graphqlclient "github.com/vstarapp/go-shopify-graphql/v6/graphql"
)

const (
	defaultShopifyAPIVersion = "2022-04"
)

type Client struct {
	gql graphql.GraphQL

	Product       ProductService
	Variant       VariantService
	Inventory     InventoryService
	Collection    CollectionService
	Order         OrderService
	Fulfillment   FulfillmentService
	Location      LocationService
	Metafield     MetafieldService
	BulkOperation BulkOperationService
}

type Option func(shopClient *Client)

func WithGraphQLClient(gql graphql.GraphQL) Option {
	return func(c *Client) {
		c.gql = gql
	}
}

func NewDefaultClient(opts ...Option) *Client {
	accessToken := os.Getenv("STORE_ACCESS_TOKEN")
	storeName := os.Getenv("STORE_NAME")
	if accessToken == "" || storeName == "" {
		log.Fatalln("Shopify app API Key and/or Password and/or Store Name not set")
	}

	return NewClient(accessToken, storeName, "", opts...)
}

func NewClient(accessToken string, storeName string, apiVersion string, opts ...Option) *Client {
	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	if c.gql == nil {
		c.gql = newShopifyGraphQLClientWithToken(accessToken, storeName, apiVersion)
	}

	c.Product = &ProductServiceOp{client: c}
	c.Variant = &VariantServiceOp{client: c}
	c.Inventory = &InventoryServiceOp{client: c}
	c.Collection = &CollectionServiceOp{client: c}
	c.Order = &OrderServiceOp{client: c}
	c.Fulfillment = &FulfillmentServiceOp{client: c}
	c.Location = &LocationServiceOp{client: c}
	c.Metafield = &MetafieldServiceOp{client: c}
	c.BulkOperation = &BulkOperationServiceOp{client: c}

	return c
}

func newShopifyGraphQLClient(apiKey string, password string, storeName string) *graphql.Client {
	opts := []graphqlclient.Option{
		graphqlclient.WithVersion(defaultShopifyAPIVersion),
		graphqlclient.WithPrivateAppAuth(apiKey, password),
	}
	return graphqlclient.NewClient(storeName, opts...)
}

func newShopifyGraphQLClientWithToken(accessToken string, storeName string, apiVersion string) *graphql.Client {
	if apiVersion == "" {
		apiVersion = defaultShopifyAPIVersion
	}
	opts := []graphqlclient.Option{
		graphqlclient.WithVersion(apiVersion),
		graphqlclient.WithToken(accessToken),
	}
	return graphqlclient.NewClient(storeName, opts...)
}

func (c *Client) GraphQLClient() graphql.GraphQL {
	return c.gql
}
