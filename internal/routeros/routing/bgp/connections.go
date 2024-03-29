package bgp

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type ListBGPConnectionsInput struct{}
type ListBGPConnectionsOutput struct{}

type GetBGPConnectionInput struct{}
type GetBGPConnectionOutput struct{}

type CreateBGPConnectionInput struct{}
type CreateBGPConnectionOutput struct{}

type UpdateBGPConnectionInput struct{}
type UpdateBGPConnectionOutput struct{}

type DeleteBGPConnectionInput struct{}
type DeleteBGPConnectionOutput struct{}

type Connections interface {
	ListBGPConnections(input *ListBGPConnectionsInput) (*ListBGPConnectionsOutput, error)
	GetBGPConnection(input *GetBGPConnectionInput) (*GetBGPConnectionOutput, error)
	CreateBGPConnection(input *CreateBGPConnectionInput) (*CreateBGPConnectionOutput, error)
	UpdateBGPConnection(input *UpdateBGPConnectionInput) (*UpdateBGPConnectionOutput, error)
	DeleteBGPConnection(input *DeleteBGPConnectionInput) (*DeleteBGPConnectionOutput, error)
}

type ConnectionsImpl struct {
	Client *resty.Client
}

func (c ConnectionsImpl) ListBGPConnections(input *ListBGPConnectionsInput) (*ListBGPConnectionsOutput, error) {
	return nil, errors.New("not implemented")
}

func (c ConnectionsImpl) GetBGPConnection(input *GetBGPConnectionInput) (*GetBGPConnectionOutput, error) {
	return nil, errors.New("not implemented")
}

func (c ConnectionsImpl) CreateBGPConnection(input *CreateBGPConnectionInput) (*CreateBGPConnectionOutput, error) {
	return nil, errors.New("not implemented")
}

func (c ConnectionsImpl) UpdateBGPConnection(input *UpdateBGPConnectionInput) (*UpdateBGPConnectionOutput, error) {
	return nil, errors.New("not implemented")
}

func (c ConnectionsImpl) DeleteBGPConnection(input *DeleteBGPConnectionInput) (*DeleteBGPConnectionOutput, error) {
	return nil, errors.New("not implemented")
}
