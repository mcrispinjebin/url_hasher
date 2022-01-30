package service

import (
	"crypto/md5"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
	"url_hasher/helpers"
	"url_hasher/repository"
)

var apiClient = helpers.NewAPIClient(&http.Client{})

func ParseRequests() {
	var req []string
	parallelCount := flag.Int("parallel", 10, "parallel call count")
	flag.Parse()

	if os.Args[1] == "-parallel" {
		req = os.Args[3:]
	} else {
		req = os.Args[1:]
	}

	ProcessRequests(req, *parallelCount)
}

func ProcessRequests(urlSlice []string, parallelCount int) {
	var wg sync.WaitGroup

	for i := 0; i < len(urlSlice); i++ {
		wg.Add(1)
		go GetResponseAndHash(urlSlice[i], apiClient, &wg)

		if (i+1)%parallelCount == 0 {
			wg.Wait()
		}
	}
	// if parallel count > total requests to wait for all to get done
	wg.Wait()
}

func GetResponseAndHash(url string, clientInterface repository.APIClient, wg *sync.WaitGroup) string {
	defer wg.Done()
	parsedURL := FormatURL(url)

	respBody, err := clientInterface.GetHTTP(parsedURL)

	if err != nil {
		return ""
	}

	hashedValue := fmt.Sprintf("%x", md5.Sum(respBody))

	fmt.Println(url, " ", hashedValue)

	return hashedValue
}

func FormatURL(requestURL string) string {
	parsedURL, _ := url.Parse(requestURL)

	if !parsedURL.IsAbs() {
		parsedURL.Scheme = "http"
	}

	return fmt.Sprintf("%v", parsedURL)
}
