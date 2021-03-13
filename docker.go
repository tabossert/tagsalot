package tagsalot

import (
	"github.com/docker/docker/client"
	dockerclient "github.com/docker/docker/client"
)

// Client for local Docker
type ImageInfo struct {
	Client    *dockerclient.Client
	ImageName string
	ImageTags []string
}

func (i *ImageInfo) PullRecentImageTags() error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	for _, tag := range c.ImageTags {
		_, err := cli
	}

}
