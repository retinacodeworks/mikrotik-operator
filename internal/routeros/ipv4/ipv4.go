package ipv4

import (
	"github.com/go-resty/resty/v2"
)

type Ipv4 struct {
	Client *resty.Client
}

func (i *Ipv4) Addresses() Addresses {
	return AddressesImpl{Client: i.Client}
}
