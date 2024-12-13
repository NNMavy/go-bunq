package bunq

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/OGKevin/go-bunq/model"
)

type deviceServerService service

func (d *deviceServerService) create() (*model.ResponseDeviceServer, error) {
	bodyStruct := model.RequestDeviceServer{
		Description:  d.client.description,
		Secret:       d.client.apiKey,
		PermittedIps: d.client.permittedIps,
	}
	bodyRaw, err := json.Marshal(bodyStruct)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(
		http.MethodPost,
		d.client.formatRequestURL(endpointDeviceServerCreate),
		bytes.NewBuffer(bodyRaw),
	)
	if err != nil {
		return nil, err
	}

	res, err := d.client.do(r)
	if err != nil {
		return nil, err
	}

	var resSessionServer model.ResponseDeviceServer
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&resSessionServer)

	return &resSessionServer, err
}
