package firewall

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type Mangle struct {
}

type ListMangleRulesInput struct{}
type ListMangleRulesOutput struct{}

type GetMangleRuleInput struct{}
type GetMangleRuleOutput struct{}

type CreateMangleRuleInput struct{}
type CreateMangleRuleOutput struct{}

type UpdateMangleRuleInput struct{}
type UpdateMangleRuleOutput struct{}

type DeleteMangleRuleInput struct{}
type DeleteMangleRuleOutput struct{}

type Mangles interface {
	ListMangleRules(input *ListMangleRulesInput) (*ListMangleRulesOutput, error)
	GetMangleRule(input *GetMangleRuleInput) (*GetMangleRuleOutput, error)
	CreateMangleRule(input *CreateMangleRuleInput) (*CreateMangleRuleOutput, error)
	UpdateMangleRule(input *UpdateMangleRuleInput) (*UpdateMangleRuleOutput, error)
	DeleteMangleRule(input *DeleteMangleRuleInput) (*DeleteMangleRuleOutput, error)
}

type ManglesImpl struct {
	Client *resty.Client
}

func (n ManglesImpl) ListMangleRules(input *ListMangleRulesInput) (*ListMangleRulesOutput, error) {
	return nil, errors.New("not implemented")
}

func (n ManglesImpl) GetMangleRule(input *GetMangleRuleInput) (*GetMangleRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n ManglesImpl) CreateMangleRule(input *CreateMangleRuleInput) (*CreateMangleRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n ManglesImpl) UpdateMangleRule(input *UpdateMangleRuleInput) (*UpdateMangleRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n ManglesImpl) DeleteMangleRule(input *DeleteMangleRuleInput) (*DeleteMangleRuleOutput, error) {
	return nil, errors.New("not implemented")
}
