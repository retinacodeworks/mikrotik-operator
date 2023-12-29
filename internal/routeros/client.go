package routeros

import (
	"crypto/tls"
	"errors"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type Options struct {
	Address string
	Headers map[string]string
	Tls     *tls.Config
}

type RouterOS struct {
	Client *resty.Client
}

func New(opts *Options) *RouterOS {
	client := resty.New()
	client.SetBaseURL(opts.Address)
	client.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		switch response.StatusCode() {
		case http.StatusInternalServerError:
			return errors.New("an unexpected error occured")
		case http.StatusBadRequest:
			return errors.New(response.String())
		default:
			return nil
		}
	})
	client.SetHeaders(opts.Headers)
	client.SetTLSClientConfig(opts.Tls)

	return &RouterOS{
		Client: client,
	}
}

type VersionInfoResponse struct {
	BoardName string `json:"board-name"`
	Version   string `json:"version"`
}

func (r *RouterOS) GetVersionInfo() (VersionInfoResponse, error) {
	var resp VersionInfoResponse
	_, err := r.Client.R().
		SetResult(&resp).
		Get("/rest/system/resource")
	return resp, err
}
