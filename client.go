package ratsit

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	// APIURL defines the Ratsit API URL for use in production
	APIURL           = "https://api.ratsit.se/api/v1"
	pkgPersonSearch  = "personsok"
	pkgCompanySearch = "foretagsok"
)

var (
	ErrInvalidInput       = errors.New("invalid input")
	ErrInternalServer     = errors.New("internal server error")
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
	url := generatePersonLookupURL(r.apiURL, ssn)

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
func (r *Ratsit) SearchPerson(name string, location string, limit int, recordFrom int) (personSearchResults SearchResults, err error) {
	url := generatePersonSearchURL(r.apiURL, name, location, limit, recordFrom)
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
	return
}

// SearchCompany searches the Ratsit database for companies with the name and location given in the parameters
func (r *Ratsit) SearchCompany(name string, location string, limit int, recordFrom int) (companySearchResults CompanySearchResults, err error) {
	url := generateCompanySearchURL(r.apiURL, name, location, limit, recordFrom)
	fmt.Println("SearchCompany: " + url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return
	}
	authorizeRequest(req, r.apiKey, pkgCompanySearch)
	resp, err := r.client.Do(req)
	if err != nil {
		log.Println(err)
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
	err = json.Unmarshal(jsonBody, &companySearchResults)
	return
}
