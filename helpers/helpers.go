package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"url_hasher/repository"
)

type HTTPClient struct {
	Client *http.Client
}

func NewAPIClient(client *http.Client) repository.APIClient {
	return &HTTPClient{Client: client}
}

func (h *HTTPClient) GetHTTP(requestURL string) ([]byte, error) {
	resp, err := h.Client.Get(requestURL)

	if err != nil {
		fmt.Println("Error occured in fetching response", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		fmt.Println("Fetched response is not 200, status code is: ", resp.StatusCode)
	}

	return body, nil
}
