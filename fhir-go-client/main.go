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

func main() {
	GetPatientPOC()
	GetEncounterPOC()
}
