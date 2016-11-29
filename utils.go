package ratsit

import (
	"net/http"
	"net/url"
)

func generatePersonLookupURL(ssn string) (q string) {
	v := url.Values{}
	v.Add("SSN", ssn)
	q = apiUrl + "/personinformation?" + v.Encode()
	return
}

func authorizeRequest(r *http.Request, apiKey string, pkg string) {
	r.Header.Add("Authorization", apiKey)
	r.Header.Add("package", pkg)
}

func handleResponseError(r *http.Response) (err error) {
	if r.StatusCode == 500 {
		err = ErrInternalServer
	} else if r.StatusCode == 400 {
		err = ErrInvalidInput
	} else if r.StatusCode == 401 {
		err = ErrInvalidCredentials
	}
	return
}
