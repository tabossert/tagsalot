package main

import (
	"log"
	"os"

	"github.com/tabossert/tagsalot"
)

type errorResponse struct {
	Status string `json:"status_code"`
	Error  string
}

func main() {
	/*var image string

	c := tagsalot.NewHTTPClient(nil)
	resp, err := c.ListTags(image)
	if err != nil {
		log.Println(err)
	}
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
	}*/

	destinationRegistry := os.Getenv("DEST_REGISTRY")

	imageList := tagsalot.LoadConfigMapFile("../images.yaml")

	err := imageList.CopyImageTags(destinationRegistry)
	if err != nil {
		log.Println(err)
	}

}
