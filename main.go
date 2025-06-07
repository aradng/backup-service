package main

import (
	"backup-service/internal/api"
	"backup-service/internal/docker"
	"backup-service/internal/utils"
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/docker/docker/client"
)

type App struct {
	Containers []docker.Container
}

func main() {
	app := &App{Containers: []docker.Container{}}
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	(*app).Containers, err = docker.GetContainers(ctx, cli)
	if err != nil {
		panic(err)
	}
	(*app).Containers = utils.Apply((*app).Containers, func(container docker.Container) docker.Container {
		val, _ := container.GetCredentials(ctx, cli)
		return val
	})

	for _, c := range (*app).Containers {
		if !c.DB {
			continue
		}
		val := reflect.ValueOf(c)
		typ := reflect.TypeOf(c)
		for k := range val.NumField() {
			field := typ.Field(k)
			value := val.Field(k)
			if value.Kind() == reflect.Ptr {
				if value.IsNil() {
					fmt.Printf("%s: <nil>\n", field.Name)
				} else {
					fmt.Printf("%s: %v\n", field.Name, value.Elem().Interface())
				}
			} else {
				fmt.Printf("%s: %v\n", field.Name, value.Interface())
			}
		}
		fmt.Printf("Hash: %s\n", c.Hash())
		fmt.Println("-----")
	}
	http.Handle("/api/", http.StripPrefix("/api", http.HandlerFunc(api.RouteHandlers)))
	http.ListenAndServe(":8090", nil)
}
