package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	moviogo "github.com/henomis/movio-go"
	moviorequest "github.com/henomis/movio-go/pkg/request"
)

func main() {

	movioClient := moviogo.New(
		moviogo.MovioEndpointV1,
		"YOUR_LICENSE_KEY",
	)

	response, err := movioClient.ListAvatars(context.Background(), &moviorequest.ListAvatars{})
	if err != nil {
		log.Fatal(err)
	}

	if !response.IsSuccess() {
		log.Println(response.RawBody)
		log.Fatal(response.Error())
	}

	bytes, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))

}
