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

	response, err := movioClient.GenerateVideo(context.Background(), &moviorequest.GenerateVideo{
		Background: "#ffffff",
		Clips: []moviorequest.Clip{
			{
				AvatarID:    "Daisy-inskirt-20220818",
				AvatarStyle: moviorequest.AvatarStyleNormal,
				InputText:   "Hello World",
				Offset: &moviorequest.Offset{
					X: 0,
					Y: 0,
				},
				Scale:   1,
				VoiceID: "1bd001e7e50f421d891986aad5158bc8",
			},
		},
		Ratio:   "16:9",
		Test:    true,
		Version: "v1alpha",
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
