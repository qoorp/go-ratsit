package ratsit

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	// APIURL defines the Ratsit API URL for use in production
	APIURL           = "https://api.checkbiz.se/api/v1"
	pkgPersonSearch  = "personsok"
	pkgCompanySearch = "foretagsok"
)

var (
	// ErrInvalidInput ...
	ErrInvalidInput = errors.New("invalid input")
	// ErrInternalServer ...
	ErrInternalServer = errors.New("internal server error")
	// ErrInvalidCredentials ...
	ErrInvalidCredentials = errors.New("authentication failed")
)

// Ratsit is the the client
type Ratsit struct {
	apiURL string
	apiKey string
	client *http.Client
}

// New creates a new client to interact with the Ratsit API
func New(apiURL string, key string) (r Ratsit) {
	r.apiURL = apiURL
	r.apiKey = key
	r.client = new(http.Client)
	return
}

// GetPerson returns a person from the database by looking up their unique personnummer
func (r *Ratsit) GetPerson(ssn string, pkg string) (person Person, err error) {
	// TODO: Validate SSN and pkg
	url := generatePersonInformationURL(r.apiURL, ssn)
	body, err := doHTTPRequest(r.client, url, r.apiKey, pkg)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &person)
	return
}

// SearchPerson searches the Ratsit database for people with the name and location given in the parameters
func (r *Ratsit) SearchPerson(name string, location string, limit int, recordFrom int) (personSearchResults SearchResults, err error) {
	url := generatePersonSearchURL(r.apiURL, name, location, limit, recordFrom)
	body, err := doHTTPRequest(r.client, url, r.apiKey, pkgPersonSearch)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &personSearchResults)
	return
}

// GetCompany returns a company from the database by looking up their unique organization number
func (r *Ratsit) GetCompany(organizationNumber string, pkg string) (company CompanyInformationResponse, err error) {
	// TODO: Validate organizationNumber and pkg
	url := generateCompanyInformationURL(r.apiURL, organizationNumber)
	body, err := doHTTPRequest(r.client, url, r.apiKey, pkg)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &company)
	return
}

// SearchCompany searches the Ratsit database for companies with the name and location given in the parameters
func (r *Ratsit) SearchCompany(name string, location string, limit int, recordFrom int) (companySearchResults CompanySearchResults, err error) {
	url := generateCompanySearchURL(r.apiURL, name, location, limit, recordFrom)
	body, err := doHTTPRequest(r.client, url, r.apiKey, pkgCompanySearch)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &companySearchResults)
	return
}

func doHTTPRequest(client *http.Client, url string, apiKey string, pkg string) (body []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	authorizeRequest(req, apiKey, pkg)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	err = handleResponseError(resp)
	if err != nil {
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	return
}
