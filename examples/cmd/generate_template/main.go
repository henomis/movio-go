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

	response, err := movioClient.GenerateTemplate(context.Background(), &moviorequest.GenerateTemplate{
		TemplateID: "516c266ceded4fc5aa81bb49e055a7b0",
		Test:       true,
		Title:      "Test",
		Variables: []moviorequest.Variable{
			{
				Name: "text",
				Properties: []moviorequest.Properties{
					{
						Default: "Hello happy World",
						Name:    "text",
					},
				},
			},
		},
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
