package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"url_hasher/helpers"
)

var formatURLTestCases = []struct {
	TestURL         string
	ExpFormattedURL string
}{
	{"google.com", "http://google.com"},
	{"http://google.com", "http://google.com"},
	{"dfgd", "http://dfgd"},
	{"https://google.com", "https://google.com"},
}

var hashResponseTestCases = []struct {
	ExpectedResponse string
	ExpectedHash     string
}{
	{"OK", "e0aa021e21dddbd6d8cecec71e9cf564"},
	{"Hello world", "3e25960a79dbc69b674cd4ec67a72c62"},
}

func TestGetResponseAndHash(t *testing.T) {
	for _, testCase := range hashResponseTestCases {
		var wg sync.WaitGroup
		wg.Add(1)

		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			fmt.Println("test server")
			rw.Write([]byte(testCase.ExpectedResponse))
		}))
		defer server.Close()

		testAPIClient := helpers.NewAPIClient(server.Client())

		actualResult := GetResponseAndHash(server.URL, testAPIClient, &wg)
		if actualResult != testCase.ExpectedHash {
			t.Errorf("GetResponseAndHash() with args %v : FAILED, expected %v but got value %v", server.URL, testCase.ExpectedHash, actualResult)
		} else {
			t.Logf("GetResponseAndHash() with args %v : PASSED, expected %v and got value %v", server.URL, testCase.ExpectedHash, actualResult)
		}
	}
}

func TestFormatURL(t *testing.T) {
	for _, testCase := range formatURLTestCases {
		actualResult := FormatURL(testCase.TestURL)

		if actualResult != testCase.ExpFormattedURL {
			t.Errorf("FormatURL() with args %v : FAILED, expected %v but got value %v", testCase.TestURL, testCase.ExpFormattedURL, actualResult)
		} else {
			t.Logf("FormatURL() with args %v : PASSED, expected %v and got value %v", testCase.TestURL, testCase.ExpFormattedURL, actualResult)
		}
	}
}
