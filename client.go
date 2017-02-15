package ratsit

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	// APIURL defines the Ratsit API URL for use in production
	APIURL          = "https://api.ratsit.se/api/v1"
	pkgPersonSearch = "personsok"
)

var (
	ErrInvalidInput       = errors.New("invalid input")
	ErrInternalServer     = errors.New("internal server error")
	ErrInvalidCredentials = errors.New("authenication failed")
)

// Ratsit is the the client
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
func (r Ratsit) GetPerson(ssn string, pkg string) (person Person, err error) {
	// TODO: Validate SSN and pkg
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
	err = json.Unmarshal(j, &person)
	return
}

// SearchPerson searches the Ratsit database for people with the name and location given in the parameters
func (r Ratsit) SearchPerson(name string, location string, limit int) (personSearchResults SearchResults, err error) {
	url := generatePersonSearchURL(name, location, limit)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	authorizeRequest(req, r.apiKey, pkgPersonSearch)
	resp, err := r.client.Do(req)
	if err != nil {
		return
	}
	err = handleResponseError(resp)
	if err != nil {
		return
	}
	jsonBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonBody, &personSearchResults)
	if err != nil {
		return
	}
	return
}
