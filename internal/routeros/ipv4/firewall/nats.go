package firewall

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type Nat struct {
}

type ListNatRulesInput struct{}
type ListNatRulesOutput struct{}

type GetNatRuleInput struct{}
type GetNatRuleOutput struct{}

type CreateNatRuleInput struct{}
type CreateNatRuleOutput struct{}

type UpdateNatRuleInput struct{}
type UpdateNatRuleOutput struct{}

type DeleteNatRuleInput struct{}
type DeleteNatRuleOutput struct{}

type Nats interface {
	ListNatRules(input *ListNatRulesInput) (*ListNatRulesOutput, error)
	GetNatRule(input *GetNatRuleInput) (*GetNatRuleOutput, error)
	CreateNatRule(input *CreateNatRuleInput) (*CreateNatRuleOutput, error)
	UpdateNatRule(input *UpdateNatRuleInput) (*UpdateNatRuleOutput, error)
	DeleteNatRule(input *DeleteNatRuleInput) (*DeleteNatRuleOutput, error)
}

type NatsImpl struct {
	Client *resty.Client
}

func (n NatsImpl) ListNatRules(input *ListNatRulesInput) (*ListNatRulesOutput, error) {
	return nil, errors.New("not implemented")
}

func (n NatsImpl) GetNatRule(input *GetNatRuleInput) (*GetNatRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n NatsImpl) CreateNatRule(input *CreateNatRuleInput) (*CreateNatRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n NatsImpl) UpdateNatRule(input *UpdateNatRuleInput) (*UpdateNatRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n NatsImpl) DeleteNatRule(input *DeleteNatRuleInput) (*DeleteNatRuleOutput, error) {
	return nil, errors.New("not implemented")
}
