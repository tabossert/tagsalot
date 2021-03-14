package tagsalot

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

const (
	dockerHubAPIVersion = "DOCKER_HUB_API_VERSION"
	dockerHubURL        = "DOCKER_HUB_URL"
	defaultHubURL       = "https://registry.hub.docker.com"
	defaultAPIVersion   = "1"
	mediaType           = "application/json"
)

// Client for Dockerhub API
type DockerHubClient struct {
	Client     *http.Client
	UserAgent  string
	APIVersion string
	HubURL     *url.URL
}

// NewHTTPClient create new docker client
func NewHTTPClient(client *http.Client) *DockerHubClient {
	if client == nil {
		client = http.DefaultClient
	}

	apiVersion := os.Getenv(dockerHubAPIVersion)
	if apiVersion == "" {
		apiVersion = defaultAPIVersion
	}

	hubURLAddr := os.Getenv(dockerHubURL)
	if hubURLAddr == "" {
		hubURLAddr = defaultHubURL
	}
	hubURL, err := url.Parse(hubURLAddr)
	if err != nil {
		panic(err)
	}

	c := &DockerHubClient{
		Client:     client,
		UserAgent:  "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1",
		APIVersion: apiVersion,
		HubURL:     hubURL,
	}

	return c
}

// NewRequest create new http request
func (c *DockerHubClient) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	relPath, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := c.HubURL.ResolveReference(relPath)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", mediaType)
	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

// Make the http request
func (c *DockerHubClient) Do(req *http.Request) (resp *http.Response, err error) {
	resp, err = c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
