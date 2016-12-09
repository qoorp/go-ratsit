package ratsit

import (
	"errors"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

func generatePersonLookupURL(ssn string) (q string) {
	v := url.Values{}
	v.Add("SSN", ssn)
	q = apiUrl + "/personinformation?" + v.Encode()
	return
}

func generatePersonSearchURL(name string, location string, limit int) (q string) {
	v := url.Values{}
	v.Add("MaxNrRecords", strconv.Itoa(limit))
	v.Add("Where", location)
	v.Add("Who", name)
	q = apiUrl + "/personsearch?" + v.Encode()
	return
}

func authorizeRequest(r *http.Request, apiKey string, pkg string) {
	r.Header.Add("Authorization", apiKey)
	r.Header.Add("package", pkg)
}

func handleResponseError(r *http.Response) (err error) {
	if r.StatusCode == http.StatusInternalServerError {
		err = ErrInternalServer
	} else if r.StatusCode == http.StatusBadRequest {
		err = ErrInvalidInput
	} else if r.StatusCode == http.StatusUnauthorized {
		err = ErrInvalidCredentials
	}
	return
}

// ErrInvalidIdentificationNumber is an error returned when an identification number fails validation
var ErrInvalidIdentificationNumber = errors.New("invalid identification number")

// validateIdentityNumber takes a Swedish identification number (personnummer) and checks if it is a valid number
func validateIdentityNumber(number string) (err error) {
	// Check length/size of identityNumber
	if len(number) != 10 && len(number) != 12 {
		err = ErrInvalidIdentificationNumber
		return
	}
	if len(number) == 12 {
		number = number[2:]
	}

	// Check if string is only numbers
	_, err = strconv.Atoi(number)
	if err != nil {
		err = ErrInvalidIdentificationNumber
		return
	}

	// Check if identiy number passes checksum test as specified by Skatteverket: https://www.skatteverket.se/privat/sjalvservice/blanketterbroschyrer/broschyrer/info/704.4.39f16f103821c58f680007993.html
	var sum int
	for i, x := range number {
		if i != 9 {
			n := int(x - '0')
			m := math.Mod(float64(n), 2)
			if m == 0 {
				sum = sum + (n * 2)
			} else {
				sum = sum + (n * 1)
			}
		}
	}
	s := strconv.Itoa(sum)
	check := 10 - int(s[1]-'0')
	if check != int(number[9]) {
		err = ErrInvalidIdentificationNumber
	}
	return nil
}
