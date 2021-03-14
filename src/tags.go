package tagsalot

import (
	"fmt"
	"net/http"
)

type tag struct {
	Name string
}

type imageResults struct {
	Name string
	Tags []tag
}

var tagsPath = map[string]string{
	"1": "/v1/repositories/%s/tags",
	"2": "/v2/%s/tags/list",
}

// ListTags list all image tags
func (c *DockerHubClient) ListTags(image string) (*http.Response, error) {
	path := fmt.Sprintf(tagsPath[c.APIVersion], image)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PullImages downloads images for given repo name and tags
/*func PullImages(image string, responseBody []byte) error {
	imageTags := imageResults{Name: image}
	_ = json.Unmarshal(responseBody, &imageTags.Tags)
	err := imageTags.PullRecentImageTags()
	if err != nil {
		return err
	}

	return nil
}*/
