package main_test

import (
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestGetUserName_EscapePlus(t *testing.T) {
	name := "john+doe"
	name = url.QueryEscape(name)
	t.Log(name)

	u := "http://localhost:8080/getUserName?name=" + name
	resp, err := http.Get(u)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := string(bodyBytes)
		want := "john+doe"
		if got != want {
			t.Errorf("name = %q; want %q", got, want)
		}
	}
}

func TestGetUserName_NotEscapePlus(t *testing.T) {
	name := "john+doe"
	t.Log(name)

	u := "http://localhost:8080/getUserName?name=" + name
	resp, err := http.Get(u)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := string(bodyBytes)
		want := "john doe"
		if got != want {
			t.Errorf("name = %q; want %q", got, want)
		}
	}
}

func TestGetUserName_EscapeSpace(t *testing.T) {
	name := "john doe"
	name = url.QueryEscape(name)
	t.Log(name)

	u := "http://localhost:8080/getUserName?name=" + name
	resp, err := http.Get(u)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := string(bodyBytes)
		want := "john doe"
		if got != want {
			t.Errorf("name = %q; want %q", got, want)
		}
	}
}
