// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main provides example code for using jsonformat to move back and
// forth from FHIR JSON <--> FHIR Proto.
//
// To run with bazel:
//
//	bazel run //go/google/fhir_examples:parse_patients $WORKSPACE
//
// To run with native Go:
//
//	go run parse_patients.go $WORKSPACE
//
// where $WORKSPACE is the location of a synthea dataset.
// For instructions on setting up your workspace, see the top-level README.md
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/fhir/go/fhirversion"
	"github.com/google/fhir/go/jsonformat"
	r4pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
)

const (
	timeZone = "Australia/Sydney"
)

func GetPatientPOC() {
	httpGetURL := "http://localhost:8080/fhir/Patient/DDONYVATHBD6R3KW"

	request, error := http.NewRequest("GET", httpGetURL, nil)
	if error != nil {
		log.Fatalf("Failed to create request %v", error)
	}
	request.Header.Set("Content-Type", "application/fhir+json; charset=UTF-8")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		log.Fatalf("Failed to read response body %v", error)
	}

	um, err := jsonformat.NewUnmarshaller(timeZone, fhirversion.R4)
	if err != nil {
		log.Fatalf("Failed to create unmarshaller %v", err)
	}
	unmarshalled, err := um.Unmarshal(body)
	if err != nil {
		log.Fatalf("Failed to unmarshal patient %v", err)
	}

	contained := unmarshalled.(*r4pb.ContainedResource)
	patient := contained.GetPatient()
	fmt.Println(patient.GetName()[0].GetGiven())
}

func GetEncounterPOC() {
	httpGetURL := "http://localhost:8080/fhir/Encounter/DDONYVATHBD6R32Y"

	request, error := http.NewRequest("GET", httpGetURL, nil)
	if error != nil {
		log.Fatalf("Failed to create request %v", error)
	}
	request.Header.Set("Content-Type", "application/fhir+json; charset=UTF-8")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		log.Fatalf("Failed to read response body %v", error)
	}

	um, err := jsonformat.NewUnmarshaller(timeZone, fhirversion.R4)
	if err != nil {
		log.Fatalf("Failed to create unmarshaller %v", err)
	}
	unmarshalled, err := um.Unmarshal(body)
	if err != nil {
		log.Fatalf("Failed to unmarshal patient %v", err)
	}

	contained := unmarshalled.(*r4pb.ContainedResource)
	encounter := contained.GetEncounter()
	fmt.Println(encounter.GetSubject().GetPatientId())
}

// Appointment represents the structure of the FHIR Appointment resource
type Appointment struct {
	ResourceType    string `json:"resourceType"`
	ID              string `json:"id"`
	Meta            Meta   `json:"meta"`
	Text            Text   `json:"text"`
	Identifier      []Identifier `json:"identifier"`
	Status          string `json:"status"`
	ServiceCategory []ServiceCategory `json:"serviceCategory"`
	AppointmentType AppointmentType `json:"appointmentType"`
	ReasonCode      []ReasonCode `json:"reasonCode"`
	Priority        int `json:"priority"`
	Description     string `json:"description"`
	MinutesDuration int `json:"minutesDuration"`
	Created         string `json:"created"`
	Comment         string `json:"comment"`
	Participant     []Participant `json:"participant"`
}

type Meta struct {
	Security []Security `json:"security"`
}

type Security struct {
	System  string `json:"system"`
	Code    string `json:"code"`
	Display string `json:"display"`
}

type Text struct {
	Status string `json:"status"`
	Div    string `json:"div"`
}

type Identifier struct {
	System string `json:"system"`
	Value  string `json:"value"`
}

type ServiceCategory struct {
	Coding []Coding `json:"coding"`
}

type Coding struct {
	System  string `json:"system"`
	Code    string `json:"code"`
	Display string `json:"display"`
}

type AppointmentType struct {
	Coding []Coding `json:"coding"`
}

type ReasonCode struct {
	Coding []Coding `json:"coding"`
}

type Participant struct {
	Actor    Actor `json:"actor"`
	Required string `json:"required"`
	Status   string `json:"status"`
}

type Actor struct {
	Reference string `json:"reference"`
}

func CreateAppointmentPOC() {
	appointment := Appointment{
		ResourceType: "Appointment",
		ID:           "e1316ca3b7ca4c6b9314e7baaf64097b",
		Meta: Meta{
			Security: []Security{
				{
					System:  "http://terminology.hl7.org/CodeSystem/v3-ActReason",
					Code:    "HTEST",
					Display: "test health data",
				},
			},
		},
		Text: Text{
			Status: "generated",
			Div:    "<div xmlns=\"http://www.w3.org/1999/xhtml\">Touchstone Test Data - Appointment: A follow up visit from a previous appointment</div>",
		},
		Identifier: []Identifier{
			{
				System: "http://happyvalley.com/appointment",
				Value:  "13816582032-310",
			},
		},
		Status: "proposed",
		ServiceCategory: []ServiceCategory{
			{
				Coding: []Coding{
					{
						System:  "http://terminology.hl7.org/CodeSystem/service-type",
						Code:    "124",
						Display: "General Practice",
					},
				},
			},
		},
		AppointmentType: AppointmentType{
			Coding: []Coding{
				{
					System:  "http://terminology.hl7.org/CodeSystem/v2-0276",
					Code:    "FOLLOWUP",
					Display: "A follow up visit from a previous appointment",
				},
			},
		},
		ReasonCode: []ReasonCode{
			{
				Coding: []Coding{
					{
						System:  "http://snomed.info/sct",
						Code:    "813001",
						Display: "Ankle instability",
					},
				},
			},
		},
		Priority:        5,
		Description:     "Discuss results of recent MRI",
		MinutesDuration: 15,
		Created:         "2021-03-06",
		Comment:         "Further expand on the results of the MRI and determine the next actions that may be appropriate.",
		Participant: []Participant{
			{
				Actor: Actor{
					Reference: "Patient/DDONYVATHBD6R3KW",
				},
				Required: "required",
				Status:   "needs-action",
			},
		},
	}

	// Convert the appointment struct to JSON
	jsonData, err := json.Marshal(appointment)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create a PUT request with the JSON data
	req, err := http.NewRequest("PUT", "http://localhost:8080/fhir/Appointment/e1316ca3b7ca4c6b9314e7baaf64097b", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode == http.StatusOK {
		fmt.Println("PUT request successful")
	} else {
		fmt.Println("PUT request failed with status code:", resp.StatusCode)
	}
}

func GetAppointmentPOC() {
	httpGetURL := "http://localhost:8080/fhir/Appointment/e1316ca3b7ca4c6b9314e7baaf64097b"

	request, error := http.NewRequest("GET", httpGetURL, nil)
	if error != nil {
		log.Fatalf("Failed to create request %v", error)
	}

	request.Header.Set("Content-Type", "application/fhir+json; charset=UTF-8")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		log.Fatalf("Failed to read response body %v", error)
	}

	um, err := jsonformat.NewUnmarshaller(timeZone, fhirversion.R4)
	if err != nil {
		log.Fatalf("Failed to create unmarshaller %v", err)
	}
	unmarshalled, err := um.Unmarshal(body)
	if err != nil {
		log.Fatalf("Failed to unmarshal patient %v", err)
	}

	contained := unmarshalled.(*r4pb.ContainedResource)
	appointment := contained.GetAppointment()
	fmt.Println(appointment)
}

func main() {
	// GetPatientPOC()
	// GetEncounterPOC()
	GetAppointmentPOC()
	// CreateAppointmentPOC()
}
