package routeros

type ListVlanResponse struct {
}

type CreateVlanRequest struct {
}

type CreateVlanResponse struct {
}

func (r *RouterOS) ListVlans() (ListVlanResponse, error) {
	var resp ListVlanResponse
	_, err := r.Client.R().
		SetResult(resp).
		SetQueryParams(map[string]string{
			"type": "vlan",
		}).
		Get("/interface/print")
	return resp, err
}

func (r *RouterOS) CreateVlan(req *CreateVlanRequest) (*CreateVlanResponse, error) {
	return nil, nil
}
