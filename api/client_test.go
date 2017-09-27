package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func startServer() {
	s := &http.Server{
		Addr: "localhost:8080",
	}

	s.ListenAndServe()
}

func TestCreateUnit(t *testing.T) {
	baseURL = "http://localhost"
	http.HandleFunc("/api/units", func(w http.ResponseWriter, r *http.Request) {
		res := &Response{
			Name:    "My Apartment",
			Enabled: true,
			Property: Property{
				Name: "Adam's Apt Complex",
				ID:   1,
			},
			UDF: nil,
			DeviceCount: DeviceCount{
				Thermostat: 10,
			},
			Devices: []uint64{1, 2, 3},
		}

		data, _ := json.Marshal(res)
		w.Write(data)
	})
	startServer()

	req := &Request{
		UDF:        nil,
		Name:       "Unit #60",
		Enabled:    true,
		PropertyID: 10,
		HubCode:    "ABCZYX",
		Type:       "common",
	}

	response, error := CreateUnit(*req)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(response)
}
