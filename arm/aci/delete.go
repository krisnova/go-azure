package aci

import "github.com/kris-nova/go-azure/arm"

type DeleteContainerRequest struct {
}

type DeleteContainerResponse struct {
	arm.HTTPResponse
}

func (c *Client) Delete(Request *DeleteContainerRequest) (*DeleteContainerResponse, error) {

	return &DeleteContainerResponse{}, nil
}
