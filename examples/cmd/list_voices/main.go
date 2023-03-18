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

	responseStatus, err := movioClient.ListVoices(context.Background(), &moviorequest.ListVoices{})
	if err != nil {
		log.Fatal(err)
	}

	if !responseStatus.IsSuccess() {
		log.Fatal(responseStatus.Error())
	}

	bytes, _ := json.MarshalIndent(responseStatus, "", "  ")
	fmt.Println(string(bytes))

}
