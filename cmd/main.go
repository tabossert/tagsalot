package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tabossert/tagsalot"
)

type tag struct {
	Name string
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
		fmt.Println(content.Error)
		os.Exit(1)
	}

	var tags []tag
	_ = json.Unmarshal(body, &tags)
	for _, tag := range tags {
		fmt.Println(tag.Name)
	}

}
