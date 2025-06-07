package main

import (
	"backup-service/api"
)

func main() {
	// containers := docker.Containers 


	// for _, c := range containers {
	// 	if !c.DB {
	// 		continue
	// 	}
	// 	val := reflect.ValueOf(*c)
	// 	typ := reflect.TypeOf(*c)
	// 	for k := range val.NumField() {
	// 		field := typ.Field(k)
	// 		value := val.Field(k)
	// 		if value.Kind() == reflect.Ptr {
	// 			if value.IsNil() {
	// 				fmt.Printf("%s: <nil>\n", field.Name)
	// 			} else {
	// 				fmt.Printf("%s: %v\n", field.Name, value.Elem().Interface())
	// 			}
	// 		} else {
	// 			fmt.Printf("%s: %v\n", field.Name, value.Interface())
	// 		}
	// 	}
	// 	fmt.Printf("Hash: %s\n", c.Hash())
	// 	fmt.Println("-----")
	// }

	api.App.Run(":8080")
}
