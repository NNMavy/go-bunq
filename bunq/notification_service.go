package bunq

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OGKevin/go-bunq/model"
	"github.com/pkg/errors"
	//"github.com/pkg/errors"
)

type notificationService service

func (n *notificationService) GetNotificationFilterUrl(params ...model.QueryParam) (*model.ResponseNotificationFilterUrlGet, error) {
	userID, err := n.client.GetUserID()
	if err != nil {
		return nil, err
	}

	res, err := n.client.preformRequest(http.MethodGet, n.client.formatRequestURL(fmt.Sprintf(endpointUserNotificationFilterUrl, userID)), nil, params...)
	if err != nil {
		return nil, errors.Wrap(err, "bunq: request to get notification filters failed")
	}

	var resStruct model.ResponseNotificationFilterUrlGet

	return &resStruct, n.client.parseResponse(res, &resStruct)
}

func (n *notificationService) CreateNotificationFilterUrl(create model.NotificationFilterUrlCreate) (*model.ResponseBunqID, error) {
	userID, err := n.client.GetUserID()
	if err != nil {
		return nil, err
	}

	bodyRaw, err := json.Marshal(create)
	if err != nil {
		return nil, errors.Wrap(err, "bunq: could not marshal body")
	}

	return n.client.doCURequest(n.client.formatRequestURL(fmt.Sprintf(endpointUserNotificationFilterUrl, userID)), bodyRaw, http.MethodPost)
}
