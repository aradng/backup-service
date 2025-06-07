package docker

import (
	"backup-service/internal/model"
	"backup-service/internal/utils"
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var Containers []*Container

func GetContainers(ctx context.Context, cli *client.Client) ([]*Container, error) {
	services := make([]*Container, 0)

	conts, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	for _, cont := range conts {
		dcont := &Container{
			Names:   cont.Names,
			Image:   cont.Image,
			State:   cont.State,
			Status:  cont.Status,
			Project: cont.Labels["com.docker.compose.project"],
			Dir:     cont.Labels["com.docker.compose.project.working_dir"],
		}
		volumes := utils.Reduce(func(accumulator []string, mount container.MountPoint) []string {
			accumulator = append(accumulator, mount.Source)
			return accumulator
		}, cont.Mounts, []string{})
		sort.Strings(volumes)
		dcont.Volumes = volumes
		dcont.ID = dcont.Hash()
		services = append(services, dcont)
	}

	return services, nil
}

func (container *Container) GetCredentials(ctx context.Context, cli *client.Client) (*Container, error) {
	inspection, err := cli.ContainerInspect(ctx, container.ID)
	if err != nil {
		fmt.Printf("Error inspecting container %s: %v\n", container.ID, err)
		return container, err
	}

	for db := range model.SecretSchemas {
		if strings.Contains(container.Image, model.SecretSchemas[db].Name) {
			container.Type = &db
		}
	}
	if container.Type == nil {
		return container, nil
	}
	container.DB = true
	for _, env := range inspection.Config.Env {
		key := strings.Split(env, "=")[0]
		value := strings.Split(env, "=")[1]
		if key == model.SecretSchemas[*container.Type].Username {
			container.UserName = &value
		}
		if key == model.SecretSchemas[*container.Type].Password {
			container.Password = &value
		}
	}
	if container.Password != nil && container.UserName == nil && model.SecretSchemas[*container.Type].DefaultPassword != "" {
		schema := model.SecretSchemas[*container.Type]
		container.UserName = &schema.DefaultPassword
	}
	return container, nil
}

func init() {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	Containers, err = GetContainers(ctx, cli)
	if err != nil {
		panic(err)
	}
	for i, c := range Containers {
		val, err := c.GetCredentials(ctx, cli)
		if err != nil {
			fmt.Printf("Error getting credentials for container %s: %v\n", c.ID, err)
			continue
		}
		Containers[i] = val
	}
}
