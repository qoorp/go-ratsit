package ratsit

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const apiUrl = "https://api.ratsit.se/api/v1"

var (
	ErrInvalidInput       = errors.New("invalid input")
	ErrInternalServer     = errors.New("internal server error")
	ErrInvalidCredentials = errors.New("authenication failed")
)

// Client issuess request to the ratsit server
type Ratsit struct {
	apiKey string
	client *http.Client
}

func New(key string) (r Ratsit) {
	r.apiKey = key
	r.client = new(http.Client)
	return
}

// GetPerson returns a person from the database
func (r Ratsit) GetPerson(ssn string, pkg string) (p Person, err error) {

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

	err = json.Unmarshal(j, &p)

	return
}

func (r Ratsit) SearchPerson(name string, location string) (p []Person, err error) {

	// req, err := http.NewRequest(http.MethodGet, apiUrl+"/personsok")

	return
}
