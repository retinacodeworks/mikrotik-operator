package firewall

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type Filter struct {
}

type ListFilterRulesInput struct{}
type ListFilterRulesOutput struct{}

type GetFilterRuleInput struct{}
type GetFilterRuleOutput struct{}

type CreateFilterRuleInput struct{}
type CreateFilterRuleOutput struct{}

type UpdateFilterRuleInput struct{}
type UpdateFilterRuleOutput struct{}

type DeleteFilterRuleInput struct{}
type DeleteFilterRuleOutput struct{}

type Filters interface {
	ListFilterRules(input *ListFilterRulesInput) (*ListFilterRulesOutput, error)
	GetFilterRule(input *GetFilterRuleInput) (*GetFilterRuleOutput, error)
	CreateFilterRule(input *CreateFilterRuleInput) (*CreateFilterRuleOutput, error)
	UpdateFilterRule(input *UpdateFilterRuleInput) (*UpdateFilterRuleOutput, error)
	DeleteFilterRule(input *DeleteFilterRuleInput) (*DeleteFilterRuleOutput, error)
}

type FiltersImpl struct {
	Client *resty.Client
}

func (n FiltersImpl) ListFilterRules(input *ListFilterRulesInput) (*ListFilterRulesOutput, error) {
	return nil, errors.New("not implemented")
}

func (n FiltersImpl) GetFilterRule(input *GetFilterRuleInput) (*GetFilterRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n FiltersImpl) CreateFilterRule(input *CreateFilterRuleInput) (*CreateFilterRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n FiltersImpl) UpdateFilterRule(input *UpdateFilterRuleInput) (*UpdateFilterRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n FiltersImpl) DeleteFilterRule(input *DeleteFilterRuleInput) (*DeleteFilterRuleOutput, error) {
	return nil, errors.New("not implemented")
}
