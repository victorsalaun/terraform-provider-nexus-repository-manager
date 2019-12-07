package nexus

import (
	"crypto/tls"
	"io"
	"net/http"
)

var baseUrl = "/service/rest"
var usersServiceUrl = "/beta/security/users"

func nexusClient(config Config, method string, url string, body io.Reader) (*http.Response, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, _ := http.NewRequest(method, config.Url+baseUrl+url, body)
	req.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4xMjM=")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// USERS

func userCreate(config Config, body io.Reader) (*http.Response, error) {
	return nexusClient(config, http.MethodPost, usersServiceUrl, body)
}

func userRead(config Config, userId string) (*http.Response, error) {
	url := usersServiceUrl
	if userId != "" {
		url += "/" + userId
	}
	return nexusClient(config, http.MethodGet, url, nil)
}

func userUpdate(config Config, userId string, body io.Reader) (*http.Response, error) {
	url := usersServiceUrl + "/" + userId
	return nexusClient(config, http.MethodPut, url, body)
}

func userDelete(config Config, userId string) (*http.Response, error) {
	var url = usersServiceUrl + "/" + userId
	return nexusClient(config, http.MethodDelete, url, nil)
}
