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

package resources

import original "github.com/sggms/azure-sdk-for-go/services/resources/mgmt/2018-02-01/resources"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type DeploymentOperationsClient = original.DeploymentOperationsClient
type DeploymentsClient = original.DeploymentsClient
type GroupsClient = original.GroupsClient
type DeploymentMode = original.DeploymentMode

const (
	Complete    DeploymentMode = original.Complete
	Incremental DeploymentMode = original.Incremental
)

type OnErrorDeploymentType = original.OnErrorDeploymentType

const (
	LastSuccessful     OnErrorDeploymentType = original.LastSuccessful
	SpecificDeployment OnErrorDeploymentType = original.SpecificDeployment
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	None                       ResourceIdentityType = original.None
	SystemAssigned             ResourceIdentityType = original.SystemAssigned
	SystemAssignedUserAssigned ResourceIdentityType = original.SystemAssignedUserAssigned
	UserAssigned               ResourceIdentityType = original.UserAssigned
)

type AliasPathType = original.AliasPathType
type AliasType = original.AliasType
type BasicDependency = original.BasicDependency
type CreateOrUpdateByIDFuture = original.CreateOrUpdateByIDFuture
type CreateOrUpdateFuture = original.CreateOrUpdateFuture
type DebugSetting = original.DebugSetting
type DeleteByIDFuture = original.DeleteByIDFuture
type DeleteFuture = original.DeleteFuture
type Dependency = original.Dependency
type Deployment = original.Deployment
type DeploymentExportResult = original.DeploymentExportResult
type DeploymentExtended = original.DeploymentExtended
type DeploymentExtendedFilter = original.DeploymentExtendedFilter
type DeploymentListResult = original.DeploymentListResult
type DeploymentListResultIterator = original.DeploymentListResultIterator
type DeploymentListResultPage = original.DeploymentListResultPage
type DeploymentOperation = original.DeploymentOperation
type DeploymentOperationProperties = original.DeploymentOperationProperties
type DeploymentOperationsListResult = original.DeploymentOperationsListResult
type DeploymentOperationsListResultIterator = original.DeploymentOperationsListResultIterator
type DeploymentOperationsListResultPage = original.DeploymentOperationsListResultPage
type DeploymentProperties = original.DeploymentProperties
type DeploymentPropertiesExtended = original.DeploymentPropertiesExtended
type DeploymentsCreateOrUpdateFuture = original.DeploymentsCreateOrUpdateFuture
type DeploymentsDeleteFuture = original.DeploymentsDeleteFuture
type DeploymentValidateResult = original.DeploymentValidateResult
type ExportTemplateRequest = original.ExportTemplateRequest
type GenericResource = original.GenericResource
type GenericResourceFilter = original.GenericResourceFilter
type Group = original.Group
type GroupExportResult = original.GroupExportResult
type GroupFilter = original.GroupFilter
type GroupListResult = original.GroupListResult
type GroupListResultIterator = original.GroupListResultIterator
type GroupListResultPage = original.GroupListResultPage
type GroupPatchable = original.GroupPatchable
type GroupProperties = original.GroupProperties
type GroupsDeleteFuture = original.GroupsDeleteFuture
type HTTPMessage = original.HTTPMessage
type Identity = original.Identity
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type ManagementErrorWithDetails = original.ManagementErrorWithDetails
type MoveInfo = original.MoveInfo
type MoveResourcesFuture = original.MoveResourcesFuture
type OnErrorDeployment = original.OnErrorDeployment
type OnErrorDeploymentExtended = original.OnErrorDeploymentExtended
type ParametersLink = original.ParametersLink
type Plan = original.Plan
type Provider = original.Provider
type ProviderListResult = original.ProviderListResult
type ProviderListResultIterator = original.ProviderListResultIterator
type ProviderListResultPage = original.ProviderListResultPage
type ProviderOperationDisplayProperties = original.ProviderOperationDisplayProperties
type ProviderResourceType = original.ProviderResourceType
type Resource = original.Resource
type Sku = original.Sku
type SubResource = original.SubResource
type TagCount = original.TagCount
type TagDetails = original.TagDetails
type TagsListResult = original.TagsListResult
type TagsListResultIterator = original.TagsListResultIterator
type TagsListResultPage = original.TagsListResultPage
type TagValue = original.TagValue
type TargetResource = original.TargetResource
type TemplateLink = original.TemplateLink
type UpdateByIDFuture = original.UpdateByIDFuture
type UpdateFuture = original.UpdateFuture
type ValidateMoveResourcesFuture = original.ValidateMoveResourcesFuture
type ProvidersClient = original.ProvidersClient
type Client = original.Client
type TagsClient = original.TagsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewDeploymentOperationsClient(subscriptionID string) DeploymentOperationsClient {
	return original.NewDeploymentOperationsClient(subscriptionID)
}
func NewDeploymentOperationsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentOperationsClient {
	return original.NewDeploymentOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDeploymentsClient(subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClient(subscriptionID)
}
func NewDeploymentsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewGroupsClient(subscriptionID string) GroupsClient {
	return original.NewGroupsClient(subscriptionID)
}
func NewGroupsClientWithBaseURI(baseURI string, subscriptionID string) GroupsClient {
	return original.NewGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleDeploymentModeValues() []DeploymentMode {
	return original.PossibleDeploymentModeValues()
}
func PossibleOnErrorDeploymentTypeValues() []OnErrorDeploymentType {
	return original.PossibleOnErrorDeploymentTypeValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func NewProvidersClient(subscriptionID string) ProvidersClient {
	return original.NewProvidersClient(subscriptionID)
}
func NewProvidersClientWithBaseURI(baseURI string, subscriptionID string) ProvidersClient {
	return original.NewProvidersClientWithBaseURI(baseURI, subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewTagsClient(subscriptionID string) TagsClient {
	return original.NewTagsClient(subscriptionID)
}
func NewTagsClientWithBaseURI(baseURI string, subscriptionID string) TagsClient {
	return original.NewTagsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
