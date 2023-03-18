# Movio SDK for Go


[![Build Status](https://github.com/henomis/movio-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/henomis/movio-go/actions/workflows/test.yml?query=branch%3Amain) [![GoDoc](https://godoc.org/github.com/henomis/movio-go?status.svg)](https://godoc.org/github.com/henomis/movio-go) [![Go Report Card](https://goreportcard.com/badge/github.com/henomis/movio-go)](https://goreportcard.com/report/github.com/henomis/movio-go) [![GitHub release](https://img.shields.io/github/release/henomis/movio-go.svg)](https://github.com/henomis/movio-go/releases)

This is Movio's **unofficial** Go client, designed to enable you to use Movio's services easily from your own applications.

## Movio

Movio is a cloud-based analytics service that through APIs allows you generate video contents.

## SDK versions

API version | SDK version
--- | ---
v1 | v1.0.0


## Getting started

### Installation

You can load movio-go into your project by using:
```
go get github.com/henomis/movio-go
```


### Configuration

The only thing you need to start using Movio's APIs is the developer license key. Copy it and paste it in the corresponding place in the code, select the API you want to use and the parameters you want to use, and that's it.


### Usage

Please refer to the [examples folder](examples/cmd/) to see how to use the SDK.

Here below a simple usage example:

```go
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
```
