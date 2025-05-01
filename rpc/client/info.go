package client

import (
	"github.com/google/uuid"
	"github.com/jamestunnell/slang/rpc/models"
)

func (client *Client) IsRunning() bool {
	const method = "Info.IsRunning"

	var reply bool

	if err := client.rpcClient.Call(method, &models.Empty{}, &reply); err != nil {
		client.logFailedMethodCall(method, err)

		return false
	}

	return reply
}

func (client *Client) GetName() string {
	const method = "Info.GetName"

	var reply string

	if err := client.rpcClient.Call(method, &models.Empty{}, &reply); err != nil {
		client.logFailedMethodCall(method, err)

		return ""
	}

	return reply
}

func (client *Client) GetID() uuid.UUID {
	const method = "Info.GetID"

	var reply uuid.UUID

	if err := client.rpcClient.Call(method, &models.Empty{}, &reply); err != nil {
		client.logFailedMethodCall(method, err)

		return uuid.Nil
	}

	return reply
}
