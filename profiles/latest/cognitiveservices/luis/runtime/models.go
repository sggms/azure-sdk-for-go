// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package runtime

import original "github.com/sggms/azure-sdk-for-go/services/cognitiveservices/v2.0/luis/runtime"

type BaseClient = original.BaseClient
type AzureRegions = original.AzureRegions

const (
	Australiaeast  AzureRegions = original.Australiaeast
	Brazilsouth    AzureRegions = original.Brazilsouth
	Eastasia       AzureRegions = original.Eastasia
	Eastus         AzureRegions = original.Eastus
	Eastus2        AzureRegions = original.Eastus2
	Northeurope    AzureRegions = original.Northeurope
	Southcentralus AzureRegions = original.Southcentralus
	Southeastasia  AzureRegions = original.Southeastasia
	Westcentralus  AzureRegions = original.Westcentralus
	Westeurope     AzureRegions = original.Westeurope
	Westus         AzureRegions = original.Westus
	Westus2        AzureRegions = original.Westus2
)

type APIError = original.APIError
type CompositeChildModel = original.CompositeChildModel
type CompositeEntityModel = original.CompositeEntityModel
type EntityModel = original.EntityModel
type EntityWithResolution = original.EntityWithResolution
type EntityWithScore = original.EntityWithScore
type IntentModel = original.IntentModel
type LuisResult = original.LuisResult
type Sentiment = original.Sentiment
type PredictionClient = original.PredictionClient

func New(azureRegion AzureRegions) BaseClient {
	return original.New(azureRegion)
}
func NewWithoutDefaults(azureRegion AzureRegions) BaseClient {
	return original.NewWithoutDefaults(azureRegion)
}
func PossibleAzureRegionsValues() []AzureRegions {
	return original.PossibleAzureRegionsValues()
}
func NewPredictionClient(azureRegion AzureRegions) PredictionClient {
	return original.NewPredictionClient(azureRegion)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
