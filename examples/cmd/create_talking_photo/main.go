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

	response, err := movioClient.CreateTalkingPhoto(context.Background(), &moviorequest.CreateTalkingPhoto{
		StorageKey: "prod/movio/url_upload/user_upload/99a5b551999a4c9183dd9944b53da168/5b85e9352acd43e7bc74a30bffd61c69.jpeg",
	})
	if err != nil {
		log.Fatal(err)
	}

	if !response.IsSuccess() {
		log.Fatal(response.Error())
	}

	bytes, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))

}
