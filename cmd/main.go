package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/tabossert/tagsalot"
)

type imageResults struct {
	Name string
	Tags []string
}

type errorResponse struct {
	Status string `json:"status_code"`
	Error  string
}

func main() {
	var image string
	flag.StringVar(&image, "image", "", "Image name")

	flag.Parse()

	c := tagsalot.NewHTTPClient(nil)

	resp, err := c.ListTags(image)
	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	if resp.StatusCode != 200 {
		var content errorResponse
		_ = json.Unmarshal(body, &content)
		log.Println(content.Error)
		os.Exit(1)
	}

	var tags []tag
	imageTags := imageResults{Name: image}
	_ = json.Unmarshal(body, &tags)
	imageTags.Tags = tags
	err := imageTags.PullRecentImageTags()
	if err {
		log.Println(err)
	}

}
