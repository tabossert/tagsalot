package tagsalot

import (
	"context"
	"fmt"
	"io"
	"os"

	dockType "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func (i *inputtedImages) CopyImageTags(destinationRegistry string) error {
	ctx := context.Background()
	pullOptions := dockType.ImagePullOptions{}
	pushOptions := dockType.ImagePushOptions{
		All:          true,
		RegistryAuth: "123",
	}
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	for _, image := range i.Images {
		for _, tag := range image.ImageTags {
			imageWithTag := fmt.Sprintf("docker.io/library/%s:%s", image.ImageName, tag)
			destImageWithTag := fmt.Sprintf("%s/%s:%s", destinationRegistry, image.ImageName, tag)
			CreateECRRepository(image.ImageName)
			reader, err := cli.ImagePull(ctx, imageWithTag, pullOptions)
			if err != nil {
				return err
			}
			io.Copy(os.Stdout, reader)

			err = cli.ImageTag(ctx, imageWithTag, destImageWithTag)
			if err != nil {
				return err
			}

			reader, err = cli.ImagePush(ctx, destImageWithTag, pushOptions)
			if err != nil {
				return err
			}
			io.Copy(os.Stdout, reader)
		}
	}

	return nil

}
