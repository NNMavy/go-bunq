package bunq

import (
	"fmt"
	"net/http"

	"github.com/OGKevin/go-bunq/model"
	"github.com/pkg/errors"
)

type accountService service

func (a *accountService) GetAllMonetaryAccountBank(params ...model.QueryParam) (*model.ResponseMonetaryAccountBankGet, error) {
	userID, err := a.client.GetUserID()
	if err != nil {
		return nil, err
	}

	res, err := a.client.preformRequest(http.MethodGet, a.client.formatRequestURL(fmt.Sprintf(endpointMonetaryAccountBankListing, userID)), nil, params...)
	if err != nil {
		return nil, errors.Wrap(err, "bunq: request to get all MA bank failed")
	}

	var resMaGet model.ResponseMonetaryAccountBankGet

	return &resMaGet, a.client.parseResponse(res, &resMaGet)
}

func (a *accountService) GetMonetaryAccountBank(id int) (*model.ResponseMonetaryAccountBankGet, error) {
	userID, err := a.client.GetUserID()
	if err != nil {
		return nil, err
	}

	res, err := a.client.preformRequest(http.MethodGet, a.client.formatRequestURL(fmt.Sprintf(endpointMonetaryAccountBankGet, userID, id)), nil)
	if err != nil {
		return nil, errors.Wrap(err, "bunq: request to get MA bank failed")
	}

	var resMaGet model.ResponseMonetaryAccountBankGet

	return &resMaGet, a.client.parseResponse(res, &resMaGet)
}

func (a *accountService) GetAllMonetaryAccountSaving(params ...model.QueryParam) (*model.ResponseMonetaryAccountSavingGet, error) {
	userID, err := a.client.GetUserID()
	if err != nil {
		return nil, err
	}

	res, err := a.client.preformRequest(http.MethodGet, a.client.formatRequestURL(fmt.Sprintf(endpointMonetaryAccountSavingsListing, userID)), nil, params...)
	if err != nil {
		return nil, errors.Wrap(err, "bunq: request to get all MA saving failed")
	}

	var resStruct model.ResponseMonetaryAccountSavingGet

	return &resStruct, a.client.parseResponse(res, &resStruct)
}

func (a *accountService) GetMonetaryAccountSaving(id int) (*model.ResponseMonetaryAccountSavingGet, error) {
	userID, err := a.client.GetUserID()
	if err != nil {
		return nil, err
	}

	res, err := a.client.preformRequest(http.MethodGet, a.client.formatRequestURL(fmt.Sprintf(endpointMonetaryAccountSavingsGet, userID, id)), nil)
	if err != nil {
		return nil, errors.Wrap(err, "bunq: request to get MA saving failed")
	}

	var resStruct model.ResponseMonetaryAccountSavingGet

	return &resStruct, a.client.parseResponse(res, &resStruct)
}

func (a *accountService) GetAllMonetaryAccountJoint() (*model.ResponseMonetaryAccountJointGet, error) {
	userID, err := a.client.GetUserID()
	if err != nil {
		return nil, err
	}

	res, err := a.client.preformRequest(http.MethodGet, a.client.formatRequestURL(fmt.Sprintf(endpointMonetaryAccountJointListing, userID)), nil)
	if err != nil {
		return nil, errors.Wrap(err, "bunq: request to get all MA bank failed")
	}

	var resMaGet model.ResponseMonetaryAccountJointGet

	return &resMaGet, a.client.parseResponse(res, &resMaGet)
}

func (a *accountService) GetMonetaryAccountJoint(id int) (*model.ResponseMonetaryAccountJointGet, error) {
	userID, err := a.client.GetUserID()
	if err != nil {
		return nil, err
	}

	res, err := a.client.preformRequest(http.MethodGet, a.client.formatRequestURL(fmt.Sprintf(endpointMonetaryAccountJointGet, userID, id)), nil)
	if err != nil {
		return nil, errors.Wrap(err, "bunq: request to get MA bank failed")
	}

	var resMaGet model.ResponseMonetaryAccountJointGet

	return &resMaGet, a.client.parseResponse(res, &resMaGet)
}
