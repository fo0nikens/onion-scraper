package main

import (
	"os"
	"strings"
	"testing"
)

func TestClient(t *testing.T) {

	client, err := NewClient("", 0)
	if err != nil {
		t.Error(err)
		return
	}

	// official tor onion link
	resp, err := client.Request(Address("expyuzz4wqqyqhjn"))
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Title: %s\nDescription: %s\n", resp.Title, resp.Description)

	resp, err = client.Request(Address("bla"))
	if err == nil {
		t.Error("We should get error, insted we go response:", *resp)
	}
}

func TestStore(t *testing.T) {
	_, err := NewStore().CSV("")
	if err == nil {
		t.Error("Filename empty but we got a file!")
		return
	}

	_, err = NewStore().CSV("testing.csv")
	if err != nil {
		t.Error(err)
		return
	}

	// file should be created
	// now remove it
	os.Remove("testing.csv")

}

func TestGenerator(t *testing.T) {
	checkOnionURL(t, RandOnionV2())
	checkOnionURL(t, RandOnionV3())
}

func checkOnionURL(t *testing.T, addr Address) {
	if strings.Contains(string(addr), "http://") {
		t.Error("It should not contain http", addr)
		return
	}

	if !strings.Contains(addr.Addr(), "http://") {
		t.Error("Format onion url wrong", addr.Addr())
		return
	}
}
