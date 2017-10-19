package aci

type GetContainerInput struct {
}

type GetContainerOutput struct {
}

func (c *Client) Get(input *GetContainerInput) (*GetContainerOutput, error) {

	return &GetContainerOutput{}, nil
}
