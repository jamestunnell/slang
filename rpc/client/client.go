package client

import (
	"net/rpc"
)

type Client struct {
	rpcClient *rpc.Client
}

func New(c *rpc.Client) *Client {
	return &Client{rpcClient: c}
}
