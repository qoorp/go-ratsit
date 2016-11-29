package ratsit

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	apiKey := "testKey"
	c := New(apiKey)
	expected := Ratsit{
		apiKey: apiKey,
		client: new(http.Client),
	}
	if reflect.DeepEqual(c, expected) == false {
		t.Fail()
	}
}
