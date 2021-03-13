package tagsalot

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
	dockType "github.com/docker/engine-api/types"
)

// Client for local Docker
type ImageInfo struct {
	ImageName string
	ImageTags []string
}

func (i *ImageInfo) PullRecentImageTags() error {
	ctx := context.Background()
	pullOptions := dockType.ImagePullOptions{}
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	for _, tag := range i.ImageTags {
		imageWithTag := fmt.Sprintf("%s:%s", i.ImageName, tag)
		_, err := cli.ImagePull(ctx, imageWithTag, pullOptions)
		if err != nil {
			return err
		}
	}

}
