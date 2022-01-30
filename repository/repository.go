package repository

type APIClient interface {
	GetHTTP(requestURL string) ([]byte, error)
}
