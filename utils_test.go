package ratsit

import (
	"net/http"
	"testing"
)

func TestPersonLookupURLGeneration(t *testing.T) {
	expected := apiUrl + "/personinformation?SSN=0123456789"
	actual := generatePersonLookupURL("0123456789")
	if actual != expected {
		t.Fail()
	}
}

func TestPersonSearchURLGeneration(t *testing.T) {
	expected := apiUrl + "/personsearch?MaxNrRecords=10&Where=Stockholm&Who=Svensson"
	actual := generatePersonSearchURL("Svensson", "Stockholm", 10)
	if actual != expected {
		t.Fail()
	}
}

func TestRequestAuthorization(t *testing.T) {

	apikey := "testkey"
	pkg := "testpkg"

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Error(err)
	}
	authorizeRequest(r, apikey, pkg)
	if r.Header.Get("Authorization") != apikey {
		t.Fail()
	} else if r.Header.Get("package") != pkg {
		t.Fail()
	}
}

func TestResponseError500(t *testing.T) {
	r := new(http.Response)
	r.StatusCode = 500

	err := handleResponseError(r)
	if err != ErrInternalServer {
		t.Fail()
	}
}

func TestResponseError400(t *testing.T) {
	r := new(http.Response)
	r.StatusCode = 400

	err := handleResponseError(r)
	if err != ErrInvalidInput {
		t.Fail()
	}
}

func TestResponseError401(t *testing.T) {
	r := new(http.Response)
	r.StatusCode = 401

	err := handleResponseError(r)
	if err != ErrInvalidCredentials {
		t.Fail()
	}
}
