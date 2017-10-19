package aci

import "github.com/kris-nova/go-azure/arm"

type GetContainerRequest struct {
}

type GetContainerResponse struct {
	arm.HTTPResponse
}

func (c *Client) Get(Request *GetContainerRequest) (*GetContainerResponse, error) {

	return &GetContainerResponse{}, nil
}
