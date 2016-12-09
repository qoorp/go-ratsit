package ratsit

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const apiUrl = "https://api.ratsit.se/api/v1"
const PackagePersonadress = "personadress"

var (
	ErrInvalidInput       = errors.New("invalid input")
	ErrInternalServer     = errors.New("internal server error")
	ErrInvalidCredentials = errors.New("authenication failed")
)

type Ratsit struct {
	apiKey string
	client *http.Client
}

// New creates a new client to interact with the Ratsit API
func New(key string) (r Ratsit) {
	r.apiKey = key
	r.client = new(http.Client)
	return
}

// GetPerson returns a person from the database by looking up their unique personnummer
func (r Ratsit) GetPerson(ssn string, pkg string) (p Person, err error) {

	// TODO: Validate SSN and pkg
	err = validateIdentityNumber(ssn)
	if err != nil {
		return
	}

	url := generatePersonLookupURL(ssn)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	authorizeRequest(req, r.apiKey, pkg)

	resp, err := r.client.Do(req)
	if err != nil {
		return
	}

	err = handleResponseError(resp)
	if err != nil {
		return
	}

	j, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	err = json.Unmarshal(j, &p)

	return
}

var ErrPersonNotFound = errors.New("person with specified name and address not found in records")

// SearchPerson searches the Ratsit database for people with the name and location given in the parameters
func (r Ratsit) SearchPerson(name string, location string, limit int) (p []PersonBasic, err error) {
	url := generatePersonSearchURL(name, location, limit)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	authorizeRequest(req, r.apiKey, "personsok")

	resp, err := r.client.Do(req)
	if err != nil {
		return
	}

	err = handleResponseError(resp)
	if err != nil {
		return
	}

	j, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	var data SearchResults

	err = json.Unmarshal(j, &data)
	if err != nil {
		return
	}

	if data.ExtendedResult.TotalRecordsFound == 0 {
		err = ErrPersonNotFound
		return
	}

	p = data.ExtendedResult.Records

	return
}
