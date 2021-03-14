package tagsalot

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type images struct {
	ImageName string   `yaml:"name"`
	ImageTags []string `yaml:"tags"`
}

type inputtedImages struct {
	Images []images `yaml:"images"`
}

// loadConfigMapFile returns struct for each image repo with tags provided
func LoadConfigMapFile(filePath string) *inputtedImages {
	var inputtedImagesInstance inputtedImages

	source, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &inputtedImagesInstance)
	if err != nil {
		log.Fatalf("error loading yaml: %v", err)
	}

	return &inputtedImagesInstance
}
