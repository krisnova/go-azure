package aci

import "github.com/kris-nova/go-azure/arm"

type CreateContainerRequest struct {
}

type CreateContainerResponse struct {
	arm.HTTPResponse
}

func (c *Client) Create(Request *CreateContainerRequest) (*CreateContainerResponse, error) {

	return &CreateContainerResponse{}, nil
}
