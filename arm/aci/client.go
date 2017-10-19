package aci

import "github.com/kris-nova/go-azure/arm"

type Client struct {
	authentication *arm.Authentication
}

func NewClient(auth *arm.Authentication) *Client {
	return &Client{
		authentication: auth,
	}
}
