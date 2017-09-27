package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var baseURL = "http://casaiq.com"

type Request struct {
	UDF        interface{} `json:"udf"`
	Name       string      `json:"name"`
	Enabled    bool        `json:"enabled"`
	PropertyID uint64      `json:"property_id"`
	HubCode    string      `json:"hub_code"`
	Type       string      `json:"type"`
}

type Property struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

type DeviceCount struct {
	Thermostat uint `json:"thermostat"`
}

type Response struct {
	Name        string      `json:"name"`
	Enabled     bool        `json:"enabled"`
	UDF         interface{} `json:"udf"`
	Devices     []uint64    `json:"devices"`
	Property    Property    `json:"property"`
	DeviceCount DeviceCount `json:"device_count"`
}

func CreateUnit(req Request) (Response, error) {
	fmt.Println("URL:>", baseURL)
	endpoint := "/api/units"

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(req)
	res, _ := http.Post(baseURL+endpoint, "application/json; charset=utf-8", b)

	var response Response
	json.NewDecoder(res.Body).Decode(&response)
	fmt.Println(response)
	return response, nil
}
