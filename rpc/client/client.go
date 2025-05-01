package client

import (
	"log"
	"net/rpc"
)

type Client struct {
	rpcClient *rpc.Client
}

func New(c *rpc.Client) *Client {
	return &Client{rpcClient: c}
}

func (c *Client) logFailedMethodCall(method string, err error) {
	log.Printf("RPC: client failed to call method %s: %w", method, err)
}
